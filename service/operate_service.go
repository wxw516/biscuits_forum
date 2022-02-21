package service

import (
	"biscuits-forum/dao"
	"biscuits-forum/model"
	"biscuits-forum/tool"
)

type OperateService struct {

}

//点赞
func (os *OperateService) Likes (model, id string) bool{
	od := dao.OperateDao{tool.GetDb()}

	if model == "1"{
		err := od.PostLikes(id)
		if err != nil {
			return false
		}
		return true
	}
	if model == "2"{
		err := od.ComentLikes(id)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

//收藏
func (os *OperateService) Focus (id string) bool{
	od := dao.OperateDao{tool.GetDb()}

	err := od.Focus(id)
	if err != nil{
		return false
	}
	return true
}

//获取收藏列表
func (os *OperateService) GetFavorityList (userId string) ([]model.Post,int){
	od := dao.OperateDao{tool.GetDb()}

	post, n := od.GetFavorityList(userId)
	return post, n
}

//获取关注列表
func (os *OperateService) GetFocusList (username string)([]model.User,int){
	od := dao.OperateDao{tool.GetDb()}

	user, n := od.GetFocusList(username)
	return user, n
}

//关注用户
func (os *OperateService) FocusUser (username, focusname string) bool{
	od := dao.OperateDao{tool.GetDb()}

	err := od.FocusUser(username, focusname)
	if err != nil {
		return false
	}
	return true
}