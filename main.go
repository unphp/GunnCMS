// use martini as web framework
package main

import (
	"./routers"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

/**
* Load all Controllers.
 */
func loadControllers(app *martini.ClassicMartini) {
	routers.NewBaseRouter(app)
	routers.NewBaseAdminRouter(app)
}

/**
* Martini application configuration.
 */
func configuration(app *martini.ClassicMartini) {
	app.Use(martini.Static("public"))
	app.Use(render.Renderer(render.Options{
		Directory:  "templets",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))
}

func main() {
	app := martini.Classic()
	configuration(app)
	loadControllers(app)
	app.Run()
}
