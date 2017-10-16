package controllers

import (
	"blog/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type AdminPostController struct {
	beego.Controller
}

// @router /admin/post/:page [get]
func (c *AdminPostController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	page, e := strconv.Atoi(c.Ctx.Input.Param(":page"))

	if e == nil {
		fmt.Println("page", page)
		c.Data["Articles"] = models.GetPostList(page)
	}

	c.Layout = "admin/layout.tpl"
	c.TplName = "admin/index.tpl"
}
