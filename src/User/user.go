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
	para := structure.NewParams()
	uploadFile := structure.File()
	artIds := &structure.ArtIds{}
	router.GET("/params", handler.GetPara(para))
	router.GET("/keygen", handler.KeyGen(para, user))
	router.GET("/user/address", handler.GetAddress())
	router.GET("/user/publickey", handler.GetPubkey())
	router.POST("/user", handler.UserPost(user))
	router.GET("/user", handler.UserGet(user))
	router.GET("/user/register", handler.Register(user))
	router.POST("/file/metadata", handler.CreateMetaData(uploadFile, para, user))
	router.GET("/file/upload", handler.UploadFile(artIds, uploadFile, para, user))
	// router.GET("/file", handler.ArtGet())
	router.Run(":4000")
  }
