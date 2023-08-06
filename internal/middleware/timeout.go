package middleware

import (
	"context"
	"go-web-practice/internal/log"
	"net/http"
	"time"
)

// 超时中间件结构体
type TimeoutMiddleWare struct {
	Next http.Handler // 下一个处理器（中间件链中的下一个）
}

// ServeHTTP 方法实现了中间件的处理逻辑
func (t *TimeoutMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if t.Next == nil {
		t.Next = http.DefaultServeMux // 如果下一个处理器为空，使用默认处理器
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5) // 创建带有超时的上下文
	r = r.WithContext(ctx)                                 // 将带有超时上下文的请求重新赋值给 r

	ch := make(chan struct{})
	go func() {
		defer cancel() // 在处理完成或超时时取消上下文
		t.Next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()

	select {
	case <-ch: // 如果处理完成，不做额外操作
		return
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout) // 如果超时，返回请求超时状态码
		log.Info.Println("[请求处理超时！]")
	}
	ctx.Done()
}
