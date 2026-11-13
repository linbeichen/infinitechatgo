package service

import (
	"context"
	"infinitechatgo/common/common"
	"infinitechatgo/service/user/model/session"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type SessionService struct {
	SessionModel session.SessionModel
	rds          *redis.Redis
}

func NewSessionService(sessionModel session.SessionModel, rds *redis.Redis) *SessionService {
	return &SessionService{
		SessionModel: sessionModel,
		rds:          rds,
	}
}

/*
创建会话
*/
func (ss *SessionService) CreateSession(ctx context.Context, newSession *session.Session) int64 {
	res, err := ss.SessionModel.InsertTx(ctx, newSession)
	common.ThrowIfWithMsg(err != nil || res == nil, common.SystemError, "创建会话失败")
	id, _ := res.LastInsertId()
	return id
}
