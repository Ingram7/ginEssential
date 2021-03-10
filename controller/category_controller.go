package controller

import (
	"ginessential/models"
	"ginessential/repository"
	"ginessential/response"
	"ginessential/vo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	rep := repository.NewCategoryRepository()
	rep.DB.AutoMigrate(models.Category{})
	return CategoryController{Repository: rep}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误， 分类名称必填")
		return
	}
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		response.Fail(ctx, nil, "创建失败")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误， 分类名称必填")
		return
	}
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.Repository.DeleteById(categoryId); err != nil {
		response.Fail(ctx, nil, "删除失败, 请重试")
		return
	}
	response.Success(ctx, nil, "")
}
