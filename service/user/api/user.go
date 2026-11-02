// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"

	"infinitechatgo/common/common"
	"infinitechatgo/common/utils"
	"infinitechatgo/service/user/api/internal/config"
	"infinitechatgo/service/user/api/internal/handler"
	"infinitechatgo/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 统一错误响应
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, any) {
		if errors.As(err, &common.SystemError) || errors.As(err, &common.SystemBusy) || errors.As(err, common.OperationError) {
			var e common.ErrorCode
			errors.As(err, &e)
			if e.Cause != nil {
				common.Errorf("系统异常，原因:%s", e.Cause.Error())
			} else {
				common.Errorf("系统异常，原因: %s", e.Message)
			}
			return http.StatusOK, utils.FailWithCode(500, e.Message)
		}

		// 判断是不是其他的 ErrorCode
		var e common.ErrorCode
		if errors.As(err, &e) {
			common.Infof("业务异常，原因:%s", e.Message)
			return http.StatusOK, utils.Fail(e)
		}

		// 其他错误
		common.Errorf("系统异常，原因：%s", err.Error())
		return http.StatusInternalServerError, utils.FailWithCode(500, common.SystemError.Message)
	})
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
