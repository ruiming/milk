package main

import (
	"fmt"

	"github.com/astaxie/beego"

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
		new(models.Article),
		new(models.Tag),
	)

	orm.RunSyncdb("default", false, true)

	fmt.Println("main init")
}

func main() {
	fmt.Println("main.go")
	o := orm.NewOrm()
	o.Using("default")

	tag := new(models.Tag)
	tag.Name = "test"

	fmt.Println(o.Insert(tag))

	article := new(models.Article)
	article.Title = "testTitleOfArticle"

	fmt.Println(o.Insert(article))

	m2m := o.QueryM2M(article, "Tags")

	num, err := m2m.Add(tag)

	if err == nil {
		fmt.Println(num)
		beego.Run()
	}

}
