package controller

import (
	"go-web-practice/internal/templates"
	"net/http"
)

// 注册欢迎页面路由的函数
func registerWelcomeRoutes() {
	http.HandleFunc("/welcome", handleWelcome)
}

// 处理欢迎页面请求的函数
func handleWelcome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 使用欢迎页面模板渲染响应
		t := templates.Template.Lookup("welcome.tmpl")
		t.Execute(w, "欢迎访问 Go Web。")
	default:
		// 如果请求方法不是 GET，返回“方法不允许”状态码
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
