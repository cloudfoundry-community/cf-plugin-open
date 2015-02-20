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
  Convey("checkArgs should call 0 commands", t, func() {
    err := checkArgs(cliConn, []string{"open", "test"})
    So(err, ShouldBeNil)
    })
  Convey("checkArgs should call 1 commands when calling open", t, func() {
    err := checkArgs(cliConn, []string{"open"})
    So(err, ShouldNotBeNil)
    })

  Convey("checkArgs should call 1 commands when calling open-service", t, func() {
    err := checkArgs(cliConn, []string{"service-open"})
    So(err, ShouldNotBeNil)
    })
}


func setup() {
  cliConn = &fakes.FakeCliConnection{}
}
