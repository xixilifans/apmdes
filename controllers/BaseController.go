package controllers

import (
	//"context"
	//"context"
	"fmt"

	"github.com/astaxie/beego"

	"go.elastic.co/apm/module/apmmongo"
	"go.elastic.co/apm/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var ctx = context.Background()

// var mongoClient, _ = mongo.Connect(
// 	c.Ctx.Request.Context(),
// 	options.Client().SetMonitor(apmmongo.CommandMonitor()),
// 	options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"),
// )
var collection *mongo.Collection

type BaseController struct {
	beego.Controller
	tracer *apm.Tracer
}

var mongoClient *mongo.Client

func (c *BaseController) Prepare() {
	//fmt.Println("prepare internal")
	span, _ := apm.StartSpan(c.Ctx.Request.Context(), fmt.Sprintf("%s:%s", c.Ctx.Request.Method, c.Ctx.Request.URL.Path), "controller")
	span.End()

	mongoClient, _ = mongo.Connect(
		c.Ctx.Request.Context(),
		options.Client().SetMonitor(apmmongo.CommandMonitor()),
		options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"),
	)

}

// func (c *BaseController) GetInfos() {

// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	// Connect to MongoDB
// 	client, _ := mongo.Connect(context.TODO(), options.Client().SetMonitor(apmmongo.CommandMonitor()), clientOptions)
// 	collection := client.Database("users").Collection("system.users")
// 	fmt.Println(collection.Find(context.TODO(), nil))
// 	c.Ctx.Output.Body([]byte("prepare world info"))
// }

func (c *BaseController) MongoPing() {
	collection = mongoClient.Database("dbAPM").Collection("poc")
	_, err := collection.Find(c.Ctx.Request.Context(), bson.D{})
	if err != nil {
		fmt.Println("Error", err)
	}

	c.Ctx.Output.Body([]byte("message pong"))
}

func (c *BaseController) MongoInsert() {
	collection = mongoClient.Database("dbAPM").Collection("poc")
	result, err := collection.InsertOne(c.Ctx.Request.Context(), bson.M{"foo": "bar1", "hello": "world1", "pi": 3.14159})
	fmt.Println("RESULT IS", result)

	if err != nil {
		fmt.Println("Error", err)
		c.Ctx.Output.Body([]byte("error"))
	}
	c.Ctx.Output.Body([]byte("result"))

}
