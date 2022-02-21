package service

import (
	"biscuits-forum/dao"
	"biscuits-forum/model"
	"biscuits-forum/tool"
)

type CommentService struct {

}

//获取评论
func (cs *CommentService) GetComment (model, page, size int, targetId string) []model.CommentInfo{
	cd := dao.CommentDao{tool.GetDb()}
	if model == 1{
		cm := cd.GetPostComent(targetId, page, size)
		return cm
	}
	if model == 2{
		cn := cd.GetCommentComment(targetId, page, size)
		return cn
	}
	return nil
}

//发表评论
func (cs *CommentService) PublishComment (model, targetId, content string) bool{
	cd := dao.CommentDao{tool.GetDb()}
	if model == "1"{
		err := cd.PostComment(targetId,content)
		if err != nil {
			return false
		}
	return true
	}
	if model == "2"{
		err := cd.PublishCommentComment(targetId,content)
		if err != nil{
			return false
		}
		return true
	}
	return false
}

//更新评论
func (cs *CommentService) ChangeComment (commentId, content string) bool{
	cd := dao.CommentDao{tool.GetDb()}
	err := cd.ChangeComment(commentId, content)
	if err != nil {
		return false
	}
	return true
}

//删除评论
func (cs *CommentService) DeleteComent (id string) bool{
	cd := dao.CommentDao{tool.GetDb()}
	err := cd.DeleteComent(id)
	if err != nil {
		return  false
	}
	return true
}