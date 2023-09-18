package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (s *Server) View(c *gin.Context) {
	pathAny, _ := c.Get("path")
	pathStr := pathAny.(string)
	fileType := strings.ToLower(filepath.Ext(pathStr))

	//后缀是pdf直接读取文件类容返回
	if fileType == ".txt" {
		c.Header("Content-Type", "application/pdf;charset=UTF-8")
		file, err := os.ReadFile(pathStr)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Data(http.StatusOK, "text/plain;charset=UTF-8", file)
		return
	}

	if fileType == ".pdf" {
		c.Header("Content-Type", "application/pdf;charset=UTF-8")
		file, err := os.ReadFile(pathStr)
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
