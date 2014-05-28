package tblmaps

import (
  "time"
)

type User struct {
  Id          uint64
  Name        string
  LockVersion uint
  CreatedBy   uint
  UpdatedBy   uint
  CreatedAt   time.Time
  UpdatedAt   time.Time
}
