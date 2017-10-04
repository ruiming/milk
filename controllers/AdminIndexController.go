package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type AdminIndexController struct {
	beego.Controller
}

func (c *AdminIndexController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Articles"] = models.GetArticleList(1)

	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/index.tpl"
}
