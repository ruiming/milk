package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

// @router / [get]
func (c *IndexController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["postlist"] = models.GetPostList(1)

	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
}
