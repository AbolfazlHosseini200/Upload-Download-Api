package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"sina_project/DataBaseController"
	"sina_project/Models"
	"sina_project/UploadDownloadController"
)

func main() {
	r := gin.Default()
    DataBaseController.Start()
	r.POST("/login", func(c *gin.Context){
		var user Models.Users
		c.BindJSON(&user)
        DataBaseController.Login(c,user)
	})
	r.POST("/register", func(c *gin.Context) {
		var user Models.Users
		c.BindJSON(&user)
		DataBaseController.Register(c,user)
	})
	r.POST("/upload", func(c *gin.Context) {
		UploadDownloadController.Upload(c)
	})
	r.POST("/download", func(c *gin.Context) {
		UploadDownloadController.Download(c)
		fmt.Println("heeb da")
	})
    r.POST("/addPermission",func(c *gin.Context){
    	DataBaseController.AddPermisssion(c.PostForm("username"),c.PostForm("fileid"))
	})
	r.Run(":8585")
}

