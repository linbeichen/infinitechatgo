package middleware

import (
	"infinitechatgo/common/common"
	"infinitechatgo/common/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 自定义中间件函数, middleware(next http.HandlerFunc) http.HandlerFunc
// 全局 panic 捕捉中间件
func RecoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// 返回匿名函数
	return func(w http.ResponseWriter, r *http.Request) {
		// defer 注册保护逻辑
		defer func() {
			// 用 recover 捕捉 panic
			if err := recover(); err != nil {
				// 通过类型推断判断是否是正常错误
				if bizErr, ok := err.(common.ErrorCode); ok {
					// 系统级错误走 ErrorCtx, 让 user.go 的 SetErrorHandlerCtx 统一管理
					if bizErr.Code >= 50000 {
						httpx.ErrorCtx(r.Context(), w, bizErr)
						return
					}
					// 业务错误正常返回
					httpx.OkJsonCtx(r.Context(), w, utils.Fail(bizErr))
					return
				}
				common.Errorf("panic recovered, 原因:%v", err)
				httpx.ErrorCtx(r.Context(), w, common.SystemError)
			}
		}()
		next(w, r)
	}
}
