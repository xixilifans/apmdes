package main

import (
	"github.com/astaxie/beego"
	"github.com/yufeng/test/controllers"

	"go.elastic.co/apm/module/apmbeego/v2"
)

// type thingController struct {
// 	beego.Controller
// }

// func (c *thingController) Get() {
// 	span, _ := apm.StartSpan(c.Ctx.Request.Context(), "thingController.Get", "controller")
// 	span.End()
// 	fmt.Println("holo figf")
// 	c.Data["Website"] = "beego.me"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.tpl"

// 	c.Ctx.Output.Body([]byte("hello world"))

// }

func main() {

	beego.Router("/infos", &controllers.BaseController{}, "get:GetInfos")
	beego.Router("/mongo/ping", &controllers.BaseController{}, "get:MongoPing")
	beego.Router("/mongo/insert", &controllers.BaseController{}, "get:MongoInsert")
	//beego.Router("/mysql/insert", &controllers.BaseController{}, "get:MysqlInsert")
	//beego.Router("/", &thingController{})
	//beego.Router("/thing/:id:int", &thingController{}, "get:Get")
	beego.RunWithMiddleWares("localhost:8080", apmbeego.Middleware())

}
