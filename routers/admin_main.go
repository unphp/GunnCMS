// admin_main.go
package routers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

/**
* BaseRouter structure.
 */
type BaseAdminRouter struct {
	app *martini.ClassicMartini
}

func NewBaseAdminRouter(app *martini.ClassicMartini) BaseAdminRouter {
	c := BaseAdminRouter{app}
	c.routes()
	return c
}

func (this *BaseAdminRouter) routes() {
	this.app.Get("/admin/login", this.login)
	this.app.Get("/admin/main", this.dashboard)
	this.app.Get("/admin/logout", this.logout)
}

/**
* action.
 */
func (this *BaseAdminRouter) login(r render.Render) {
	r.HTML(200, "admin/login", nil)
}

func (this *BaseAdminRouter) dashboard(r render.Render) {
	r.HTML(200, "admin/dashboard", nil)
}

func (this *BaseAdminRouter) logout(r render.Render) {
	json := map[string]interface{}{"msg": "ok", "status": true}
	r.JSON(200, json)
}
