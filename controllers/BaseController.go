package controllers

import (
	//"context"
	//"context"
	"context"
	"fmt"

	"github.com/astaxie/beego"

	"go.elastic.co/apm/module/apmmongo/v2"
	"go.elastic.co/apm/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var ctx = context.Background()
//var mongoClient *mongo.Client
var mongoClient, _ = mongo.Connect(
	context.Background(),
	options.Client().SetMonitor(apmmongo.CommandMonitor()),
	options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"),
)
var collection *mongo.Collection

type BaseController struct {
	beego.Controller
	tracer *apm.Tracer
}

func (c *BaseController) Prepare() {
	//fmt.Println("prepare internal")
	span, _ := apm.StartSpan(c.Ctx.Request.Context(), fmt.Sprintf("%s:%s", c.Ctx.Request.Method, c.Ctx.Request.URL.Path), "controller")
	span.End()

	// mongoClient, _ = mongo.Connect(
	// 	c.Ctx.Request.Context(),
	// 	options.Client().SetMonitor(apmmongo.CommandMonitor()),
	// 	options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"),
	// )

}

func (c *BaseController) GetInfos() {

	c.Ctx.Output.Body([]byte("prepare world info"))
}

func (c *BaseController) MongoPing() {
	// span, _ := apm.StartSpan(c.Ctx.Request.Context(), fmt.Sprintf("%s:%s", c.Ctx.Request.Method, c.Ctx.Request.URL.Path), "controller")
	// span.End()

	span, _ := apm.StartSpan(c.Ctx.Request.Context(), "find_db_APM poc", "custom")
	defer span.End()

	collection = mongoClient.Database("dbAPM").Collection("poc")
	_, err := collection.Find(c.Ctx.Request.Context(), bson.D{})

	if err != nil {
		fmt.Println("Error", err)
	}

	c.Ctx.Output.Body([]byte("message pong"))
}

func (c *BaseController) MongoInsert() {
	// span, _ := apm.StartSpan(c.Ctx.Request.Context(), fmt.Sprintf("%s:%s", c.Ctx.Request.Method, c.Ctx.Request.URL.Path), "controller")
	// span.End()
	collection = mongoClient.Database("dbAPM").Collection("poc")
	result, err := collection.InsertOne(c.Ctx.Request.Context(), bson.M{"foo": "bar1", "hello": "world1", "pi": 3.14159})

	collection.Find(c.Ctx.Request.Context(), bson.D{})
	fmt.Println("RESULT IS", result)

	if err != nil {
		fmt.Println("Error", err)
		c.Ctx.Output.Body([]byte("error"))
	}
	c.Ctx.Output.Body([]byte("result"))

}

// func (c *BaseController) MysqlInsert() {
// 	//?????????????????????????????????
// 	db := util.InitDB()
// 	//????????????
// 	tx, err := db.Begin()
// 	if err != nil {
// 		//???????????????????????????panic
// 		panic(err)
// 	}
// 	//??????SQL??????
// 	sql := "insert into tb_user (`name`, `password`) values (?, ?)"
// 	//???SQL?????????????????????
// 	stmt, err := db.Prepare(sql)
// 	if err != nil {
// 		panic(err)
// 	}
// 	result, err := stmt.Exec("????????????", "123")
// 	if err != nil {
// 		//SQL?????????????????????panic
// 		panic(err)
// 	}
// 	//????????????
// 	tx.Commit()
// 	//?????????????????????id
// 	fmt.Println(result.LastInsertId())

// }
