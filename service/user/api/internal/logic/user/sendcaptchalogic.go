// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"context"
	"errors"

	"infinitechatgo/common/common"
	"infinitechatgo/common/utils"
	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"
	"infinitechatgo/service/user/api/internal/types/constants"
	utils2 "infinitechatgo/service/user/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type SendCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCaptchaLogic {
	return &SendCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCaptchaLogic) SendCaptcha(req *types.SendCaptchaReq) (resp string, err error) {
	// 发送验证码逻辑

	code := utils.RandamNumber(6)                            // 生成 6 位随机验证码
	common.Infof("发送验证码，账号: %s, 验证码: %s", req.Account, code) // 打印日志

	// 校验账号是否手机号或邮箱格式，拼接对应的 key
	var codeKey string
	if constants.PhoneRegex.MatchString(req.Account) {
		codeKey = constants.LoginCodePhonePrefix + req.Account
	} else if constants.EmailRegex.MatchString(req.Account) {
		codeKey = constants.LoginCodeEmailPrefix + req.Account
	}

	// 从上下文获取 redis, 并从redis获取验证码字符串
	codeStr, err := l.svcCtx.Redis.Get(codeKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		common.Errorf("从Redis获取验证码失败: %s", err.Error())
		return "", err
	}

	if constants.PhoneRegex.MatchString(req.Account) {
		if codeStr != "" {
			return constants.SmsSendFrequentlyMsg, nil
		}

	} else if constants.EmailRegex.MatchString(req.Account) {
		if codeStr != "" {
			return constants.SmsSendFrequentlyMsg, nil
		}

		// 发送邮箱验证码
		emailErr := utils2.SendEmailCode(req.Account, code)
		common.ThrowIfWithMsg(emailErr != nil, common.SystemError, "发送邮件验证码失败", emailErr)
	} else {
		return common.PhoneEmailError.Message, common.PhoneEmailError
	}
	redisErr := l.svcCtx.Redis.Setex(codeKey, code, int(constants.SmsExpireTime))
	common.ThrowIfWithMsg(redisErr != nil, common.SystemError, "发送邮件验证码", redisErr)
	return constants.SmsSendSuccessMsg, nil
}
