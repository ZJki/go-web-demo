package controller

import (
	"go-web-practice/internal/config"
	"net/http"
)

// 注册所有路由的函数
func RegisterRoutes() {
	// 注册首页路由
	registerIndexRoutes()

	// 注册欢迎页面路由
	registerWelcomeRoutes()

	// 注册查看页面路由
	registerLookRoutes()

	// 注册默认 404 处理器
	http.NotFoundHandler()

	// 注册静态文件路由
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir(config.Config.Static))))
}
