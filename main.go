package main

import (
	"fmt"

	"github.com/astaxie/beego"

	"blog/controllers"
	"blog/models"
	_ "blog/routers"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterDataBase("default", "mysql",
		beego.AppConfig.String("mysql.username")+":"+
			beego.AppConfig.String("mysql.password")+"@/"+
			beego.AppConfig.String("mysql.database")+"?"+
			"charset="+beego.AppConfig.String("mysql.charset"))

	orm.RegisterModel(
		new(models.Comment),
		new(models.Option),
		new(models.Tag),
		new(models.Post),
	)

	orm.RunSyncdb("default", false, true)

	fmt.Println("main init")
}

func main() {
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.AdminIndexController{})
	beego.Include(&controllers.AdminPostController{})

	beego.Run()

}
