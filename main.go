package main

import (
	"Go-Learing/app/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 设置为发布模式
	gin.SetMode(gin.DebugMode)

	r := router.SetupRouter()

	err := r.Run(":8881")
	if err != nil {
		panic(err)
	}
}
