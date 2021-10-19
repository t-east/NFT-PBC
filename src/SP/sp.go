package main

import (
	"github.com/gin-gonic/gin"
	"pairing_test/src/SP/structure"
	"pairing_test/src/SP/handler"
)

func main() {
	router := gin.Default()
	storage := structure.NewStorage()
	para := structure.NewParams()
	handler.GetPara(para)
	router.POST("/art", handler.ArtPost(storage))
	router.POST("/challenge", handler.AuditProofGen(para, storage))
	router.POST("/user", handler.UserPost(user))
	router.GET("/user", handler.UserGet(user))
	router.Run(":4000")
  }