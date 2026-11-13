package friend

import (
	"context"
	"database/sql"
	"fmt"
	"infinitechatgo/common/txctx"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FriendModel = (*customFriendModel)(nil)

type (
	// FriendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendModel.
	FriendModel interface {
		friendModel
		InsertTx(ctx context.Context, data *Friend) (sql.Result, error)
		UpdateTx(ctx context.Context, data *Friend) error
		DeleteTx(ctx context.Context, id int64) error
	}

	customFriendModel struct {
		*defaultFriendModel
	}
)

// NewFriendModel returns a model for the database table.
func NewFriendModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FriendModel {
	return &customFriendModel{
		defaultFriendModel: newFriendModel(conn, c, opts...),
	}
}

// 如果 ctx 中有事务 session 则走事务连接，否则走原来的 CachedConn
func (m *customFriendModel) InsertTx(ctx context.Context, data *Friend) (sql.Result, error) {
	if session := txctx.GetSession(ctx); session != nil {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, friendRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.FriendId, data.Status, data.CreatedTime, data.UpdatedTime)
	}
	return m.Insert(ctx, data)
}

// 如果 ctx 中有事务 session 则走事务连接，否则走原来的 CachedConn
func (m *customFriendModel) UpdateTx(ctx context.Context, data *Friend) error {
	if session := txctx.GetSession(ctx); session != nil {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, friendRowsWithPlaceHolder)
		_, err := session.ExecCtx(ctx, query, data.UserId, data.FriendId, data.Status, data.CreatedTime, data.UpdatedTime, data.Id)
		return err
	}
	return m.Update(ctx, data)
}

// 如果 ctx 中有事务 session 则走事务连接，否则走原来的 CachedConn
func (m *customFriendModel) DeleteTx(ctx context.Context, id int64) error {
	if session := txctx.GetSession(ctx); session != nil {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		_, err := session.ExecCtx(ctx, query, id)
		return err
	}
	return m.Delete(ctx, id)
}
