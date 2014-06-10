package main

import (
  "app"
  . "core/db"
  "github.com/go-martini/martini"
  //"log"
  "schemas"
)

func main() {
  m := martini.Classic()

  // route handler
  m.Get("/hello1", hello)
  m.Get("/hello2", app.Hello)

  m.Get("/json/:id", func(params martini.Params) string {
    var user schemas.User
    DB.First(&user, params["id"])

    //return "Hello " + params["id"]
    return "name: " + user.Name
  })

  m.Run()
}

func hello() (int, string) {
  return 200, "Hello world"
}

//func findUser(id int) tblmaps.User {
//  var u tblmaps.User
//
//  dbmap := db.GetDbmap()
//  defer dbmap.Db.Close()
//  sql := "select id, name from users where id = ?;"
//  trans, err := dbmap.Begin()
//  if err != nil {
//    log.Println(err)
//  }
//  err = trans.SelectOne(&u, sql, id)
//  if err != nil {
//    log.Println(err)
//  }
//
//  err = trans.Commit()
//
//  return u
//}
