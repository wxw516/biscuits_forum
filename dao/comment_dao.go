package dao

import (
	"biscuits-forum/model"
	"database/sql"
	"log"
)

type CommentDao struct {
	*sql.DB
}

//获取一级评论
func (cd *CommentDao) GetPostComent (id string, page, size int) []model.CommentInfo{
	var cm []model.CommentInfo
  	var cn []model.CommentInfo
	n := (page - 1) * size

	row, err := cd.Query("select comment_id,post_id,publish_time,content,pictures,user_id,nickname,reply_user_id,praise_count,is_praise,is_focus from commentinfo where post_id= ?",id)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	for i := 0; i < n; i++{
		if row.Next(){
			err := row.Scan(&cm[i].CommentId,&cm[i].PostId,&cm[i].PublishTime,&cm[i].Content,&cm[i].Pictures,&cm[i].UserId,&cm[i].Nickname,&cm[i].ReplyUserId,&cm[i].PraiseCount,&cm[i].IsPraise,&cm[i].IsFocus)
			if err != nil {
				log.Fatal(err.Error())
				return nil
			}
		}
	}

	for i := 0; i < size; i++{
		if row.Next(){
			err := row.Scan(&cn[i].CommentId,&cn[i].PostId,&cn[i].PublishTime,&cn[i].Content,&cn[i].Pictures,&cn[i].UserId,&cn[i].Nickname,&cn[i].ReplyUserId,&cn[i].PraiseCount,&cn[i].IsPraise,&cn[i].IsFocus)
			if err != nil {
				log.Fatal(err.Error())
				return nil
			}
		}
	}

	return cn
}

//获取二级评论
func (cd *CommentDao) GetCommentComment (id string, page, size int) []model.CommentInfo{
	var cm []model.CommentInfo
	var cn []model.CommentInfo
	n := (page - 1) * size

	row, err := cd.Query("select comment_id,post_id,publish_time,content,pictures,user_id,nickname,reply_user_id,praise_count,is_praise,is_focus from commentinfo where comment_id = ?",id)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	for i := 0; i < n; i++{
		if row.Next(){
			err := row.Scan(&cm[i].CommentId,&cm[i].PostId,&cm[i].PublishTime,&cm[i].Content,&cm[i].Pictures,&cm[i].UserId,&cm[i].Nickname,&cm[i].ReplyUserId,&cm[i].PraiseCount,&cm[i].IsPraise,&cm[i].IsFocus)
			if err != nil {
				log.Fatal(err.Error())
				return nil
			}
		}
	}

	for i := 0; i < size; i++{
		if row.Next(){
			err := row.Scan(&cn[i].CommentId,&cn[i].PostId,&cn[i].PublishTime,&cn[i].Content,&cn[i].Pictures,&cn[i].UserId,&cn[i].Nickname,&cn[i].ReplyUserId,&cn[i].PraiseCount,&cn[i].IsPraise,&cn[i].IsFocus)
			if err != nil {
				log.Fatal(err.Error())
				return nil
			}
		}
	}

	return cn
}

//发布一级评论
func (cd *CommentDao) PostComment (id, content string) error{
	_,err := cd.Exec("insert into commentinfo(post_id, content) values (?,?)", id, content)
	return err
}

//发布二级评论
func (cd *CommentDao) PublishCommentComment (id, content string) error{
	_,err := cd.Exec("insert into commentinfo(comment_id, content) values (?,?)", id, content)
	return err
}

//更新评论
func (cd *CommentDao) ChangeComment (id, content string) error{
	_,err := cd.Exec("update commentinfo set content=? where commnent_id = ?", content, id)
	return err
}

//删除评论
func (cd *CommentDao) DeleteComent (id string) error{
	_,err := cd.Exec("delete from commentinfo where comment_id = ?", id)
	return err
}