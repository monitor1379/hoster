package main

/*
 * @Date: 2020-11-29 15:37:19
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 16:25:48
 */

import (
	"fmt"

	"github.com/monitor1379/hoster"
)

/*

// File: ./hosts.txt

127.0.0.1 localhost

192.168.10.10 hoster-1.k8s.svc.cluster.local  hoster-2.k8s.svc.cluster.local # my kubernetes services
192.168.10.11 other-app.example.com


# The following lines are desirable for IPv6 capable hosts
::1     ip6-localhost ip6-loopback



*/

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
