package model

type PostInfo struct {
	PostId string 		`from:"post_id"`
	IsPraised bool		`from:"is_praised"`
	IsFocus bool		`from:"is_focus"`
	CommentCount int	`from:"comment_count"`
	PraiseCount int		`from:"praise_count"`
	Title string 		`from:"title"`
	TopicId string		`from:"topic_id"`
	PublishTime int		`from:"publish_time"`
	Content string		`from:"content"`
	Pictures []string	`from:"pictures"`
	UseName string		`from:"username"`
	Nickname string		`from:"nickname"`
	Avatar string		`from:"avatar"`
}

type Post struct {
	PostId string 		`from:"post_id"`
	Title string 		`from:"title"`
	PublishTime int		`from:"publish_time"`
	UseName string		`from:"username"`
	Nickname string		`from:"nickname"`
}