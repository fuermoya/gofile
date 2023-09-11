package server

import (
	"github.com/gin-gonic/gin"
	response "gofile/common"
	"gofile/utils"
	"os"
)

var exts = [...]string{".vtt", ".ass"}

// LookVideo TODO  不支持的视频需要转码为mp4 字幕格式转换为vtt
func (s *Server) LookVideo(c *gin.Context) {
	pathAny, _ := c.Get("path")
	filePath := pathAny.(string)
	c.File(filePath)
}

// GetSubtitle 获取字幕，与视频同名
func (s *Server) GetSubtitle(c *gin.Context) {
	pathAny, _ := c.Get("path")
	filePath := pathAny.(string)
	var path string
	for _, v := range exts {
		ext := utils.JoinExt(filePath, v)
		_, err := os.Stat(ext)
		if err == nil || os.IsExist(err) {
			path = ext
			response.OkWithData(gin.H{"path": path, "ext": v}, c)
			return
		}
	}
	response.OkWithData(gin.H{}, c)
}
