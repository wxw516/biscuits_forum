package service

import (
	"biscuits-forum/dao"
	"biscuits-forum/model"
	"biscuits-forum/tool"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type UserService struct {

}

//获取并存储验证码
func (us *UserService) Sendcode () string {
	//产生验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000))
	icode := model.Code{code,time.Now().Unix()}
	userDao := dao.UserDao{tool.Db}
	result := userDao.InsertCode(icode)
	if result > 0{
		return code
	}else{
		return "验证码获取失败"
	}
}

//用户注册
func (us *UserService) RegisteByPwd (userName, pwd string) bool{
	thisDao := dao.UserDao{tool.GetDb()}
	err := thisDao.InsertUser(userName, pwd)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}

//生成用户信息
func (us *UserService) RegisteUserInfo (username string) bool{
	thisDao := dao.UserDao{tool.GetDb()}
	err := thisDao.InsertUserInfo(username)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}

//检查用户是否存在
func (us *UserService) CheckUserAlive (username string) bool {
	thisDao := dao.UserDao{tool.GetDb()}
	return thisDao.QueryUsername(username)
}

//检查用户名是否符合规矩
func (us *UserService) CheckUsernameLen (username string) bool {
	if len(username) > 10||len(username) < 1{
		return false
	}
	return true
}

//检查密码是否符合规矩
func (us *UserService) CheckPwdLen (pwd string) bool {
	if len(pwd) > 20||len(pwd) < 6{
		return false
	}
	return true
}

//检查密码是否正确
func (us *UserService) LoginByPwd (username, pwd string) bool{
	thisDao := dao.UserDao{tool.GetDb()}

	rightPwd := thisDao.QueryUserPwd(username)
	if rightPwd != pwd {
		return false
	}
	return true
}

//修改密码
func (us *UserService) ChangePwd (username, pwd string) bool{
	thisDao := dao.UserDao{tool.GetDb()}

	err := thisDao.ChanegePwd(username, pwd)
	if err != nil{
		log.Fatal(err.Error())
		return false
	}
	return true
}

//判断用户信息是否符合规范
func (us *UserService) CheckUserInfo (userInfo model.UserInfo) string{

	var a [7]string
	a[0] = us.CheckNickname(userInfo.Nickname)
	a[1] = us.CheckIntroduction(userInfo.Introduction)
	a[2] = us.CheckTelephone(userInfo.Telephone)
	a[3] = us.CheckQQ(userInfo.QQ)
	a[4] = us.CheckEmail(userInfo.Email)
	a[5] = us.CheckBirthday(userInfo.Birthday)
	a[6] = us.CheckGender(userInfo.Gender)
	for i := 0;i < 7;i++{
		if a[i] != "ok"{
			return a[i]
		}
	}
	return "ok"
}

//判断昵称是否符合规范
func (us *UserService) CheckNickname (nickname string) string{
	if len(nickname) > 32||len(nickname) < 1{
		return "用户昵称不符合规范"
	}
	return "ok"
}

//判断简介是否符合规范
func (us *UserService) CheckIntroduction (introduction string) string{
	if len(introduction) > 100{
		return "用户简介不符合规范"
	}
	return "ok"
}

//判断电话号是否符合规范
func (us *UserService) CheckTelephone (telephone string) string{
	if len(telephone) != 11{
		return "用户电话号码不符合规范"
	}
	return "ok"
}

//检查qq是否符合规范
func (us *UserService) CheckQQ (qq string) string{
	if len(qq) > 10{
		return "用户QQ不符合规范"
	}
	return "ok"
}

//判断性别是否符合规范
func (us *UserService) CheckGender (gender string) string{
	if gender == "男"||gender == "女"||gender == "双性人" {
		return "ok"
	}
	return "用户性别不符合规范"
}

//检查邮箱是否符合规范
func (us *UserService) CheckEmail (email string) string{
	if len(email) > 20{
		return "用户邮箱不符合规范"
	}
	return "ok"
}

//检查生日是否符合规范
func (us *UserService) CheckBirthday (birthday string) string{

	return "ok"
}

//更改用户信息
func (us *UserService) ChanegeUserInfo (userInfo model.UserInfo) bool{
	dao := dao.UserDao{tool.GetDb()}

	err1 := dao.ChangeUserNickname(userInfo)
	if err1 != nil {
		log.Fatal(err1.Error())
		return false
	}

	err2 := dao.ChangeUserIntroduction(userInfo)
	if err2 != nil {
		log.Fatal(err2.Error())
		return false
	}

	err3 := dao.ChangeUserTelephone(userInfo)
	if err3 != nil {
		log.Fatal(err3.Error())
		return false
	}

	err4 := dao.ChangeUserQQ(userInfo)
	if err4 != nil {
		log.Fatal(err4.Error())
		return false
	}

	err5 := dao.ChangeUserEmail(userInfo)
	if err5 != nil {
		log.Fatal(err5.Error())
		return false
	}

	err6 := dao.ChangeUserGender(userInfo)
	if err6 != nil {
		log.Fatal(err6.Error())
		return false
	}

	err7 := dao.ChangeUserBirthday(userInfo)
	if err7 != nil {
		log.Fatal(err7.Error())
		return false
	}

	err8 := dao.ChangeUserAvatar(userInfo)
	if err8 != nil {
		log.Fatal(err8.Error())
		return false
	}

	return true
}

//上传头像
func (us UserService) UploadAvatar(userinfo model.UserInfo) bool{
	userDao := dao.UserDao{tool.GetDb()}
	err := userDao.ChangeUserAvatar(userinfo)
	if err != nil {
		return false
	}
	return true
}