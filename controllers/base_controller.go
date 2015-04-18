package controllers

import (
	"github.com/astaxie/beego"
)

type base_controller struct {
	beego.Controller
}

func (this *base_controller) SendJson(arr []map[string]string) {
	this.Data["json"] = arr
	this.ServeJson()
}
