// routers.go
package routers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

/**
* BaseRouter structure.
 */
type BaseRouter struct {
	app *martini.ClassicMartini
}

func NewBaseRouter(app *martini.ClassicMartini) BaseRouter {
	c := BaseRouter{app}
	c.routes()
	return c
}

/**
* define controller routing.
 */
func (this *BaseRouter) routes() {
	this.app.Get("/", this.index)
	this.app.NotFound(this.httpNotFound)
}

/**
* index action.
 */
func (this *BaseRouter) index(r render.Render) {
	r.HTML(200, "home", nil)
}

/**
* httpNotFound action.
 */
func (this *BaseRouter) httpNotFound() string {
	return "HTTP NOT FOUND !"
}
