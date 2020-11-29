package main

/*
 * @Date: 2020-11-29 15:37:19
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 16:43:17
 */

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
