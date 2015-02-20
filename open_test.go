package main

import (
    "testing"

    "github.com/cloudfoundry/cli/plugin/fakes"
    . "github.com/smartystreets/goconvey/convey"
  )

var (
    cliConn *fakes.FakeCliConnection
  )

func TestNoApp(t *testing.T){
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

func TestNoRoutes(t *testing.T){
  Convey("checkRoutes should return nil if route exists", t, func(){
    input := []string{"urls: google.com"}
    out, err := getUrlFromOutput(input)
    So(err, ShouldBeNil)
    So(out, ShouldEqual, "http://google.com")
    })

  Convey("checkRoutes should return error if no route exists", t, func(){
    input := []string{"urls: "}
    out, err := getUrlFromOutput(input)
    So(err, ShouldNotBeNil)
    So(out, ShouldEqual, "")
    })
}


func setup() {
  cliConn = &fakes.FakeCliConnection{}
}
