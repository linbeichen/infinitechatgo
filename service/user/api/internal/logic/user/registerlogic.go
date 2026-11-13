// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"infinitechatgo/common/common"
	"infinitechatgo/common/txctx"
	"infinitechatgo/common/utils"
	"infinitechatgo/service/user/api/internal/svc"
	"infinitechatgo/service/user/api/internal/types"
	"infinitechatgo/service/user/api/internal/types/constants"
	userModel "infinitechatgo/service/user/model/user"

	"github.com/zeromicro/go-zero/core/logx"
)

const DefaultAvatar = "default_avatar.png"

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegisterRequest) (resp *types.LoginResponse, err error) {
	loginCodeKey := ""
	if constants.PhoneRegex.MatchString(req.Account) {
		code := "123456"
		if code != req.Code {
			return nil, common.LoginErrorCode
		}
	} else if constants.EmailRegex.MatchString(req.Account) {
		loginCodeKey = constants.LoginCodeEmailPrefix + req.Account
		code, _ := l.svcCtx.Redis.Get(loginCodeKey)
		if code == "" || code != req.Code {
			return nil, common.LoginErrorCode
		}
	}

	user, err := l.svcCtx.UserService.GetUser(req.Account)
	common.ThrowIfWithMsg(user != nil, common.UserAlreadyExists, common.UserAlreadyExists.Message, err)

	// 判断两次密码是否一致
	if req.Password != req.ConfirmPassword {
		return nil, common.LoginPasswordError
	}

	// 雪花算法生成userId
	userId := utils.NextId()

	newUser := &userModel.User{UserId: userId}
	newUser.Avatar = DefaultAvatar
	if constants.PhoneRegex.MatchString(req.Account) {
		newUser.Phone = sql.NullString{String: req.Account, Valid: true}
	} else if constants.EmailRegex.MatchString(req.Account) {
		newUser.Email = sql.NullString{String: req.Account, Valid: true}
	} else {
		return nil, common.PhoneEmailError
	}

	encryptedPassword := l.svcCtx.UserService.EncryptPassword(req.Password)
	newUser.Password = encryptedPassword

	newUser.Nickname = req.Nickname
	newUser.CreatedTime = time.Now()
	newUser.UpdatedTime = time.Now()

	// 事务
	err = txctx.WithTransaction(l.ctx, l.svcCtx.SqlConn, func(ctx context.Context) error {
		_, err2 := l.svcCtx.UserModel.InsertTx(ctx, newUser)
		common.ThrowIfWithMsg(err2 != nil, common.OperationError, "注册失败", err2)

		l.svcCtx.BalanceService.InitializeBalance(ctx, userId, 10000)

		robotId := int64(123)
		robotUser, _ := l.svcCtx.UserService.GetUserById(robotId)
		common.ThrowIfWithMsg(robotUser == nil, common.SystemError, "机器人用户不存在，robotId:"+strconv.FormatInt(robotId, 10))

		addFriendResp := l.svcCtx.FriendService.AddFriend(ctx, newUser, robotId)
		common.Infof("用户注册成功，已添加机器人好友，userId: %d, robotId: %d, sessionId: %s", userId, robotId, addFriendResp.SessionId)

		return nil
	})

	// 事务失败
	if err != nil {
		return nil, err
	}

	twoToken := l.svcCtx.UserService.GetTwoToken(userId)

	wsServerUri := l.svcCtx.WsServerLocator.GetWsServerUri(strconv.FormatInt(userId, 10))
	l.svcCtx.Redis.Hset("wsServerUri", strconv.FormatInt(userId, 10), wsServerUri)
	return &types.LoginResponse{
		UserId:       newUser.UserId,
		Account:      req.Account,
		Nickname:     newUser.Nickname,
		Avatar:       newUser.Avatar,
		Gender:       int32(newUser.Gender),
		Description:  newUser.Description.String,
		Token:        twoToken.AccessToken,
		RefreshToken: twoToken.RefreshToken,
		WsServerUri:  wsServerUri,
	}, nil
}
