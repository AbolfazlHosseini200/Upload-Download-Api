package UploadDownloadController

import (
	"github.com/gin-gonic/gin"
	"sina_project/DataBaseController"
	_ "sina_project/DataBaseController"
)

func Download(c *gin.Context){
	fileid:=c.PostForm("fileid")
	username:=c.PostForm("username")
	path:="C:\\Users\\AmirHossein\\Desktop\\Coding Stuff\\Code Go\\project\\UploadedFiles\\"+fileid
if !DataBaseController.IsPermited(username,fileid,c){
	return
}
	c.File(path)
}
