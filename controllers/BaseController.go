package controllers

import (
	"context"
	"fmt"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"github.com/prometheus/common/log"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmmongo/v2"
)

type BaseController struct {
	beego.Controller
	tracer *apm.Tracer
}

func (c *BaseController) Prepare() {
	fmt.Println("prepare internal")
	span, _ := apm.StartSpan(c.Ctx.Request.Context(), fmt.Sprintf("%s:%s", c.Ctx.Request.Method, c.Ctx.Request.URL.Path), "controller")

	span.End()

}

func (c *BaseController) GetInfos() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, _ := mongo.Connect(context.TODO(), options.Client().SetMonitor(apmmongo.CommandMonitor()), clientOptions)
	collection := client.Database("users").Collection("system.users")
	fmt.Println(collection.Find(context.TODO(), nil))
	c.Ctx.Output.Body([]byte("prepare world info"))
}

func (c *BaseController) GetGG() {
	c.Ctx.Output.Body([]byte("prepare wGGGG info"))
}
