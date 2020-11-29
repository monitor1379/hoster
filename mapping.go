package hoster

/*
 * @Date: 2020-11-29 14:03:49
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 14:12:30
 */

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Mapping struct {
	Address string
	Hosts   []string
	Comment string
}

func (m *Mapping) Decode(line string) error {
	*m = Mapping{}

	if line == "" {
		return nil
	}

	if []rune(line)[0] == '#' {
		m.Comment = line
		return nil
	}

	i := strings.Index(line, "#")
	if i != -1 {
		m.Comment = line[i:]
		line = line[:i]
	}

	line = strings.ReplaceAll(line, "\t", "    ")
	items := strings.Split(line, " ")
	items = reduceEmptyItem(items)

	if len(items) < 2 {
		return errors.New("invalid format")
	}

	// TODO(monitor1379): validate address
	m.Address = items[0]

	// TODO(monitor1379): validate hosts
	m.Hosts = items[1:]

	return nil
}

func (m Mapping) Encode() string {
	var s []string
	if m.Address != "" && len(m.Hosts) != 0 {
		s = append(s, m.Address, strings.Join(m.Hosts, "\t"))
	}
	if m.Comment != "" {
		s = append(s, m.Comment)
	}
	return strings.Join(s, "\t")
}

func (m Mapping) String() string {
	return fmt.Sprintf("{Mapping : %s}", strconv.Quote(m.Encode()))
}

func Decode(line string) (*Mapping, error) {
	m := new(Mapping)
	err := m.Decode(line)
	if err != nil {
		return nil, err
	}
	return m, nil
}
