package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	"pairing_test/src/User/structure"
	"pairing_test/src/User/handler"
)

func main() {
	router := gin.Default()
	user := structure.NewUser()
	datas := structure.Datas()
	para := structure.NewParams()
	router.GET("/params", handler.GetPara(para))
	router.POST("/upload", handler.FilePost(datas))
	router.GET("/keygen", handler.KeyGen(para, user))
	router.POST("/user/address", handler.GetAddress(user))
	router.POST("/user", handler.UserPost(user))
	router.GET("/user", handler.UserGet(user))
	router.Run(":4000")
  }
