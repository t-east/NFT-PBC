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
	router.GET("/audit", handler.AuditVerify(para))
	router.GET("/count", handler.Count())
	router.GET("/chal", handler.AuditChallen(para))
	router.GET("/audit-logs", handler.GetAuditLogs())
	router.Run(":4002")
  }