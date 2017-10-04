package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})

	beego.Router("/admin", &controllers.AdminIndexController{})
}
