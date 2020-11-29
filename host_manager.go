package hoster

/*
 * @Date: 2020-11-29 14:01:23
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 16:12:58
 */

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/monitor1379/hoster/pkg/linesep"
)

type HostManager struct {
	hostFilePath string
	mappings     []*Mapping
	mu           *sync.RWMutex
}

// New returns a `*HostManager` with given host file path.
func New(hostFilePath string) (*HostManager, error) {
	if !isPathExist(hostFilePath) {
		return nil, os.ErrNotExist
	}

	mappings, err := ParseFile(hostFilePath)
	if err != nil {
		return nil, err
	}

	h := &HostManager{
		hostFilePath: hostFilePath,
		mappings:     mappings,
		mu:           &sync.RWMutex{},
	}
	return h, nil
}

// Default returns a `*HostManager` with system host file path.
func Default() (*HostManager, error) {
	var hostFilePath string

	switch runtime.GOOS {

	case "linux", "darwin":
		hostFilePath = "/etc/hosts"

	case "windows":
		// Some Windows operating system did not install on "C:" partition,
		// use %windir% to get system partition is more safely.
		winDirPath := os.Getenv("windir")
		if winDirPath == "" {
			return nil, errors.New("can not found environment variable \"windir\"")
		}
		hostFilePath = filepath.Join(winDirPath, "System32", "drivers", "etc", "hosts")

	// TODO(monitor1379): support more os platforms.
	default:
		return nil, fmt.Errorf("unsupported os type: \"%s\"", runtime.GOOS)
	}

	h, err := New(hostFilePath)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (h *HostManager) HostFilePath() string {
	return h.hostFilePath
}

func (h *HostManager) Backup(backupHostFilePath string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	return flush(backupHostFilePath, h.mappings)
}

func (h *HostManager) Duplicate(newHostFilePath string) (*HostManager, error) {

	h.mu.Lock()
	defer h.mu.Unlock()
	err := flush(newHostFilePath, h.mappings)
	if err != nil {
		return nil, err
	}

	newH, err := New(newHostFilePath)
	if err != nil {
		return nil, err
	}
	return newH, nil
}

func (h *HostManager) LookupByAddress(address string) (*Mapping, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, mapping := range h.mappings {
		if mapping.Address == address {
			m := new(Mapping)
			*m = *mapping
			return m, true
		}
	}
	return nil, false
}

func (h *HostManager) LookupByHost(host string) (*Mapping, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, mapping := range h.mappings {
		for _, mappingHost := range mapping.Hosts {
			if mappingHost == host {
				m := new(Mapping)
				*m = *mapping
				return m, true
			}
		}
	}
	return nil, false
}

// Set set the ip of host as address, and flush to host file.
func (h *HostManager) Set(host, address, comment string) error {
	if address == "" && host == "" && comment == "" {
		return errors.New("can not set an empty content")
	}
	if (address != "" && host == "") || (address == "" && host != "") {
		return errors.New("invalid format")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	newMappings := make([]*Mapping, 0)
	for _, mapping := range h.mappings {
		diffHosts, _ := diffset(mapping.Hosts, []string{host})
		if len(diffHosts) != 0 {
			mapping.Hosts = diffHosts
		} else {
			continue
		}
		newMappings = append(newMappings, mapping)
	}
	newMappings = append(newMappings, &Mapping{
		Address: address,
		Hosts:   []string{host},
		Comment: comment,
	})
	h.mappings = newMappings

	err := flush(h.hostFilePath, h.mappings)
	if err != nil {
		return err
	}
	return nil
}

func (h *HostManager) DeleteHost(host string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	newMappings := make([]*Mapping, 0)
	for _, mapping := range h.mappings {
		diffHosts, _ := diffset(mapping.Hosts, []string{host})
		if len(diffHosts) != 0 {
			mapping.Hosts = diffHosts
		} else {
			continue
		}
		newMappings = append(newMappings, mapping)
	}

	h.mappings = newMappings

	err := flush(h.hostFilePath, h.mappings)
	if err != nil {
		return err
	}
	return nil
}

func (h *HostManager) String() string {
	s := make([]string, 0)
	for _, mapping := range h.mappings {
		s = append(s, mapping.Encode())
	}
	return strings.Join(s, linesep.DefaultLineSeperator)
}

func parse(data []byte) ([]*Mapping, error) {
	mappings := make([]*Mapping, 0)
	for _, line := range strings.Split(string(data), linesep.DefaultLineSeperator) {
		mapping, err := Decode(line)
		if err != nil {
			return nil, err
		}
		mappings = append(mappings, mapping)
	}
	return mappings, nil
}

func Parse(data []byte) ([]*Mapping, error) {
	return parse(data)
}

func ParseFile(hostFilePath string) ([]*Mapping, error) {
	data, err := ioutil.ReadFile(hostFilePath)
	if err != nil {
		return nil, err
	}
	return parse(data)
}

func flush(hostFilePath string, mappings []*Mapping) error {
	s := make([]string, 0)
	for _, mapping := range mappings {
		s = append(s, mapping.Encode())
	}
	return ioutil.WriteFile(hostFilePath, []byte(strings.Join(s, linesep.DefaultLineSeperator)), os.ModePerm)
}

func Flush(hostFilePath string, mappings []*Mapping) error {
	return flush(hostFilePath, mappings)
}
