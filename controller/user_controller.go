package controller

import (
	"biscuits-forum/dao"
	"biscuits-forum/model"
	"biscuits-forum/service"
	"biscuits-forum/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

type UserController struct {

}

func (uc *UserController) Router (engine *gin.Engine) {
	engine.POST("/user/register",uc.registe)
	engine.GET("/user/token",uc.login)
	engine.PUT("/user/password",uc.changePwd)
	engine.GET("/user/info/{user_id}",uc.getUserInfo)
	engine.PUT("/user/info",uc.changeUserInfo)
}

//用户注册
func (uc *UserController) registe (context *gin.Context){

	userName := context.PostForm("username")
	password := context.PostForm("password")
	us := new(service.UserService)

	flag := us.CheckUserAlive(userName)
	if flag {
		tool.PrintInfo(context, "该用户已存在")
		return
	}

	unlen := us.CheckUsernameLen(userName)
	if !unlen{
		tool.PrintInfo(context,"用户名不符合规范")
		return
	}

	pwdlen := us.CheckPwdLen(password)
	if !pwdlen{
		tool.PrintInfo(context,"密码不符合规范")
		return
	}

	ok := us.RegisteByPwd(userName, password)
	if !ok {
		tool.PrintInfo(context,"注册失败！")
		return
	}

	okk := us.RegisteUserInfo(userName)
	if !okk {
		tool.PrintInfo(context,"注册失败！")
		return
	}

	tool.PrintInfo(context,"注册成功!")

	fmt.Println("RegisteUserInfo:", userName, password)
}

//用户登录
func (uc *UserController) login (context *gin.Context){

	userName := context.Query("username")
	password := context.Query("password")
	us := new(service.UserService)
	fmt.Println("LoginUserInfo:", userName, password)
	flag := us.CheckUserAlive(userName)
	if !flag {
		tool.PrintInfo(context, "该用户不存在")
		return
	}

 	su := us.LoginByPwd(userName, password)
	if !su {
		tool.PrintInfo(context,"密码错误！")
		return
	}

	//创建一个token有效期两分钟
	tokenString, err := tool.GetToken(userName, "TOKEN", 1200)
	if err != nil {
		fmt.Println("CreateTokenErr:", err)
		tool.PrintInfo(context, "系统错误")
		log.Fatal(err.Error())
		return
	}

	//创建一个REFRESH_TOKEN有效期一周
	refreshToken, err := tool.GetToken(userName,"REFRESH_TOKEN", 604800)
	if err != nil {
		fmt.Println("CreateRefreshTokenErr:", err)
		tool.PrintInfo(context, "系统错误")
		return
	}

	context.JSON(200, gin.H{
		"status":       true,
		"data":         userName,
		"token":        tokenString,
		"refreshToken": refreshToken,
	})


}

//刷新token
func (uc *UserController) getToken(context *gin.Context) {
	refreshToken := context.Query("refreshToken")
	username := context.PostForm("username")


	//判断refreshToken状态
	model, err := tool.ParseRefreshToken(refreshToken)
	if err != nil {
		if err.Error()[:16] == "token is expired" {
			tool.Failed(context, "refreshToken失效")
			return
		}

		fmt.Println("getTokenParseTokenErr:", err)
		tool.Failed(context, "refreshToken不正确或系统错误")
		return

	}

	if model.Type == "ERR" {
		tool.Failed(context, "refreshToken不正确或系统错误")
		return
	}

	//创建新token
	newToken, err := tool.GetToken(username, "TOKEN", 1200)
	if err != nil {
		fmt.Println("getTokenCreateErr:", err)
		tool.Failed(context, "系统错误")
		return
	}

	tool.PrintInfo(context, newToken)
}

//修改密码
func (uc *UserController) changePwd (context *gin.Context){

	userName := context.PostForm("username")
	oldPassword := context.PostForm("oldPassword")
	newPassword := context.PostForm("newPassword")
	token := context.PostForm("token")
	us := new(service.UserService)

	fmt.Println(token)
	claims, err := tool.ParseToken(token)
	flag1 := tool.CheckTokenErr(context, claims, err)
	if !flag1  {
		return
	}

	su := us.LoginByPwd(userName, oldPassword)
	if !su{
		tool.PrintInfo(context,"密码错误！")
		return
	}

	pwdlen := us.CheckPwdLen(newPassword)
	if !pwdlen{
		tool.PrintInfo(context,"密码不符合规范")
		return
	}

	ok := us.ChangePwd(userName, newPassword)
	if !ok {
		tool.PrintInfo(context,"修改失败！")
		return
	}
	tool.PrintInfo(context,"修改成功！")
}

//获取用户信息
func (uc *UserController) getUserInfo (context *gin.Context) {

	userName := context.Query("username")
	us := new(service.UserService)
	dao := dao.UserDao{tool.GetDb()}
	var userInfo model.UserInfo
	token := context.Query("token")

	claims, err := tool.ParseToken(token)
	flag1 := tool.CheckTokenErr(context, claims, err)
	if flag1 == false {
		return
	}

	flag := us.CheckUserAlive(userName)
	if !flag {
		tool.PrintInfo(context, "该用户不存在")
		return
	}

	userInfo = dao.QueryUserInfo(userName)
	context.JSON(200,gin.H{
		"status":       10000,
		"info":        "success",
		"username":     userInfo.Username,
		"nickname":     userInfo.Nickname,
		"introduction": userInfo.Introduction,
		"phone":        userInfo.Telephone,
		"qq":           userInfo.QQ,
		"gender":       userInfo.Gender,
		"email":        userInfo.Email,
		"birthday":     userInfo.Birthday,
		"avatar": 		userInfo.Avatar,
	})
}

//更改用户信息
func (uc *UserController) changeUserInfo (context *gin.Context) {

	var userInfo model.UserInfo
	us := new(service.UserService)
	token := context.PostForm("token")

	err := context.ShouldBind(&userInfo)
	if err != nil {
		tool.PrintInfo(context,"参数获取失败！")
		return
	}

	claims, err := tool.ParseToken(token)
	flag1 := tool.CheckTokenErr(context, claims, err)
	if flag1 == false {
		return
	}

	flag := us.CheckUserAlive(userInfo.Username)
	if !flag {
		tool.PrintInfo(context, "该用户未存在")
		return
	}

	file, err1 := context.FormFile("avatar")
	if err1 != nil{
		tool.Failed(context,"参数解析失败")
	}

	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err2 := context.SaveUploadedFile(file, fileName)
	if err2 != nil {
		tool.Failed(context,"头像返回失败")
		return
	}

	path := us.UploadAvatar(userInfo)
	if !path {
		tool.Failed(context,"上传失败")
	}

    ok := us.CheckUserInfo(userInfo)
	if ok != "ok"{
		tool.PrintInfo(context, ok)
		return
	}

	ts := us.ChanegeUserInfo(userInfo)
	if !ts {
		tool.PrintInfo(context,"更改信息失败")
		return
	}
	tool.PrintInfo(context, "更改信息成功")
}

//

//验证
func (uc *UserController) sendCode (context *gin.Context){

	us := service.UserService{}
	code := us.Sendcode()
	if code == "验证码获取失败" {
		context.JSON(200,gin.H{
			"错误": code,
		})
	}else{
		context.JSON(200,gin.H{
			"验证码": code,
		})
	}
}