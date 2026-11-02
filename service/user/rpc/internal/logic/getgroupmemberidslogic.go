package logic

import (
	"context"

	"infinitechatgo/service/user/rpc/internal/svc"
	"infinitechatgo/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMemberIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMemberIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberIdsLogic {
	return &GetGroupMemberIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupMemberIdsLogic) GetGroupMemberIds(in *user.SessionIdReq) (*user.GroupMemberIdsResp, error) {
	// todo: add your logic here and delete this line

	return &user.GroupMemberIdsResp{}, nil
}
