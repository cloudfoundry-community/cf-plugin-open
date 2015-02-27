package main

import (
	"testing"
	"strings"

	"github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	cliConn *fakes.FakeCliConnection
)

func TestNoApp(t *testing.T) {
	setup()
	Convey("checkArgs should not return error with open test", t, func() {
		err := checkArgs(cliConn, []string{"open", "test"})
		So(err, ShouldBeNil)
	})

	Convey("checkArgs should not return error with service-open test", t, func() {
		err := checkArgs(cliConn, []string{"service-open", "test"})
		So(err, ShouldBeNil)
	})

	Convey("checkArgs should return error with open", t, func() {
		err := checkArgs(cliConn, []string{"open"})
		So(err, ShouldNotBeNil)
	})

	Convey("checkArgs should return error with service-open", t, func() {
		err := checkArgs(cliConn, []string{"service-open"})
		So(err, ShouldNotBeNil)
	})
}

func TestNoRoutes(t *testing.T) {
	Convey("getUrlFromOuput should return nil if route exists", t, func() {
		input := []string{"urls: google.com"}
		out, err := getUrlFromOutput(input)
		So(err, ShouldBeNil)
		So(out[0], ShouldEqual, "http://google.com")
	})

	Convey("getUrlFromOuput should return error if no route exists", t, func() {
		input := []string{"urls: "}
		out, err := getUrlFromOutput(input)
		So(err, ShouldNotBeNil)
		So(out[0], ShouldEqual, "")
	})

	Convey("getUrlFromOuput should handle multiple routes", t, func() {
		input := []string{"urls: google.com, apple.com, github.com"}
		out, err := getUrlFromOutput(input)
		So(err, ShouldBeNil)
		So(out[0], ShouldEqual, "http://google.com")
		So(out[1], ShouldEqual, "http://apple.com")
		So(out[2], ShouldEqual, "http://github.com")
	})
}

func TestRoutesMenu(t *testing.T) {
	Convey("multiRoutesMenu should return url if there is one route", t, func() {
		input := []string{"http://google.com"}
		So(multiRoutesMenu(strings.NewReader(""), input), ShouldEqual, "http://google.com")
	})

	Convey("multiRoutesMenu should return 1nd url if first route is chosen", t, func() {
		input := []string{"http://google.com", "http://apple.com"}
		So(multiRoutesMenu(strings.NewReader("1"), input), ShouldEqual, "http://google.com")
	})

	Convey("multiRoutesMenu should return 2nd url if second route is chosen", t, func() {
		input := []string{"http://google.com", "http://apple.com"}
		So(multiRoutesMenu(strings.NewReader("2"), input), ShouldEqual, "http://apple.com")
	})
}

func setup() {
	cliConn = &fakes.FakeCliConnection{}
}
