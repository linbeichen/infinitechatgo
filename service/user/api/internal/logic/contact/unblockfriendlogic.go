// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package contact

import (
	"context"

	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnblockFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnblockFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnblockFriendLogic {
	return &UnblockFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnblockFriendLogic) UnblockFriend(req *types.UserIdAndReceiveuserIdReq) (resp bool, err error) {
	// todo: add your logic here and delete this line

	return
}
