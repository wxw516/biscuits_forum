package dao

import (
	"biscuits-forum/model"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type UserDao struct {
	*sql.DB
}

func (ud *UserDao) InsertCode (code model.Code) int{
	sqlStr := "insert into user_code(code, time) values (?,?)"
	_, err := ud.Exec(sqlStr,&code)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}else{
		return 1
	}
}

//插入注册信息
func (dao *UserDao) InsertUser (username, pwd string) error {
	registerTime := time.Now().Unix()
	_,err := dao.Exec("insert into userid(username, password, registertime) values (?,?,?)", username, pwd, registerTime)
	return err
}

//注册时插入用户信息
func (dao *UserDao) InsertUserInfo (username string) error{
	_,err := dao.Exec("insert into userinfo(username) values (?)", username)
	return err
}

//查询用户是否存在
func (dao *UserDao) QueryUsername(name string) bool {
	row := dao.QueryRow("select username from userid where username = ? ", name)
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

//返回密码
func (dao *UserDao) QueryUserPwd (username string) string {
	row := dao.QueryRow("select password from userid where username = ? ", username)

	err := row.Err()
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	var pwd string
	row.Scan(&pwd)
	return pwd
}

//更改密码
func (dao *UserDao) ChanegePwd (username, pwd string) error{
	_,err := dao.Exec("update userid set password=? where username = ?", pwd, username)
	return err
}

//更改用户信息
func (dao *UserDao) ChangeUserNickname (info model.UserInfo) error{
	_,err := dao.Exec("update userinfo set nickname=? where username = ?", info.Nickname, info.Username)
	return err
}

func (dao *UserDao) ChangeUserIntroduction (info model.UserInfo) error{
	_,err := dao.Exec("update userinfo set introduction=? where username = ?", info.Introduction, info.Username)
	return err
}

func (dao *UserDao) ChangeUserTelephone (info model.UserInfo) error{
	_,err := dao.Exec("update userinfo set telephone=? where username = ?", info.Telephone, info.Username)
	return err
}

func (dao *UserDao) ChangeUserQQ (info model.UserInfo) error{
	_,err := dao.Exec("update userinfo set qq=? where username = ?", info.QQ, info.Username)
	return err
}

func (dao *UserDao) ChangeUserEmail (info model.UserInfo) error{
	_,err := dao.Exec("update userinfo set email=? where username = ?", info.Email, info.Username)
	return err
}

func (dao *UserDao) ChangeUserGender (info model.UserInfo) error{
	_,err := dao.Exec("update userinfo set gender=? where username = ?", info.Gender, info.Username)
	return err
}

func (dao *UserDao) ChangeUserBirthday (info model.UserInfo) error{
	_,err := dao.Exec("update userinfo set birthday=? where username = ?", info.Birthday, info.Username)
	return err
}

func (dao *UserDao) ChangeUserAvatar (info model.UserInfo) error{
	_,err := dao.Exec("update userinfo set avatar=? where username = ?", info.Avatar, info.Username)
	return err
}

//查询用户信息
func (dao *UserDao) QueryUserInfo (username string) model.UserInfo {
	row := dao.QueryRow("select nickname, introduction, telephone, qq, gender, email, birthday from userinfo where username = ? ", username)

	err := row.Err()
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	var UI model.UserInfo
	UI.Username = username
	row.Scan(&UI.Nickname, UI.Introduction, UI.Telephone, UI.QQ, UI.Gender, UI.Email, UI.Birthday)
	return UI
}

