// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"context"

	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KickGroupMembersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKickGroupMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KickGroupMembersLogic {
	return &KickGroupMembersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KickGroupMembersLogic) KickGroupMembers(req *types.KickGroupMembersRequest) (resp *types.KickGroupMembersResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
