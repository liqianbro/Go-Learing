package router

import (
	"Go-Learing/app/protoc/cloud/protos"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
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
	}

	return r
}

func ProtoResponse(c *gin.Context, data proto.Message) {
	if c.Query("format") == "json" {
		b, _ := protojson.Marshal(data)
		c.String(http.StatusOK, "%s", b)
	} else {
		c.ProtoBuf(http.StatusOK, data)
	}
}