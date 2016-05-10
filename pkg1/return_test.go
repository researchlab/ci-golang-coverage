package pkg1

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReturnStr(t *testing.T) {
	Convey("the ReturnStr shouldEqual 'ABC'", t, func() {
		So(ReturnStr(), ShouldEqual, "ABC")
	})
}

func TestReturnInt(t *testing.T) {
	Convey("The ReturnInt ShouldEqual 10", t, func() {
		So(ReturnInt(), ShouldEqual, 10)
	})
}
