// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExitGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExitGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExitGroupLogic {
	return &ExitGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExitGroupLogic) ExitGroup(req *types.GroupExitRequestDTO) (resp bool, err error) {
	// todo: add your logic here and delete this line

	return
}
