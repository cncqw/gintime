package controller

import (
	"alldu.cn/ginproject/common"
	"alldu.cn/ginproject/model"
	"alldu.cn/ginproject/response"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "分类名称必填")
		return
	}

	c.DB.Create(&requestCategory)

	response.Success(ctx, gin.H{"category": requestCategory}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	//绑定Body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "分类名称必填")
		return
	}

	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory model.Category

	err := c.DB.First(&updateCategory, categoryId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Fail(ctx, nil, "分类不存在")
	}

	//更新分类
	//map
	//struct
	//name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
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
