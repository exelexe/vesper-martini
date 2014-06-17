package main

import (
    "app"
    "github.com/go-martini/martini"
    "github.com/codegangsta/martini-contrib/render"
)

func main() {
    // setup middleware
    m := martini.Classic()
    m.Use(render.Renderer())

    // setup routes
    m.Get("/hello", app.Hello)
    m.Get("/user/:id", app.FindUser)

    m.Run()
}
