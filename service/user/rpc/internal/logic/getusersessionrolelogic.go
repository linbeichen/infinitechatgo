package logic

import (
	"context"

	"infinitechatgo/service/user/rpc/internal/svc"
	"infinitechatgo/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSessionRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSessionRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSessionRoleLogic {
	return &GetUserSessionRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserSessionRoleLogic) GetUserSessionRole(in *user.UserSessionRoleReq) (*user.UserSessionRoleResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserSessionRoleResp{}, nil
}
