package utils

import (
	"github.com/astaxie/beego"
)

func increase(in int) (out int) {
	out = in + 1
	return
}

func decrease(in int) (out int) {
	out = in - 1
	return
}

func init() {
	beego.AddFuncMap("increase", increase)
	beego.AddFuncMap("decrease", decrease)
}
