package controller

import (
	"biscuits-forum/service"
	"biscuits-forum/tool"
	"github.com/gin-gonic/gin"
)

type CommentController struct {

}

func (cc *CommentController) Router (engine *gin.Engine){
	engine.GET("/comment", cc.getPostComment)
	engine.POST("/comment", cc.publishComment)
	engine.PUT("/comment/{comment_id}", cc.chanegComment)
	engine.DELETE("/comment/{comment_id}", cc.deleteComment)
}

//获取某个帖子/评论下的评论
func (cc *CommentController) getPostComment (context *gin.Context){
	model := context.GetInt("model")
	targetId := context.Query("target_id")
	page := context.GetInt("page")
	size := context.GetInt("size")
	cs := new(service.CommentService)

	cm := cs.GetComment(model, page, size, targetId)
	if cm == nil {
		tool.Failed(context, "参数获取失败")
		return
	}

	for i := 0; i < size; i++{
		context.JSON(200, gin.H{
			"status":       	10000,
			"info":        			"success",
			"comment_id": 			cm[i].CommentId,
			"post_id": 				cm[i].PostId,
			"publish_time": 		cm[i].PublishTime,
			"content":				cm[i].Content,
			"pictures":				cm[i].Pictures,
			"user_id":				cm[i].UserId,
			"avatar":				cm[i].Avater,
			"nickname":				cm[i].Nickname,
			"reply_user_id":		cm[i].ReplyUserId,
			"reply_user_nickname":	cm[i].RelpyUserNickname,
			"praise_count":			cm[i].PraiseCount,
			"is_praised":			cm[i].IsPraise,
			"is_focus":				cm[i].IsFocus,
		})
	}
}

//发布评论
func (cc *CommentController) publishComment (context *gin.Context){
	model := context.PostForm("model")
	targetId := context.PostForm("target_id")
	content := context.PostForm("content")
	cs := new(service.CommentService)

	ok := cs.PublishComment(model, targetId, content)
	if !ok {
		tool.Failed(context,"发表评论失败")
		return
	}
	tool.PrintInfo(context, "发布成功")
}

//更新评论
func (cc *CommentController) chanegComment (context *gin.Context){
	commentId := context.PostForm("comment_id")
	content := context.PostForm("content")
	cs := new(service.CommentService)

	ok := cs.ChangeComment(commentId, content)
	if !ok {
		tool.Failed(context, "更改失败")
		return
	}
	tool.PrintInfo(context, "更改成功")
}

//删除评论
func (cc *CommentController) deleteComment (context *gin.Context){
	commentId := context.PostForm("comment_id")
	cs := new(service.CommentService)

	ok := cs.DeleteComent(commentId)
	if !ok{
		tool.Failed(context, "删除失败")
		return
	}
	tool.PrintInfo(context,"成功删除")
}