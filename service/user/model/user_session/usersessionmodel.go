package user_session

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserSessionModel = (*customUserSessionModel)(nil)

type (
	UserSessionModel interface {
		userSessionModel
	}

	customUserSessionModel struct {
		*defaultUserSessionModel
	}
)

func NewUserSessionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserSessionModel {
	return &customUserSessionModel{
		defaultUserSessionModel: newUserSessionModel(conn, c, opts...),
	}
}
