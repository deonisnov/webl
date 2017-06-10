package main

import (
  "fmt"
  "github.com/deonisnov/webl/api"
)

func showVersion() {
  webl.INFO.Println(fmt.Sprintf("weblui %s", webl.Version()))
}
