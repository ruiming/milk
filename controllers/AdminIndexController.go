package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type AdminIndexController struct {
	beego.Controller
}

// @router /admin [get]
func (c *AdminIndexController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Articles"] = models.GetPostList(1)

	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/index.tpl"
}
