package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Article Model
type Article struct {
	Id        int
	Title     string
	Content   string `orm:"type(text)"`
	Thumbnail string
	Tags      []*Tag    `orm:"rel(m2m)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func GetArticleById(id int) (bool, Article) {
	o := orm.NewOrm()
	var article Article
	err := o.QueryTable(article).Filter("Id", id).One(&article)
	return err != orm.ErrNoRows, article
}

func CreateArticle(article *Article) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(article)
	return id
}

func UpdateArticle(article *Article) {
	o := orm.NewOrm()
	o.Update(article)
}
