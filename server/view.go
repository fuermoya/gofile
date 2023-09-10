package server

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strings"
)

func (s *Server) View(c *gin.Context) {
	filePath := c.Query("path")
	decodeString, err := base64.StdEncoding.DecodeString(filePath)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	filePath = string(decodeString)
	fileType := strings.ToLower(path.Ext(filePath))

	//后缀是pdf直接读取文件类容返回
	if fileType == ".txt" {
		c.Header("Content-Type", "application/pdf;charset=UTF-8")
		file, err := os.ReadFile(filePath)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Data(http.StatusOK, "text/plain;charset=UTF-8", file)
		return
	}

	if fileType == ".pdf" {
		c.Header("Content-Type", "application/pdf;charset=UTF-8")
		file, err := os.ReadFile(filePath)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Data(http.StatusOK, "application/pdf", file)
		return
	}

	//暂时不支持
	c.Status(http.StatusNotFound)

}
