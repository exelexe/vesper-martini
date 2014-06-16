package main

import (
    "app"
    //    . "core/db"
    "github.com/go-martini/martini"
    //    "log"
    //    "schemas"
)

func main() {
    m := martini.Classic()

    // route handler
    m.Get("/hello", app.Hello)
    m.Get("/user/:id", app.FindUser)
    /*
       m.Get("/json/:id", func(params martini.Params) string {
           var user schemas.User
           DB.First(&user, params["id"])
           log.Println(user)
           //return "Hello " + params["id"]
           return "name: " + user.Name
       })
    */
    m.Run()
}
