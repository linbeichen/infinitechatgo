package logic

import (
	"context"

	"infinitechatgo/service/user/rpc/internal/svc"
	"infinitechatgo/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateSingleMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateSingleMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateSingleMessageLogic {
	return &ValidateSingleMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateSingleMessageLogic) ValidateSingleMessage(in *user.ValidateSingleMessageReq) (*user.MessageValidateResp, error) {
	// todo: add your logic here and delete this line

	return &user.MessageValidateResp{}, nil
}
