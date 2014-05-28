package app

import (
  //  "database/sql"
  //  "fmt"
  //  "github.com/coopernurse/gorp"
  //  _ "github.com/mattn/go-sqlite3"
  //  "log"
  "time"
)

type User struct {
  Id          uint32 `json:"id"`
  Name        string `json:"name"`
  CreatedBy   uint32
  UpdatedBy   uint32
  LockVersion uint32
  CreatedAt   time.Time
  UpdatedAt   time.Time
}
