package router

import (
	"Go-Learing/app/db/mysqldb"
	"Go-Learing/app/protoc/cloud/protos"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/hello", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "ok", "data": "Hello,Gin!!!"})
		})
		api.GET("/pro", func(c *gin.Context) {
			ret := &protos.Hello{
				Name: "Gin",
			}
			ProtoResponse(c, ret)
		})
		api.GET("/get/user", func(context *gin.Context) {
			count, err := mysqldb.DB.User.Query().Count(context)
			if err != nil {
				return
			}
			context.JSON(http.StatusOK, gin.H{"data": count})
		})
	}

	return r
}

func ProtoResponse(c *gin.Context, data proto.Message) {
	if c.Query("format") == "json" {
		c.JSON(http.StatusOK, data)
	} else {
		c.ProtoBuf(http.StatusOK, data)
	}
}
