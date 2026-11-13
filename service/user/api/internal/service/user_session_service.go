package service

import (
	"context"
	"infinitechatgo/common/common"
	"infinitechatgo/common/txctx"
	"infinitechatgo/service/user/api/internal/types/constants"
	"infinitechatgo/service/user/model/user_session"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserSessionService struct {
	UserSessionModel user_session.UserSessionModel
	rds              *redis.Redis
	SqlConn          sqlx.SqlConn
}

func NewUserSessionService(userSessionModel user_session.UserSessionModel, rds *redis.Redis, sqlConn sqlx.SqlConn) *UserSessionService {
	return &UserSessionService{
		UserSessionModel: userSessionModel,
		rds:              rds,
		SqlConn:          sqlConn,
	}
}

func (ss *UserSessionService) CreateUserSessions(ctx context.Context, userId int64, friendId int64, sessionId int64) {
	err := txctx.WithTransaction(ctx, ss.SqlConn, func(ctx context.Context) error {
		// 创建第一条用户会话关系
		_, err1 := ss.UserSessionModel.InsertTx(ctx, &user_session.UserSession{
			UserId:      userId,
			SessionId:   sessionId,
			Role:        constants.UserRoleNormal,
			Status:      constants.StatusSuccess,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		})
		// 创建第二条用户会话关系
		_, err2 := ss.UserSessionModel.InsertTx(ctx, &user_session.UserSession{
			UserId:      friendId,
			SessionId:   sessionId,
			Role:        constants.UserRoleNormal,
			Status:      constants.StatusSuccess,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		})
		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}
		return nil
	})
	common.ThrowIfWithMsg(err != nil, common.SystemError, "创建用户会话关系失败")
}
