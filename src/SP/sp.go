package main

import (
	"github.com/gin-gonic/gin"
	"pairing_test/src/SP/structure"
	"pairing_test/src/SP/handler"
)

func main() {
	router := gin.Default()
	storage := &structure.Storage{}
	para := structure.NewParams()
	handler.GetPara(para)
	router.POST("/art", handler.ArtPost(storage))
	router.GET("/art/:id", handler.ArtGet())
	router.POST("/proof", handler.AuditProofGen(para, storage))
	router.Run(":4001")
  }