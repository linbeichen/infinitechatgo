package logic

import (
	"context"

	"infinitechatgo/service/user/rpc/internal/svc"
	"infinitechatgo/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetUserInfosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetUserInfosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetUserInfosLogic {
	return &BatchGetUserInfosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchGetUserInfosLogic) BatchGetUserInfos(in *user.BatchGetUserInfosReq) (*user.BatchGetUserInfosResp, error) {
	// todo: add your logic here and delete this line

	return &user.BatchGetUserInfosResp{}, nil
}
