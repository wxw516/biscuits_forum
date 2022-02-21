package main

import (
	"biscuits-forum/controller"
	"biscuits-forum/tool"
	"github.com/gin-gonic/gin"
)

func main()  {

	cfg, err := tool.ParseConfig("./config/web.json")
	if err != nil {
		panic(err.Error())
	}

	tool.SqlEngine()

	web := gin.Default()

	registerRouter(web)

	web.Run(cfg.WebHost + ":" + cfg.WebPort)
}

//路由设置
func registerRouter(router *gin.Engine)  {
	new(controller.HelloController).Router(router)
	new(controller.UserController).Router(router)
	new(controller.PostController).Router(router)
	new(controller.CommentController).Router(router)
	new(controller.OperateController).Router(router)
}