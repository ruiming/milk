package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["blog/controllers:AdminIndexController"] = append(beego.GlobalControllerRouter["blog/controllers:AdminIndexController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/admin`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:AdminPostController"] = append(beego.GlobalControllerRouter["blog/controllers:AdminPostController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/admin/post/:page`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
