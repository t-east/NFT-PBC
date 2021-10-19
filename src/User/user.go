package main

import (
	"github.com/gin-gonic/gin"
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
	router.POST("/keygen", handler.KeyGenPost(para, user))
	router.POST("/user", handler.UserPost(user))
	router.GET("/user", handler.UserGet(user))
	router.Run(":4000")
  }
