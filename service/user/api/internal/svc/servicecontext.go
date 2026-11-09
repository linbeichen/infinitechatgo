// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"infinitechatgo/common/common"
	"infinitechatgo/service/user/api/internal/config"
	"infinitechatgo/service/user/api/internal/service"

	// model
	"infinitechatgo/service/user/model/apply_friend"
	"infinitechatgo/service/user/model/balance_log"
	"infinitechatgo/service/user/model/friend"
	"infinitechatgo/service/user/model/session"
	"infinitechatgo/service/user/model/user"
	"infinitechatgo/service/user/model/user_balance"
	"infinitechatgo/service/user/model/user_session"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis

	// WsServerLocator 用于定位到正确的 WebSocket 服务器
	WsServerLocator *service.WsServerLocator

	// 模型
	UserModel        user.UserModel
	UserBalanceModel user_balance.UserBalanceModel
	ApplyFriendModel apply_friend.ApplyFriendModel
	BalanceModel     balance_log.BalanceLogModel
	SessionModel     session.SessionModel
	UserSessionModel user_session.UserSessionModel
	FriendModel      friend.FriendModel

	// service
	UserService *service.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.DataSource) // 数据库连接
	// model
	userModel := user.NewUserModel(conn, c.Cache)
	userBalanceModel := user_balance.NewUserBalanceModel(conn, c.Cache)
	applyFriendModel := apply_friend.NewApplyFriendModel(conn, c.Cache)
	balanceModel := balance_log.NewBalanceLogModel(conn, c.Cache)
	sessionModel := session.NewSessionModel(conn, c.Cache)
	userSessionModel := user_session.NewUserSessionModel(conn, c.Cache)
	friendModel := friend.NewFriendModel(conn, c.Cache)

	rds := redis.MustNewRedis(c.Redis) // 创建 Redis 实例

	locator, err := service.NewWsServerLocator(c.Etcd.Endpoints, c.Etcd.Prefix)
	if err != nil {
		common.Errorf("etcd 连接失败，wsServerUri 将为空：%s", err.Error())
		locator = nil
	}

	// service
	userService := service.NewUserService(userModel, rds)

	return &ServiceContext{
		Config:           c,
		Redis:            rds,
		WsServerLocator:  locator,
		UserModel:        userModel,
		UserBalanceModel: userBalanceModel,
		ApplyFriendModel: applyFriendModel,
		BalanceModel:     balanceModel,
		SessionModel:     sessionModel,
		UserSessionModel: userSessionModel,
		FriendModel:      friendModel,
		UserService:      userService,
	}
}
