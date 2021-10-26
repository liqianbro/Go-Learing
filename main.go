package main

import (
	"Go-Learing/app/router"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			for i := 3; ; i++ {
				pc, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				fmt.Println(pc, file, line)
			}
		}
	}()
	// 设置为发布模式
	gin.SetMode(gin.DebugMode)

	r := router.SetupRouter()

	err := r.Run(":8881")
	if err != nil {
		panic(err)
	}
}
