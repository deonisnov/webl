package main

import (
  "fmt"
  "github.com/deonisnov/webl/api"
)

func showVersion() {
  webl.INFO.Println(fmt.Sprintf("weblonsole %s", webl.Version()))
}

func showMissingUrl() {
  webl.WARN.Println("Please provide a URL to parse, e.g.")
  webl.WARN.Println("webl -url=a4word.com")
}
