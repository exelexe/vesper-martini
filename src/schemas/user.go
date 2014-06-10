package schemas

import (
  "time"
)

type User struct {
  Id          int       `db:"id"`
  Name        string    `db:"name"`
  LockVersion int       `db:"lock_version"`
  CreatedBy   int       `db:"created_by"`
  UpdatedBy   int       `db:"updated_by"`
  CreatedAt   time.Time `db:"created_at"`
  UpdatedAt   time.Time `db:"updated_at"`
}
