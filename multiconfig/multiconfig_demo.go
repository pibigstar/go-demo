package main

import (
  "github.com/koding/multiconfig"
  "github.com/smallnest/rpcx/log"
  "fmt"
)

type Server struct {
  Name    string
  Port    int    `default:"6060"`
  Enabled bool
  Users   []string
}

func main() {

  m := multiconfig.NewWithPath("multiconfig/config/config.toml")
  serverConf := new(Server)

  err := m.Load(serverConf) // Check for error
  if err!=nil {
    log.Errorf("%s",err.Error())
  }
  m.MustLoad(serverConf)    // Panic's if there is any error

  fmt.Println(serverConf.Port)

}
