// use martini as web framework
package main

import (
	//"fmt"
	//"net/http"
	//"time"
	//"html/template"

	"./utils"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	//martini.Logger()
	// render html templates from templates directory
	m.Use(render.Renderer(render.Options{
		Directory:  "templets",                 // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs:      []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		//Delims:     render.Delims{"{[{", "}]}"},    // Sets delimiters to the specified strings.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
	}))

	m.Get("/", func(r render.Render) {
		//messages := "系统将在3秒后跳转后后台登录页......"
		p := utils.NewPlatform()
		p.GetAll()
		r.HTML(200, "hemo", p)
	})

	m.Get("/admin/login", func(r render.Render) {
		token := "1234567890"
		r.HTML(200, "admin/login", token)
	})

	m.Run()
}
