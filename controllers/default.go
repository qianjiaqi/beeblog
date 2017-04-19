package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)
var globalSession *session.Manager
type MainController struct {
	beego.Controller
}

func init() {
	globalSession, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "gclifetime":3600}`)
	go globalSession.GC()
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	sess := globalSession.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	sess.Set("mySession", "this is my session ------------------")
}
