package controller

import (
	"biscuits-forum/dao"
	"biscuits-forum/model"
	"biscuits-forum/service"
	"biscuits-forum/tool"
	"github.com/gin-gonic/gin"
)

type PostController struct {

}

func (pc *PostController) Router (engine *gin.Engine) {
	engine.GET("/post/list", pc.getPostList)
	engine.GET("/post/single/{post_id}", pc.getPostInfo)
	engine.POST("/post/single", pc.publishPost)
	engine.PUT("/post/single/{post_id}", pc.publishPost)
	engine.DELETE("/post/single/{post_id}", pc.deletePost)
	engine.GET("/post/search", pc.searchPost)
	engine.GET("/topic/list", pc.getPostTopic)
}

//获取帖子列表
func (pc *PostController) getPostList (context *gin.Context){
	var page, size int
	var post []model.PostInfo
	ps := new(service.PostService)

	page = context.GetInt("page")
	size = context.GetInt("size")
	n := (page - 1) * size + 1

	ok := ps.CheckPostId(string(n))
	if !ok {
		tool.Failed(context, "帖子不存在")
	}

	post = ps.GetPagePost(n, size)
	for i := 0; i < size; i++{
		context.JSON(200,gin.H{
			"status":       	10000,
			"info":        		"success",
			"post_id":			post[i].PostId,
			"is_praise":	 	post[i].IsPraised,
			"is_focus": 		post[i].IsFocus,
			"comment_count": 	post[i].CommentCount,
			"praise_count":  	post[i].PraiseCount,
			"title": 		 	post[i].Title,
			"publish_time":  	post[i].PublishTime,
			"content": 		 	post[i].Content,
			"picture": 			post[i].Pictures,
			"nickname": 		post[i].Nickname,
			"avatar": 			post[i].Avatar,
		})
	}
}

//获取某个帖子内容
func (pc *PostController) getPostInfo (context *gin.Context){
	postId := context.PostForm("post_id")
	ps := new(service.PostService)
	var post model.PostInfo

	ok := ps.CheckPostId(postId)
	if !ok {
		tool.Failed(context, "帖子不存在")
	}

	post = ps.GetPost(postId)
	context.JSON(200, gin.H{
		"status":       	10000,
		"info":        		"success",
		"post_id":			post.PostId,
		"is_praise":	 	post.IsPraised,
		"is_focus": 		post.IsFocus,
		"comment_count": 	post.CommentCount,
		"praise_count":  	post.PraiseCount,
		"title": 		 	post.Title,
		"publish_time":  	post.PublishTime,
		"content": 		 	post.Content,
		"picture": 			post.Pictures,
		"nickname": 		post.Nickname,
		"avatar": 			post.Avatar,
	})
}

//发布帖子
func (pc *PostController) publishPost (context *gin.Context){
	var post model.PostInfo
	ps := new(service.PostService)

	err := context.ShouldBind(&post)
	if err != nil {
		tool.PrintInfo(context,"参数获取失败！")
		return
	}

	ok := ps.CheckPost(post)
	if ok != "ok"{
		tool.Failed(context, ok)
		return
	}

	ok1 := ps.InsertPost(post)
	if !ok1 {
		tool.Failed(context,"发布失败")
		return
	}
	tool.PrintInfo(context,"发布成功")
}

//更新帖子
func (pc *PostController) changePost (context *gin.Context){
	var post model.PostInfo
	ps := new(service.PostService)

	err := context.ShouldBind(&post)
	if err != nil {
		tool.PrintInfo(context,"参数获取失败！")
		return
	}

	ok := ps.CheckPost(post)
	if ok != "ok"{
		tool.Failed(context, ok)
		return
	}

	ok1 := ps.ChangePost(post)
	if !ok1 {
		tool.Failed(context,"发布失败")
		return
	}
	tool.PrintInfo(context,"发布成功")
}

//删除帖子
func (pc *PostController) deletePost (context *gin.Context){
	ps := new(service.PostService)
	postId := context.PostForm("post_id")

	ok := ps.CheckPostId(postId)
	if !ok {
		tool.Failed(context,"未找到该帖子")
		return
	}

	olk :=ps.DeletePost(postId)
	if !olk {
		tool.Failed(context,"删除失败")
		return
	}
	tool.PrintInfo(context,"删除成功")
}

//搜索帖子
func (pc *PostController) searchPost (context *gin.Context){
	var page, size int
	var post []model.PostInfo
	ps := new(service.PostService)

	page = context.GetInt("page")
	size = context.GetInt("size")
	n := (page - 1) * size + 1

	ok := ps.CheckPostId(string(n))
	if !ok {
		tool.Failed(context, "帖子不存在")
	}

	post = ps.GetPagePost(n, size)
	for i := 0; i < size; i++{
		context.JSON(200,gin.H{
			"status":       	10000,
			"info":        		"success",
			"post_id":			post[i].PostId,
			"is_praise":	 	post[i].IsPraised,
			"is_focus": 		post[i].IsFocus,
			"comment_count": 	post[i].CommentCount,
			"praise_count":  	post[i].PraiseCount,
			"title": 		 	post[i].Title,
			"publish_time":  	post[i].PublishTime,
			"content": 		 	post[i].Content,
			"picture": 			post[i].Pictures,
			"nickname": 		post[i].Nickname,
			"avatar": 			post[i].Avatar,
		})
	}
}

//获取所有主题
func (pc *PostController) getPostTopic (context *gin.Context){
	pd := dao.PostDao{tool.GetDb()}
	topic, n := pd.GetPostTopic()
	for i := 0; i < n; i++{
		context.JSON(200, gin.H{
			"topic": topic[i],
 		})
	}
}