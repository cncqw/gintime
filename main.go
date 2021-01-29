package main

import (
	"alldu.cn/ginproject/common"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB() //初始化
	//defer sqlDB.Close() //gorm v1.20已取消关闭方法

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}
