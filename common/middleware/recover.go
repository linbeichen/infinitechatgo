package middleware

import (
	"infinitechatgo/common/common"
	"infinitechatgo/common/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
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
