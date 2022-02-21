package model

type UserId struct {
	Message string `form:"username"`
	Username string `form:"password"`
}

type UserInfo struct {
	Username string     `form:"username"`
	Nickname string	    `form:"nickname"`
	Avatar string		`from:"avatar"`
	Introduction string `form:"introduction"`
	Telephone string    `form:"telephone"`
	QQ string           `form:"qq"`
	Gender string       `form:"gender"`
	Email string		`form:"email"`
	Birthday string		`form:"birthday"`
}

type User struct {
	Username string `from:"username"`
	Nickname string	    `form:"nickname"`
	Introduction string `form:"introduction"`
}

