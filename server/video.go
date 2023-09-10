package server

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"os"
)

// LookVideo TODO  不支持的视频需要转码，还有字幕
func (s *Server) LookVideo(c *gin.Context) {
	filePath := c.Query("path")
	decodeString, err := base64.StdEncoding.DecodeString(filePath)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	filePath = string(decodeString)
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	c.File(filePath)
}
