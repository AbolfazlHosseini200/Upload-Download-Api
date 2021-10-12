package UploadDownloadController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"sina_project/DataBaseController"
	"sina_project/Models"
)

func Upload(c *gin.Context) {
	var fileInfo Models.Files
	file, err := c.FormFile("file")
	fileInfo.FileId= c.PostForm("fileid")
	fileInfo.Username=c.PostForm("username")
	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	if !DataBaseController.IsUnique(fileInfo.FileId,c){
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := fileInfo.FileId + extension
    DataBaseController.AddPermisssion(fileInfo.Username,newFileName)
	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "C:\\Users\\AmirHossein\\Desktop\\Coding Stuff\\Code Go\\project\\UploadedFiles\\"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}
