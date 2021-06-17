package controller

import (
	"alldu.cn/ginproject/model"
	"alldu.cn/ginproject/repository"
	"alldu.cn/ginproject/response"
	"alldu.cn/ginproject/vo"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	// DB *gorm.DB
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()

	// db := common.GetDB()
	repository.DB.AutoMigrate(model.Category{})

	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "分类名称必填")
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		response.Fail(ctx, nil, "创建失败")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "分类名称必填")
		return
	}

	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//var updateCategory model.Category
	//err := c.DB.First(&updateCategory, categoryId).Error
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	response.Fail(ctx, nil, "分类不存在")
	//}

	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	//更新分类
	//map
	//struct
	//name value
	category ,err := c.Repository.Update(*updateCategory, requestCategory.Name)
	//c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	if err != nil{
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "修改成功")
}


func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var category model.Category
	err := c.DB.First(&category, categoryId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	err := c.DB.Delete(model.Category{}, categoryId).Error
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}

	response.Success(ctx, nil, "删除成功")
}
