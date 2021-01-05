package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (t *MainController) Get() {
	t.Data["Website"] = "beego.me"
	t.Data["Email"] = "astaxie@gmail.com"
	t.TplName = "index.tpl"
}
