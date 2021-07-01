package controller

import (
	"alldu.cn/ginproject/common"
	"alldu.cn/ginproject/model"
	"alldu.cn/ginproject/response"
	"alldu.cn/ginproject/vo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IPostController interface {
	RestController
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {

	db := common.GetDB()
	db.AutoMigrate(model.Post{})

	return PostController{DB: db}
}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		response.Fail(ctx, nil, "数据验证错误")
		return
	}
	// 获取登录用户
	user, _ := ctx.Get("user")

	//创建post
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	if err := p.DB.Create(&post).Error;err != nil{
		panic(err)
		return
	}

	response.Success(ctx,nil,"创建成功")
}


func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 获取登录用户
	user, _ := ctx.Get("user")

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Post
	if p.DB.Where("id = ?",postId).First(&post).RecordNotFound(){
		response.Fail(ctx,nil,"文章不存在")
	}
	panic("implement me")
}

func (p PostController) Show(ctx *gin.Context) {
	panic("implement me")
}

func (p PostController) Delete(ctx *gin.Context) {
	panic("implement me")
}
