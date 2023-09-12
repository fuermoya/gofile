package middleware

import (
	"encoding/base64"
	"github.com/fuermoya/gofile/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DecodePath 解码path的base64
func DecodePath() gin.HandlerFunc {
	return func(c *gin.Context) {
		filePath := c.Query("path")
		decodeString, err := base64.StdEncoding.DecodeString(filePath)
		if err != nil {
			c.Status(http.StatusNotFound)
			c.Abort()
			return
		}
		c.Set("path", string(decodeString))
		c.Next()
	}
}

// IsLocalIp 校验是否是局域网
func IsLocalIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !utils.IsLocalIp(ip) {
			c.Status(http.StatusNotFound)
			c.Abort()
			return
		}
		c.Next()
	}
}
