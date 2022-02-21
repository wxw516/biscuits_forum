package service

import (
	"biscuits-forum/dao"
	"biscuits-forum/model"
	"biscuits-forum/tool"
)

type PostService struct {

}

//获取帖子列表


//判断帖子是否存在
func (ps *PostService) CheckPostId (postId string) bool{
	pd := dao.PostDao{tool.GetDb()}
	return pd.CheckPostAlive(postId)
}

//获取帖子
func (ps *PostService) GetPost (postId string) model.PostInfo{
	pd := dao.PostDao{tool.GetDb()}
	return pd.GetPost(postId)
}

//审核发布的帖子
func (ps *PostService) CheckPost (info model.PostInfo) string{
	err1 := ps.CheckPostTitle(info.Title)
	if err1 != "ok"{
		return err1
	}

	err2 := ps.CheckPostContent(info.Content)
	if err2 != "ok"{
		return err2
	}

	err3 := ps.CheckTopic(info.TopicId)
	if err3 != "ok"{
		return err3
	}

	err4 := ps.CheckPostPhoto(info.Pictures)
	if err4 != "ok"{
		return err4
	}
	return "ok"
}


//判断帖子标题是否符合规范
func (ps *PostService) CheckPostTitle (title string) string{
	if len(title) > 10||len(title) < 1{
		return "title不符合规范"
	}
	return "ok"
}

//判断内容是否符合规范
func (ps *PostService) CheckPostContent (content string) string{
	if len(content) > 3000||len(content) < 1{
		return "内容不符合规范（不可以涩涩哦）"
	}
	return "ok"
}

//判断帖子id是否存在
func (ps *PostService) CheckTopic (topic string) string{
	if len(topic) > 3000||len(topic) < 1{
		return "不可以涩涩哦"
	}
	return "ok"
}

//判断图片是否符合规范
func (ps *PostService) CheckPostPhoto (photo []string) string{
	if len(photo) > 9{
		return "图片数量最多为9张"
	}
	return "ok"
}

//存储帖子信息
func (ps *PostService) InsertPost (post model.PostInfo) bool{
	pd := dao.PostDao{tool.GetDb()}
	err := pd.InsertPost(post)
	if err != nil {
		return false
	}
	return true
}

//更新帖子
func (ps *PostService) ChangePost (post model.PostInfo) bool{
	pd := dao.PostDao{tool.GetDb()}
	err := pd.ChangePost(post)
	if err != nil {
		return false
	}
	return true
}

//删除帖子
func (ps *PostService) DeletePost (postId string) bool{
	pd := dao.PostDao{tool.GetDb()}
	err := pd.DeletePost(postId)
	if err != nil {
		return false
	}
	return true
}

//获取一页帖子
func (ps *PostService) GetPagePost (n, size int) []model.PostInfo{
	var page []model.PostInfo
	pd := dao.PostDao{tool.GetDb()}

	for i := 0; i < size; i++{
		page[i] = pd.GetPost(string(n))
		n++
	}
	return page
}