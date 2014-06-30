package main

import (
	"github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/h-miyamo/vesper-martini/app"
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
