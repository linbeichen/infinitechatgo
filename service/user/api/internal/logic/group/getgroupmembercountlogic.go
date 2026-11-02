// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMemberCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupMemberCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberCountLogic {
	return &GetGroupMemberCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupMemberCountLogic) GetGroupMemberCount(req *types.SessionIdPathReq) error {
	// todo: add your logic here and delete this line

	return nil
}
