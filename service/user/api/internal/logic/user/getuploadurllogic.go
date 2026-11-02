// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"context"

	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUploadUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUploadUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUploadUrlLogic {
	return &GetUploadUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUploadUrlLogic) GetUploadUrl(req *types.GetUploadUrlRequest) (resp *types.UploadUrlResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
