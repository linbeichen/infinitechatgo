// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"context"
	"strconv"

	"infinitechatgo/common/common"
	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"
	"infinitechatgo/service/user/api/internal/types/constants"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginCodeLogic {
	return &LoginCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginCodeLogic) LoginCode(req *types.LoginCodeRequest) (resp *types.LoginResponse, err error) {
	loginCodeKey := ""
	if constants.PhoneRegex.MatchString(req.Account) {
		code := "123456" // 写死 123456
		if code != req.Code {
			return nil, common.LoginErrorCode
		}
		// 手机号场景使用固定验证码，不需要操作 Redis
	} else if constants.EmailRegex.MatchString(req.Account) {
		loginCodeKey = constants.LoginCodeEmailPrefix + req.Account
		code, _ := l.svcCtx.Redis.Get(loginCodeKey)
		if code == "" || code != req.Code {
			return nil, common.LoginErrorCode
		}
	} else {
		return nil, common.PhoneEmailError
	}

	// 验证码校验通过
	// 验证用户是否存在
	user, err := l.svcCtx.UserService.GetUser(req.Account)
	common.ThrowIfWithMsg(err != nil, common.UserNotExists, common.UserNotExists.Message)

	// 删除 redis 保存的验证码
	if loginCodeKey != "" {
		l.svcCtx.Redis.Del(loginCodeKey)
	}
	twoToken := l.svcCtx.UserService.GetTwoToken(user.UserId)

	wsServerUri := l.svcCtx.WsServerLocator.GetWsServerUri(strconv.FormatInt(user.UserId, 10))
	l.svcCtx.Redis.Hset("wsServerUri", strconv.FormatInt(user.UserId, 10), wsServerUri)

	return &types.LoginResponse{
		UserId:       user.UserId,
		Account:      req.Account,
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		Gender:       int32(user.Gender),
		Description:  user.Description.String,
		Token:        twoToken.AccessToken,
		RefreshToken: twoToken.RefreshToken,
		WsServerUri:  wsServerUri,
	}, nil
}
