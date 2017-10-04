package models

import "github.com/astaxie/beego/orm"

// Option Model
type Option struct {
	Name  string `orm:"pk"`
	Value string
}

// GetOptionList 获取配置列表
func GetOptionList() []Option {
	o := orm.NewOrm()
	var option Option
	var list []Option
	o.QueryTable(option).All(&list)
	return list
}

// GetOptionValue 获取配置
func GetOptionValue(name string) (bool, Option) {
	o := orm.NewOrm()
	var option Option
	err := o.QueryTable(option).Filter("name", name).One(&option)
	return err != orm.ErrMissPK, option
}

// CreateOption 创建配置
func CreateOption(option *Option) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(option)
	return id
}

// UpdateOption 更新配置
func UpdateOption(option *Option) {
	o := orm.NewOrm()
	o.Update(option)
}
