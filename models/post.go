package models

import (
	"blog/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Post Model
type Post struct {
	Id        int
	Title     string
	Content   string `orm:"type(text)"`
	Thumbnail string
	View      string
	Like      string
	Comments  []*Comment `orm:"reverse(many)"`
	Tags      []*Tag     `orm:"rel(m2m)"`
	CreatedAt time.Time  `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time  `orm:"auto_now_add;type(datetime)"`
}

// GetPostById 根据文章 ID 获取文章信息
func GetPostById(id int) (bool, Post) {
	o := orm.NewOrm()
	var post Post
	err := o.QueryTable(post).Filter("Id", id).One(&post)
	return err != orm.ErrNoRows, post
}

// GetPostList 获取文章列表
func GetPostList(page int) utils.Page {
	perpage := beego.AppConfig.DefaultInt("perpage", 10)
	o := orm.NewOrm()
	var post Post
	var list []Post

	qs := o.QueryTable(post)

	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-CreatedAt").Limit(perpage).Offset((page - 1) * perpage).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, page, perpage, list)
}

// CreatePost 创建文章
func CreatePost(post *Post) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(post)
	return id
}

// UpdatePost 更新文章
func UpdatePost(post *Post) {
	o := orm.NewOrm()
	o.Update(post)
}
