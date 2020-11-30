package main

/*
 * @Date: 2020-11-29 15:37:19
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 16:14:45
 */

import (
	"fmt"
	"time"

	"github.com/monitor1379/hoster"
)

func main() {
	// create a *HostManager
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

	time.Sleep(10 * time.Second)
}
