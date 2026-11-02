// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"context"

	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBalanceDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBalanceDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBalanceDetailLogic {
	return &GetBalanceDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBalanceDetailLogic) GetBalanceDetail(req *types.GetBalanceDetailRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
