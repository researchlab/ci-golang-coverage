package pkg2

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSumInt(t *testing.T) {
	Convey("1+2+3+5 shouldEqual 11", t, func() {
		So(SumInt(1, 2, 3, 5), ShouldEqual, 11)
	})
}

func TestDivInt(t *testing.T) {
	Convey("将两数相除", t, func() {
		Convey("除以非 0 数", func() {
			num, err := DivInt(10, 2)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, 5)
		})

		Convey("除以 0", func() {
			_, err := DivInt(10, 0)
			So(err, ShouldNotBeNil)
		})
	})
}
