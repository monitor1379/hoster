<!--
 * @Date: 2020-11-29 14:00:53
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 16:15:18
-->
# Hoster: A Golang library for manipulating your host file

Hoster is a cross-platform operating system host file management library.

## Installation

```
go get -u -v github.com/monitor1379/hoster
```


## Usage

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