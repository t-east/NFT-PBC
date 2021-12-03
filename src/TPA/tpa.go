package main

import (
	"github.com/gin-gonic/gin"
	"pairing_test/src/TPA/structure"
	"pairing_test/src/TPA/handler"
)

func main() {
	router := gin.Default()
	para := structure.NewParams()
	handler.GetPara(para)
	router.POST("/audit", handler.AuditVerify(para))
	router.GET("/count", handler.Count())
	router.GET("/chal", handler.AuditChallen(para))
	router.Run(":4002")
  }