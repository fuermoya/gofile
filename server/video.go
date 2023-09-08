package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// LookVideo TODO  不支持的视频需要转码，还有字幕
func (s *Server) LookVideo(c *gin.Context) {
	pathname := c.Query("path")
	file, err := os.Open(pathname)
	defer file.Close()
	if err != nil {
		c.Redirect(http.StatusFound, "./static/404.html")
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+file.Name())
	c.File(pathname)
}
