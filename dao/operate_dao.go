package dao

import (
	"biscuits-forum/model"
	"database/sql"
	"fmt"
)

type OperateDao struct {
	*sql.DB
}

//给帖子点赞
func (od *OperateDao) PostLikes (id string) error{
	_,err := od.Exec("update postinfo set is_praise = ? where post_id = ?","ture", id)
	return err
}

//给评论点赞
func (od *OperateDao) ComentLikes (id string) error{
	_,err := od.Exec("update commentinfo set is_praise = ? where comment_id = ?","ture", id)
	return err
}

//收藏
func (od *OperateDao) Focus (id string) error{
	_,err := od.Exec("update postinfo set is_focus = ? where post_id = ?","ture", id)
	return err
}

//获取收藏列表
func (od *OperateDao) GetFavorityList (userid string) ([]model.Post, int){
	var post []model.Post
	m := 0
	for a := 1; ; a++{
		row := od.QueryRow("select title, nickname, post_idm, publish_time, nickname from postinfo where post_id = ? ", userid)

		err := row.Err()
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		var str model.Post
		row.Scan(&str.Title, &str.Nickname, &str.PostId, &str.PublishTime, &str.UseName)
		if str.PostId == ""{
			break
		}
		m++
		post[a-1] = str
	}
	return post, m
}

//获取关注列表
func (od *OperateDao) GetFocusList (username string)([]model.User,int){
	var user []model.User
	m := 0
	for a := 1; ; a++{
		row := od.QueryRow("select username, nickname, introduction from userinfo where username = ? ", username)

		err := row.Err()
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		var str model.User
		row.Scan(&str.Username, &str.Nickname, &str.Introduction)
		if str.Username == ""{
			break
		}
		m++
		user[a-1] = str
	}
	return user, m
}

//关注用户
func (od *OperateDao) FocusUser (username, focusname string) error{
	_,err := od.Exec("update userfocus set focusname=? where username = ?",focusname,username )
	return err
}