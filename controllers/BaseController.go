package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	fmt.Println("prepare internal")
}

func (c *BaseController) GetInfos() {
	c.Ctx.Output.Body([]byte("prepare world info"))
}
