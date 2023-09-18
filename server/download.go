package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func (s *Server) DownloadFile(c *gin.Context) {
	pathAny, _ := c.Get("path")
	filePath := pathAny.(string)
	file, err := os.Open(filePath)
	file.Close()
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	c.File(filePath)
}
