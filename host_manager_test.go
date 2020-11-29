/*
 * @Date: 2020-11-29 14:54:33
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 15:31:02
 */
package hoster_test

import (
	"fmt"
	"testing"

	"github.com/monitor1379/hoster"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHostManager(t *testing.T) {
	Convey("TestHostManager", t, func() {
		Convey("Lookup", func() {
			hm, err := hoster.New("./tests/test-hosts-1.txt")
			So(err, ShouldBeNil)

			var m *hoster.Mapping
			var ok bool

			// 127.0.0.1       localhost
			m, ok = hm.LookupByAddress("127.0.0.1")
			So(ok, ShouldBeTrue)
			So(m.Address, ShouldEqual, "127.0.0.1")
			So(len(m.Hosts), ShouldEqual, 1)
			So(m.Hosts[0], ShouldEqual, "localhost")

			// ::1     ip6-localhost ip6-loopback
			m, ok = hm.LookupByAddress("::1")
			So(ok, ShouldBeTrue)
			So(m.Address, ShouldEqual, "::1")
			So(len(m.Hosts), ShouldEqual, 2)
			So(m.Hosts[0], ShouldEqual, "ip6-localhost")
			So(m.Hosts[1], ShouldEqual, "ip6-loopback")
		})

		Convey("SetAndDelete", func() {
			// hm is managing host file "./tests/test-hosts-2.txt"
			hm, err := hoster.New("./tests/test-hosts-2.txt")
			So(err, ShouldBeNil)
			fmt.Printf("\n%s\n", hm.String())

			// hm is managing host file "./tests/test-hosts-3.txt"
			// which content is copied from "./tests/test-hosts-2.txt"
			hm, err = hm.Duplicate("./tests/test-hosts-3.txt")
			So(err, ShouldBeNil)

			var m *hoster.Mapping
			var ok bool

			// set: 127.0.1.1	red-coast-base	# added by monitor1379
			err = hm.Set("red-coast-base", "127.0.1.1", "# added by monitor1379")
			So(err, ShouldBeNil)
			fmt.Printf("\nAfter add red-coast-base:\n%s\n", hm.String())

			// delete a host
			err = hm.DeleteHost("test4.domain.com")
			So(err, ShouldBeNil)
			fmt.Printf("\nAfter delete test4.domain.com:\n%s\n", hm.String())

			// delete a host whose ip has other hosts
			err = hm.DeleteHost("test.domain.com")
			So(err, ShouldBeNil)
			fmt.Printf("\nAfter delete test.domain.com:\n%s\n", hm.String())
			m, ok = hm.LookupByHost("test.domain.com")
			So(ok, ShouldBeFalse)
			So(m, ShouldBeNil)

			m, ok = hm.LookupByAddress("127.0.0.1")
			So(ok, ShouldBeTrue)
			So(m.Address, ShouldEqual, "127.0.0.1")
			So(len(m.Hosts), ShouldEqual, 2)
			So(m.Hosts[0], ShouldEqual, "localhost")
			So(m.Hosts[1], ShouldEqual, "test1.domain.com")

		})
	})

}
