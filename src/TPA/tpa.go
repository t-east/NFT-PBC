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
	var Challens &[]structure.Chal{}
	router.POST("/audit", handler.AuditVerify())
	router.POST("/chal", handler.PostChallens())
	router.Run(":4000")
  }