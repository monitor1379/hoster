package hoster_test

/*
 * @Date: 2020-11-29 14:10:50
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 14:21:14
 */

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/monitor1379/hoster"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMapping(t *testing.T) {
	Convey("TestMapping", t, func() {
		Convey("EmptyLine", func() {
			line := ""
			m, err := hoster.Decode(line)
			So(err, ShouldBeNil)
			So(m.Address, ShouldEqual, "")
			So(len(m.Hosts), ShouldEqual, 0)
			So(m.Comment, ShouldEqual, "")
			fmt.Printf("%s => %s \n", strconv.Quote(line), strconv.Quote(m.Encode()))
		})

		Convey("OnlyComment", func() {
			line := "# this\tis a comment"
			m, err := hoster.Decode(line)
			So(err, ShouldBeNil)
			So(m.Address, ShouldEqual, "")
			So(len(m.Hosts), ShouldEqual, 0)
			So(m.Comment, ShouldEqual, "# this\tis a comment")
			So(m.Encode(), ShouldEqual, "# this\tis a comment")
			fmt.Printf("%s => %s \n", strconv.Quote(line), strconv.Quote(m.Encode()))
		})

		Convey("IPv4", func() {
			line := "127.0.0.1\tlocalhost\t# this is   comment"
			m, err := hoster.Decode(line)
			So(err, ShouldBeNil)
			So(m.Address, ShouldEqual, "127.0.0.1")
			So(len(m.Hosts), ShouldEqual, 1)
			So(m.Hosts[0], ShouldEqual, "localhost")
			So(m.Comment, ShouldEqual, "# this is   comment")
			So(m.Encode(), ShouldEqual, "127.0.0.1\tlocalhost\t# this is   comment")
			fmt.Printf("%s => %s \n", strconv.Quote(line), strconv.Quote(m.Encode()))
		})

		Convey("IPv6", func() {
			line := "fe00::0 ip6-localnet"
			m, err := hoster.Decode(line)
			So(err, ShouldBeNil)
			So(m.Address, ShouldEqual, "fe00::0")
			So(len(m.Hosts), ShouldEqual, 1)
			So(m.Hosts[0], ShouldEqual, "ip6-localnet")
			So(m.Comment, ShouldEqual, "")
			So(m.Encode(), ShouldEqual, "fe00::0\tip6-localnet")
			fmt.Printf("%s => %s \n", strconv.Quote(line), strconv.Quote(m.Encode()))
		})

		Convey("MultiHosts", func() {
			line := "::1\t\tip6-localhost ip6-loopback"
			m, err := hoster.Decode(line)
			So(err, ShouldBeNil)
			So(m.Address, ShouldEqual, "::1")
			So(len(m.Hosts), ShouldEqual, 2)
			So(m.Hosts[0], ShouldEqual, "ip6-localhost")
			So(m.Hosts[1], ShouldEqual, "ip6-loopback")
			So(m.Comment, ShouldEqual, "")
			So(m.Encode(), ShouldEqual, "::1\tip6-localhost\tip6-loopback")
			fmt.Printf("%s => %s \n", strconv.Quote(line), strconv.Quote(m.Encode()))
		})

	})
}
