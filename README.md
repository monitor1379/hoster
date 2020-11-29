<!--
 * @Date: 2020-11-29 14:00:53
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 23:21:07
-->


<img src="./docs/logo/logo-with-name.png" style="height:100px" alt="hoster" height="100px"/>
<img src="https://travis-ci.org/monitor1379/hoster.svg" />

Hoster is a cross-platform operating system host file management library written in Go.


[Features](#features) | [Installation](#installation) | [Usage](#usage) | [CLI](#cli)

## Features

- Cross platform: support Linux / Windows / macOS(darwin)
- Lookup by address(IPv4, IPv6 etc.) and host(domains).
- Host file backup and duplication.
- CLI `cmd/hoster` for manipulating host file in terminal.

## Installation

Requirements: 
- Golang version v1.13.0 (minimum)


Installation:
```
go get -u -v github.com/monitor1379/hoster
```


## Usage


### Create

Create a `*HostManager` and print the content of host file:

```golang
package main


import (
	"fmt"
	"time"

	"github.com/monitor1379/hoster"
)

func main() {
	// create a *HostManger
	hm, err := hoster.Default()
	if err != nil {
		panic(err)
	}

	// print:
	// "/etc/hosts" in non-Windows OS
	// or
	// "C:\Windows\System32\drivers\etc\hosts" in Windows
	fmt.Println(hm.HostFilePath())

	// print your host file content
	fmt.Println(hm.String())
}

```


### Lookup

Assuming your host file content is:
```bash
127.0.0.1 localhost

192.168.10.10 hoster-1.k8s.svc.cluster.local  hoster-2.k8s.svc.cluster.local    # my kubernetes services
192.168.10.11 other-app.example.com


# The following lines are desirable for IPv6 capable hosts
::1     ip6-localhost ip6-loopback
```

Lookup by address(ip) or host(domain):
```golang
package main

import (
	"fmt"

	"github.com/monitor1379/hoster"
)

func main() {
	hm, err := hoster.New("./hosts.txt")
	if err != nil {
		panic(err)
	}

	mapping, ok := hm.LookupByAddress("192.168.10.10")
	fmt.Println(ok)              // true
	fmt.Println(mapping.Address) // 192.168.10.10
	fmt.Println(mapping.Hosts)   // []string{"hoster-1.k8s.svc.cluster.local", "hoster-2.k8s.svc.cluster.local"}
	fmt.Println(mapping.Comment) // # my kubernetes services

	mapping, ok = hm.LookupByHost("hoster-2.k8s.svc.cluster.local")
	fmt.Println(ok)              // true
	fmt.Println(mapping.Address) // 192.168.10.10
	fmt.Println(mapping.Hosts)   // []string{"hoster-1.k8s.svc.cluster.local", "hoster-2.k8s.svc.cluster.local"}
	fmt.Println(mapping.Comment) // # my kubernetes services
}

```



### Set and Delete


```golang
package main

import (
	"fmt"

	"github.com/monitor1379/hoster"
)

func main() {
	// hm is managing "./hosts.txt"
	hm, err := hoster.New("./hosts.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(hm.String())
	// print:
	// 127.0.0.1       localhost
	//
	// 192.168.10.10   hoster-1.k8s.svc.cluster.local  hoster-2.k8s.svc.cluster.local  # my kubernetes services
	// 192.168.10.11   other-app.example.com
	//
	// # The following lines are desirable for IPv6 capable hosts
	// ::1     ip6-localhost   ip6-loopback

	// hm2 is managing "./hosts2.txt" which content is copied from "./hosts.txt"
	hm2, err := hm.Duplicate("./host2.txt")
	if err != nil {
		panic(err)
	}

	err = hm2.Set("hoster-2.k8s.svc.cluster.local", "192.168.10.12", "# added by monitor1379")
	if err != nil {
		panic(err)
	}
	fmt.Println(hm2.String())
	// print:
	// 127.0.0.1       localhost
	//
	// 192.168.10.10   hoster-1.k8s.svc.cluster.local  # my kubernetes services
	// 192.168.10.11   other-app.example.com
	//
	// # The following lines are desirable for IPv6 capable hosts
	// ::1     ip6-localhost   ip6-loopback
	// 192.168.10.12   hoster-2.k8s.svc.cluster.local  # added by monitor1379

	err = hm2.DeleteHost("hoster-2.k8s.svc.cluster.local")
	if err != nil {
		panic(err)
	}
	fmt.Println(hm2.String())
	// print:
	// 127.0.0.1       localhost
	//
	// 192.168.10.10   hoster-1.k8s.svc.cluster.local  # my kubernetes services
	// 192.168.10.11   other-app.example.com
	//
	// # The following lines are desirable for IPv6 capable hosts
	// ::1     ip6-localhost   ip6-loopback

}

```

### Backup

backup host file to a new path:
```golang
package main

import (
	"github.com/monitor1379/hoster"
)

func main() {
	hm, err := hoster.New("./hosts.txt")
	if err != nil {
		panic(err)
	}

	// note that after backup, hm is still managing "./hosts.txt"
	err = hm.Backup("./hosts-backup.txt")
	if err != nil {
		panic(err)
	}
}

```


# CLI

`cmd/hoster` is a command-line program for manipulating host file in terminal.

common methods:
- `hoster version`
- `hoster list [--file <host-file-path>]`
- `hoster lookup [--file <host-file-path>]  [--address <address>] [--host <host>]`
- `hoster set [--file <host-file-path>]  [--address <address>] [--host <host>] [--comment <comment>`


```bash
> hoster

Hoster is a cross-platform operating system host file management library written in Go.

Usage:
  hoster [command]

Available Commands:
  help        Help about any command
  list        List address-host mappings
  lookup      lookup by address or host
  set         Add a address-host mapping
  version     Print the version of hoster

Flags:
  -f, --file string   specify a host file path (default "/etc/hosts")
  -h, --help          help for hoster

Use "hoster [command] --help" for more information about a command.
```


List command:

```bash
> hoster list

+-----------+-----------------+
|  ADDRESS  |      HOST       |
+-----------+-----------------+
| 127.0.0.1 | localhost       |
| 127.0.1.1 | red-coast-base  |
| ::1       | ip6-localhost   |
| ::1       | ip6-loopback    |
| fe00::0   | ip6-localnet    |
| ff00::0   | ip6-mcastprefix |
| ff02::1   | ip6-allnodes    |
| ff02::2   | ip6-allrouters  |
+-----------+-----------------+
```

