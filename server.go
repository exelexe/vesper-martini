package main

import (
	"github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/h-miyamo/vesper-martini/app"
	"github.com/martini-contrib/cors"
)

func main() {
	// setup middleware
	m := martini.Classic()
	m.Use(render.Renderer())

	// CORS for http://localhost origins, allowing:
	// - GET methods
	// - Origin header
	// - Credentials share
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:		[]string{"http://localhost"},
		AllowMethods:		[]string{"GET"},
		AllowHeaders:		[]string{"Origin"},
		ExposeHeaders:		[]string{"Content-Length"},
		AllowCredentials:	true,
	}))

	// setup routes
	m.Get("/hello", app.Hello)
	m.Get("/user/:id", app.FindUser)

	m.Run()
}
