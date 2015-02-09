package main

import (
    //"errors"
    "testing"

    "github.com/cloudfoundry/cli/plugin/fakes"
    . "github.com/smartystreets/goconvey/convey"
    //io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
    //. "github.com/cloudfoundry-community/cli-plugin-open"
  )

var (
    cliConn *fakes.FakeCliConnection
  )

func TestNoApp(t *testing.T){
  setup()
  Convey("checkArgs should call 0 commands", t, func() {
    checkArgs(cliConn, []string{"open", "test"})
    So(cliConn.CliCommandCallCount(), ShouldEqual, 0)
    })
  Convey("checkArgs should call 1 commands when calling open", t, func() {
    checkArgs(cliConn, []string{"open"})
    So(cliConn.CliCommandCallCount(), ShouldEqual, 1)
    })
}


func setup() {
  cliConn = &fakes.FakeCliConnection{}
}
