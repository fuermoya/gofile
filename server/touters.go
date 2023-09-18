package server

import (
	"embed"
	_ "embed"
	middleware "github.com/fuermoya/gofile/midd"
	"github.com/gin-gonic/gin"
	fs2 "io/fs"
	"net/http"
	"strconv"
)

func (s *Server) Routers(multi embed.FS) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.DecodePath())
	router.Use(middleware.IsLocalIp())

	fs := http.FS(multi)
	indexHtml, _ := multi.ReadFile("static/index.html")
	router.GET("/", func(c *gin.Context) {
		c.Header("Accept-Ranges", "bytes")
		c.Header("Content-Length", strconv.Itoa(len(indexHtml)))
		c.Data(http.StatusOK, "text/html;charset=UTF-8", indexHtml)
	})

	router.StaticFileFS("favicon.ico", "./static/favicon.ico", fs)
	router.StaticFileFS("404.html", "./static/404.html", fs)

	js, _ := fs2.Sub(multi, "static/js")
	plugins, _ := fs2.Sub(multi, "static/plugins")
	css, _ := fs2.Sub(multi, "static/css")
	src, _ := fs2.Sub(multi, "static/src")
	router.StaticFS("/js/", http.FS(js))
	router.StaticFS("/plugins/", http.FS(plugins))
	router.StaticFS("/css/", http.FS(css))
	router.StaticFS("/src/", http.FS(src))

	// 前端项目静态资源 方便测试
	//router.StaticFile("/", "./static/index.html")
	//router.StaticFile("/favicon.ico", "./static/favicon.ico")
	//router.StaticFile("/404.html", "./static/404.html")
	//router.Static("/js/", "./static/js")
	//router.Static("/src/", "./static/src")
	//router.Static("/plugins/", "./static/plugins")
	//router.Static("/css/", "./static/css")

	// 注册 api 分组路由
	group := router.Group("api")
	group.GET("getLogicalDrives", s.GetLogicalDrives)
	group.GET("getAllFile", s.GetAllFile)
	group.GET("readFile", s.ReadFile)
	group.GET("downloadFile", s.DownloadFile)
	group.GET("lookVideo", s.LookVideo)
	group.GET("getSubtitle", s.GetSubtitle)
	group.GET("view", s.View)
	return router
}
