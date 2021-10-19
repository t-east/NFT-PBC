package main

import (
	"github.com/gin-gonic/gin"
	"pairing_test/src/User/structure"
	"pairing_test/src/User/handler"
)

func main() {
	router := gin.Default()
	para := structure.NewParams()
	handler.GetPara(para)
	router.POST("/verify", handler.AuditVerify())
	router.Run(":4000")
  }