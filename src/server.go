package main

import (
  "app"
  "core/db"
  "core/db/tblmaps"
  "github.com/go-martini/martini"
  "log"
  "os"
)

func main() {
  // init DB
  dbmap := db.GetDbmap()
  defer dbmap.Db.Close()

  // debug
  //log.Println(config.Get())
  log.Println(os.Getenv("GOPATH"))

  m := martini.Classic()

  // route handler
  m.Get("/hello1", hello)
  m.Get("/hello2", app.Hello)

  var user tblmaps.User
  sql := "select * from users where id = ?;"
  trans, err := dbmap.Begin()
  if err != nil {
    log.Println(err)
  }
  err = trans.SelectOne(&user, sql, 1)
  if err != nil {
    log.Println(err)
  } else {
    log.Println(user)
  }

  m.Get("/json/:id", func(params martini.Params) string {
    return "Hello " + params["name"]
  })

  m.Run()
}

func hello() (int, string) {
  return 200, "Hello world"
}
