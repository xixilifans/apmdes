package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/astaxie/beego"
	"github.com/yufeng/test/controllers"

	"go.elastic.co/apm/module/apmbeego"
	"go.elastic.co/apm/v2"
)

type thingController struct {
	beego.Controller
}

func (c *thingController) Get() {
	span, _ := apm.StartSpan(c.Ctx.Request.Context(), "thingController.Get", "controller")
	span.End()
	fmt.Println("holo figf")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	c.Ctx.Output.Body([]byte("hello world"))

}

type MongoClient struct {
	DatabaseName string
	_uri         string
	Database     *mongo.Database
	Client       *mongo.Client
}

// var Client *mongo.Client

// func init() {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	// Connect to MongoDB
// 	client, _ := mongo.Connect(context.TODO(), options.Client().SetMonitor(apmmongo.CommandMonitor()), clientOptions)
// 	Client = client
// }

func main() {
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// // Connect to MongoDB
	// client, _ := mongo.Connect(context.TODO(), options.Client().SetMonitor(apmmongo.CommandMonitor()), clientOptions)

	// _ = client.Ping(context.TODO(), nil)
	//collection := client.Database("testdb").Collection("people")

	beego.Router("/infos", &controllers.BaseController{}, "get:GetInfos")
	beego.Router("/gg", &controllers.BaseController{}, "get:GetGG")
	beego.Router("/", &thingController{})
	beego.Router("/thing/:id:int", &thingController{}, "get:Get")
	beego.RunWithMiddleWares("localhost:8080", apmbeego.Middleware())

}
