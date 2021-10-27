package main

import (
	"Go-Learing/app/db/mysqldb"
	"Go-Learing/app/router"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	// 异常捕获
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
	// 初始化 viper
	err := mysqldb.InitViper()
	if err != nil {
		return
	}
	// 初始化 MySQL
	err = mysqldb.InitMySQL()
	if err != nil {
		return
	}
	// 设置为发布模式
	gin.SetMode(gin.DebugMode)
	r := router.SetupRouter()

	err = r.Run(":8881")
	if err != nil {
		panic(err)
	}
}
