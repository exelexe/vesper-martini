package db

import (
  "core/config"
  "database/sql"
  "github.com/coopernurse/gorp"
  _ "github.com/go-sql-driver/mysql"
  "log"
)

var dbmap *gorp.DbMap = initDb()

func initDb() *gorp.DbMap {
  // get json config
  c := config.Get()

  // debug
  log.Println("DB DataSource: " + c.Dsn)

  // connect to db using standard Go database/sql API
  // use whatever database/sql driver you wish
  db, err := sql.Open("mysql", c.Dsn)
  if err != nil {
    log.Println(err)
  }
  //CheckFatal(err, "sql.Open failed")

  // construct a gorp DbMap
  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

  return dbmap
}

func GetDbmap() *gorp.DbMap {
  return dbmap
}
