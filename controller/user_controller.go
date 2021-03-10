package controller

import (
	"ginessential/common"
	"ginessential/dao"
	"ginessential/models"
	"ginessential/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @title 注册
// @description 用户注册
// @tags 注册
// @accept mpfd
// @produce json
// @param name formData string false "用户名"
// @param telephone formData string true "登录账户"
// @param password formData string true "登录密码"
// @Success 200  {string} string "{"code": "200", "msg": "具体信息", "data": "具体数据"}"
// @Failure 400  {string} string "{"code": "400", "msg": "具体信息", "data": "具体数据"}"
// @router /api/auth/register [post]
func Register(c *gin.Context) {
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 校验手机位数
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	// 校验密码位数
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 如果没传name， 置于一个10位随机字符串
	if len(name) == 0 {
		name = common.RandomString(10)
	}
	// 判断手机号是否存在
	if common.IsTelephoneExist(telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "该用户已经存在")
		return
	}
	// 创建用户

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	dao.CreateUser(&newUser)
	response.Success(c, nil, "注册成功")
}

// @title 登录
// @description 用户登录
// @tags 登录
// @accept mpfd
// @produce json
// @param telephone formData string true "登录账户"
// @param password formData string true "登录密码"
// @Success 200 {string} string "{"code": 200, "msg": "具体信息", "data": "具体数据"}"
// @Failure 400 {string} string "{"code": 400, "msg": "具体信息", "data": "具体数据"}"
// @router /api/auth/login [post]
func Login(c *gin.Context) {
	// 获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	// 校验密码位数
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 判断手机号是否存在
	var user models.User
	dao.GetTelephone(&user, telephone)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(c, nil, "密码错误")
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v\n", err)
		return
	}
	// 返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")
}

// @title 用户信息
// @description 用户信息
// @tags 用户信息

// @produce json

// @Success 200 {string} string "{"code": 200, "msg": "具体信息", "data": "具体数据"}"
// @Failure 400 {string} string "{"code": 400, "msg": "具体信息", "data": "具体数据"}"
// @router /api/auth/info [get]
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, gin.H{"user": common.ToUserDto(user.(models.User))}, "")
}
