package controller

import (
	"biscuits-forum/service"
	"biscuits-forum/tool"
	"github.com/gin-gonic/gin"
)

type OperateController struct {

}

func (oc *OperateController) Router (engine *gin.Engine) {
	engine.PUT("/operate/praise", oc.likes)
	engine.GET("/operate/collect/list", oc.getFavoriteList)
	engine.PUT("/operate/collect", oc.favorite)
	engine.GET("/operate/focus/list", oc.getFollowList)
	engine.PUT("/operate/focus", oc.followUser)
}

//点赞
func (oc *OperateController) likes (context *gin.Context){
	model := context.PostForm("model")
	id := context.PostForm("target_id")
	ops := new(service.OperateService)

	ok := ops.Likes(model, id)
	if !ok {
		tool.Failed(context, "点赞失败")
		return
	}
	tool.PrintInfo(context, "点赞成功")
}

//获取用户收藏列表
func (oc *OperateController) getFavoriteList (context *gin.Context){
	ops := new(service.OperateService)
	id := context.Query("user_id")

	favor, n := ops.GetFavorityList(id)
	for i := 0; i < n; i++{
		context.JSON(200,gin.H{
			"status":       	10000,
			"info":        		"success",
			"post_id": 			favor[i].PostId,
			"title": 			favor[i].Title,
			"publish_time":		favor[i].PublishTime,
			"user_id" :			favor[i].Nickname,
			"nickname":		 	favor[i].Nickname,
		})
	}
}

//收藏
func (oc *OperateController) favorite (context *gin.Context){
	id := context.PostForm("post_id")
	ops := new(service.OperateService)

	ok := ops.Focus(id)
	if !ok {
		tool.Failed(context, "收藏失败")
		return
	}
	tool.PrintInfo(context, "收藏成功")
}

//获取用户关注列表
func (oc *OperateController) getFollowList (context *gin.Context){
	ops := new(service.OperateService)
	id := context.Query("username")

	user, n := ops.GetFocusList(id)
	for i := 0; i < n; i++{
	context.JSON(200,gin.H{
		"status":       	10000,
		"info":        		"success",
		"username" :		user[i].Nickname,
		"nickname":		 	user[i].Nickname,
		"introducion":		user[i].Introduction,
	})
	}

}

//关注用户
func (oc *OperateController) followUser (context *gin.Context){
	username := context.PostForm("username")
	focusname := context.PostForm("focusname")
	ops := new(service.OperateService)

	ok := ops.FocusUser(username, focusname)
	if !ok {
		tool.Failed(context,"关注失败")
		return
	}
	tool.PrintInfo(context, "关注成功" +
		"")
}