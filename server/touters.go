package server

import "github.com/gin-gonic/gin"

func (s *Server) Routers() *gin.Engine {
	router := gin.Default()
	// 前端项目静态资源
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.StaticFile("/404.html", "./static/404.html")
	router.Static("/js/", "./static/js")
	router.Static("/plugins/", "./static/plugins")
	router.Static("/css/", "./static/css")
	router.Static("/src/", "./static/src")
	// 注册 api 分组路由
	group := router.Group("api")
	group.GET("getLogicalDrives", s.GetLogicalDrives)
	group.GET("getAllFile", s.GetAllFile)
	group.GET("readFile", s.ReadFile)
	group.GET("lookVideo", s.LookVideo)
	group.GET("view", s.View)
	return router
}
