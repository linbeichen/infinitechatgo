// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package contact

import (
	"context"

	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnreadApplyCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUnreadApplyCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUnreadApplyCountLogic {
	return &GetUnreadApplyCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUnreadApplyCountLogic) GetUnreadApplyCount(req *types.UserIdPathRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
