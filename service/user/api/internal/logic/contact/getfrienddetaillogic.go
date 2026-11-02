// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package contact

import (
	"context"

	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendDetailLogic {
	return &GetFriendDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendDetailLogic) GetFriendDetail(req *types.GetFriendDetailReq) (resp *types.FriendDetailVO, err error) {
	// todo: add your logic here and delete this line

	return
}
