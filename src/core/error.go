package core

import (
  "log"
)

func CheckFatal(err error, msg string) {
  if err != nil {
    log.Fatalln(msg, err)
  }
}
