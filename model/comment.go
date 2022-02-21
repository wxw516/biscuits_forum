package model

type CommentInfo struct {
	CommentId string 			`form:"comment_id"`
	PostId string				`form:"post_id"`
	PublishTime int 			`form:"publish_time"`
	Content string 				`form:"content"`
	Pictures []string 			`form:"pictures"`
	UserId string 				`form:"user_id"`
	Avater string 				`form:"avater"`
	Nickname string 			`form:"nickname"`
	ReplyUserId string 			`form:"reply_user_id"`
	RelpyUserNickname string	`form:"reply_user_nickname"`
	PraiseCount string 			`form:"praise_count"`
	IsPraise string 			`form:"is_praise"`
	IsFocus string 				`form:"is_focus"`
}

