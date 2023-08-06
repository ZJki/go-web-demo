package controller

import (
	"go-web-practice/internal/templates"
	"net/http"
	"strings"
)

// 注册查看页面路由的函数
func registerLookRoutes() {
	http.HandleFunc("/look/", handleLook)
}

// 处理查看页面请求的函数
func handleLook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 切割 URL 来获取 URL 中的参数
		p := strings.Split(r.URL.Path, "/")
		switch p[2] {
		case "go":
			// 如果参数是 "go"，重定向到 Go 文档页面
			w.Header().Set("Location", "https://studygolang.com/pkgdoc")
			w.WriteHeader(302)
		case "down":
			// 渲染查看页面模板，显示 "Look down" 图片
			templates.Template.Lookup("look.tmpl").Execute(w, map[string]string{
				"title": "Look down",
				"image": "golang-down.png",
			})
		case "":
			// 渲染查看页面模板，显示 "Look right" 图片
			templates.Template.Lookup("look.tmpl").Execute(w, map[string]string{
				"title": "Look right",
				"image": "golang-right.png",
			})
		default:
			// 如果参数无法匹配，返回“未找到”状态码
			w.WriteHeader(http.StatusNotFound)
		}

	default:
		// 如果请求方法不是 GET，返回“方法不允许”状态码
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
