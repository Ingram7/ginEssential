package controller

import (
	"ginessential/dao"
	"ginessential/models"
	"ginessential/response"
	"ginessential/vo"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := dao.GetDB()
	db.AutoMigrate(models.Category{})
	return PostController{DB: db}
}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	if err := ctx.ShouldBind(&requestPost); err != nil {
		response.Fail(ctx, nil, "数据验证错误")
		return
	}
	user, _ := ctx.Get("user")
	post := models.Post{
		UserId:     user.(models.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}
	response.Success(ctx, nil, "创建成功")

}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	if err := ctx.ShouldBind(&requestPost); err != nil {
		response.Fail(ctx, nil, "数据验证错误")
		return
	}
	postId := ctx.Params.ByName("id")
	var post models.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "文章不存在")
		return
	}
	user, _ := ctx.Get("user")
	userId := user.(models.User).ID
	if userId != post.UserId {
		response.Fail(ctx, nil, "请勿非法操作")
		return
	}
	if err := p.DB.Model(&post).Update(requestPost).Error; err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}
	response.Success(ctx, gin.H{"post": post}, "更新成功")
}

func (p PostController) Show(ctx *gin.Context) {
	postId := ctx.Params.ByName("id")
	var post models.Post
	if p.DB.Preload("category").Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "文章不存在")
		return
	}
	response.Success(ctx, gin.H{"post": post}, "成功")
}

func (p PostController) Delete(ctx *gin.Context) {
	postId := ctx.Params.ByName("id")
	var post models.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "文章不存在")
		return
	}
	user, _ := ctx.Get("user")
	userId := user.(models.User).ID
	if userId != post.UserId {
		response.Fail(ctx, nil, "请勿非法操作")
		return
	}

	if p.DB.Delete(&post).Error != nil {
		response.Fail(ctx, nil, "删除失败 ")
		return
	}
	response.Success(ctx, nil, "删除成功")
}

func (p PostController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))
	// 分页
	var posts []models.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)
	// 总数
	var total int
	p.DB.Model(models.Post{}).Count(&total)
	response.Success(ctx, gin.H{"data": posts, "total": total}, "成功")
}
