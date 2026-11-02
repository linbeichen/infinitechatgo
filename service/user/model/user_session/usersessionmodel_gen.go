package user_session

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	userSessionFieldNames = builder.RawFieldNames(&UserSession{})
	userSessionRows       = strings.Join(userSessionFieldNames, ",")
	// userSessionRowsExpectAutoSet   = strings.Join(stringx.Remove(userSessionFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	// userSessionRowsWIthPlaceHolder = strings.Join(stringx.Remove(userSessionFieldNames, "`user_id`", "`session_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserSessionUserIdPrefix = "cache:userSession:userSessionId:"
)

type (
	userSessionModel interface {
		Insert(ctx context.Context, data *UserSession) (sql.Result, error)
		FindOne(ctx context.Context, userId int64, sessionId int64) (*UserSession, error)
		Update(ctx context.Context, data *UserSession) error
		Delete(ctx context.Context, userId int64, sessionId int64) error
	}

	defaultUserSessionModel struct {
		sqlc.CachedConn
		table string
	}

	UserSession struct {
		UserId      int64     `db:"user_id"`
		SessionId   int64     `db:"session_id"`
		Role        int64     `db:"role"`
		Status      int64     `db:"status"`
		CreatedTime time.Time `db:"created_time"`
		UpdatedTime time.Time `db:"updated_time"`
	}
)

func newUserSessionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserSessionModel {
	return &defaultUserSessionModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_session`",
	}
}

func (m *defaultUserSessionModel) Insert(ctx context.Context, data *UserSession) (sql.Result, error) {
	userSessionKey := fmt.Sprintf("%s%v:%v", cacheUserSessionUserIdPrefix, data.UserId, data.SessionId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (`user_id`, `session_id`, `role`, `status`) values (?, ?, ?, ?)", m.table)
		return conn.ExecCtx(ctx, query, data.UserId, data.SessionId, data.Role, data.Status)
	}, userSessionKey)
	return ret, err
}

func (m *defaultUserSessionModel) FindOne(ctx context.Context, userId int64, sessionId int64) (*UserSession, error) {
	userSessionKey := fmt.Sprintf("%s%v:%v", cacheUserSessionUserIdPrefix, userId, sessionId)
	var resp UserSession
	err := m.QueryRowCtx(ctx, &resp, userSessionKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `session_id` = ? limit 1", userSessionRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId, sessionId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserSessionModel) Update(ctx context.Context, data *UserSession) error {
	userSessionKey := fmt.Sprintf("%s%v:%v", cacheUserSessionUserIdPrefix, data.UserId, data.SessionId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `role` = ?, `status` = ? where `user_id` = ? and `session_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, data.Role, data.Status, data.UserId, data.SessionId)
	}, userSessionKey)
	return err
}

func (m *defaultUserSessionModel) Delete(ctx context.Context, userId int64, sessionId int64) error {
	userSessionKey := fmt.Sprintf("%s%v:%v", cacheUserSessionUserIdPrefix, userId, sessionId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ? and `session_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, userId, sessionId)
	}, userSessionKey)
	return err
}

func (m *defaultUserSessionModel) tableName() string {
	return m.table
}
