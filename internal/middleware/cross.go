package middleware

import "net/http"

// 跨域中间件结构体
type CrossMiddleWare struct {
	Next http.Handler // 下一个处理器（中间件链中的下一个）
}

// ServeHTTP 方法实现了中间件的处理逻辑
func (c CrossMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if c.Next == nil {
		c.Next = http.DefaultServeMux // 如果下一个处理器为空，使用默认处理器
	}
	// 设置响应头，允许跨域请求
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	c.Next.ServeHTTP(w, r) // 调用下一个处理器继续处理请求
}
