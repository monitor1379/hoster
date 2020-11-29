package main

/*
 * @Date: 2020-11-29 15:37:19
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 16:48:43
 */

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
