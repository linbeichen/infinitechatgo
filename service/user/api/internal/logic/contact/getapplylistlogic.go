// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package contact

import (
	"context"

	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplyListLogic {
	return &GetApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApplyListLogic) GetApplyList(req *types.GetApplyListReq) error {
	// todo: add your logic here and delete this line

	return nil
}
