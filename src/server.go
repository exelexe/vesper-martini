package main

import (
	"app"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
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
