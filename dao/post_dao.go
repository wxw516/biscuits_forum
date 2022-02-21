package dao

import (
	"biscuits-forum/model"
	"database/sql"
	"fmt"
	"log"
)

type PostDao struct {
	*sql.DB
}

//获取帖子
func (pd *PostDao) GetPost (postId string) model.PostInfo{
	row := pd.QueryRow("select is_praised, is_focus, comment_count, praise_count, title, topic_id, publish_time, content, pictures, username, nickname, avatar from postinfo where post_id = ? ", postId)

	err := row.Err()
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	var PI model.PostInfo
	PI.PostId = postId
	row.Scan(&PI.IsPraised, &PI.IsFocus, &PI.CommentCount, &PI.PraiseCount, &PI.Title, &PI.PublishTime, &PI.Content, &PI.Pictures, &PI.UseName, &PI.Nickname, &PI.Avatar)
	return PI
}

//查找帖子是否存在
func (pd *PostDao) CheckPostAlive (postId string) bool{
	row := pd.QueryRow("select post_id from postinfo where post_id = ? ", postId)
	err := row.Err()
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}

	var user string
	row.Scan(&user)
	if user == "" {
		return false
	}
	fmt.Println(user)
	return true
}

//查找话题是否存在
func (pd *PostDao) CheckTopicAlive (topicId string) bool{
	row := pd.QueryRow("select topic_id from postinfo where topic_id = ? ", topicId)
	err := row.Err()
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}

	var user string
	row.Scan(&user)
	if user == "" {
		return false
	}
	fmt.Println(user)
	return true
}

//存储帖子
func (pd *PostDao) InsertPost (PI model.PostInfo) error{
	_,err := pd.Exec("insert into postinfo(is_praised, is_focus, comment_count, praise_count, title, topic_id, publish_time, content, pictures, username, nickname, avatar) values (?,?,?,?,?,?,?,?,?,?,?,?)", PI.IsPraised, PI.IsFocus, PI.CommentCount, PI.PraiseCount, PI.Title, PI.PublishTime, PI.Content, PI.Pictures, PI.UseName, PI.Nickname, PI.Avatar)
	return err
}

//更新帖子
func (pd *PostDao) ChangePost (PI model.PostInfo) error{
	_,err := pd.Exec("update postinfo set title = ?,content = ?, topic_id = ?, photo + ? where username = ?",PI.Title, PI.Content, PI.TopicId, PI.Pictures)
	return err
}

//删除帖子
func (pd *PostDao) DeletePost (postId string) error{
	_,err := pd.Exec("delete from postinfo where post_id = ?", postId)
	return err
}

//获取帖子主题
func (pd *PostDao) GetPostTopic () ([]string, int){
	var topic []string
	m := 0
	for a := 1; ; a++{
		row := pd.QueryRow("select topic_id from postinfo where post_id = ? ", a)

		err := row.Err()
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		var str string
		row.Scan(&str)
		if str == ""{
			break
		}
		m++
		topic[a-1] = str
	}
	return topic, m
}