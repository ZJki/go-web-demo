package controller

import (
	"go-web-practice/internal/templates"
	"net/http"
)

// 注册首页路由的函数
func registerIndexRoutes() {
	http.HandleFunc("/", handleIndex)
}

// 处理首页请求的函数
func handleIndex(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 使用模板渲染响应
		templates.Template.Lookup("index.tmpl").Execute(w, map[string]string{
			"welcome": "/welcome",
			"look":    "/look",
			"statics": "/web",
		})
	default:
		// 如果请求方法不是 GET，返回“方法不允许”状态码
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
