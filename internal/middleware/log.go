package middleware

import (
	"go-web-practice/internal/log"
	"net/http"
)

// 日志中间件结构体
type LogMiddleWare struct {
	Next http.Handler // 下一个处理器（中间件链中的下一个）
}

// ServeHTTP 方法实现了中间件的处理逻辑
func (m LogMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.Next == nil {
		m.Next = http.DefaultServeMux // 如果下一个处理器为空，使用默认处理器
	}
	// 记录请求信息到日志
	log.Info.Println("[请求]\t- 方法：" + r.Method + "\t\tURL：" + r.URL.Path)
	m.Next.ServeHTTP(w, r) // 调用下一个处理器继续处理请求
}
