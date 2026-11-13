package apply_friend

import (
	"context"
	"database/sql"
	"fmt"
	"infinitechatgo/common/txctx"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 编译期检查，确保 customApplyFriendModel 实现了 ApplyFriendModel 接口
var _ ApplyFriendModel = (*customApplyFriendModel)(nil)

type (
	// 提供给业务层使用的公开接口，封装 applyFriendModel 接口，可以直接调用方法
	// ApplyFriendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplyFriendModel.
	ApplyFriendModel interface {
		applyFriendModel
		InsertTx(ctx context.Context, data *ApplyFriend) (sql.Result, error)
		UpdateTx(ctx context.Context, data *ApplyFriend) error
		DeleteTx(ctx context.Context, applyFriendId int64) error
	}

	// 生成层的自定义扩展实现
	// 内嵌了 defaultApplyFriendModel 结构体，继承了其方法
	// 同时 applyfriendmodel.go 内的自定义内容会在 goctl 自动生成时保留。
	customApplyFriendModel struct {
		*defaultApplyFriendModel
	}
)

// NewApplyFriendModel returns a model for the database table.
// 业务层/model 的构造函数
// conn sqlx.SqlConn Mysql 数据库连接
// c cache.CacheConf 缓存配置
// opts ...cache.Option 可选缓存配置项
// 返回 customApplyFriendModel 实例
// customApplyFriendModel 内部调用 defaultApplyFriendModel 的构造函数，初始化数据库连接和缓存配置
func NewApplyFriendModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ApplyFriendModel {
	return &customApplyFriendModel{
		defaultApplyFriendModel: newApplyFriendModel(conn, c, opts...),
	}
}

// InsertTx 如果 ctx 中有事务 session 则走事务连接，否则走原来的 CachedConn
func (m *customApplyFriendModel) InsertTx(ctx context.Context, data *ApplyFriend) (sql.Result, error) {
	if session := txctx.GetSession(ctx); session != nil {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, applyFriendRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.ApplyFriendId, data.SenderId, data.ReceiverId, data.Message, data.Status, data.CreatedTime, data.UpdatedTime)
	}
	return m.Insert(ctx, data)
}

// UpdateTx 如果 ctx 中有事务 session 则走事务连接，否则走原来的 CachedConn
func (m *customApplyFriendModel) UpdateTx(ctx context.Context, data *ApplyFriend) error {
	if session := txctx.GetSession(ctx); session != nil {
		query := fmt.Sprintf("update %s set %s where `apply_friend_id` = ?", m.table, applyFriendRowsWithPlaceHolder)
		_, err := session.ExecCtx(ctx, query, data.SenderId, data.ReceiverId, data.Message, data.Status, data.CreatedTime, data.UpdatedTime, data.ApplyFriendId)
		return err
	}
	return m.Update(ctx, data)
}

// DeleteTx 如果 ctx 中有事务 session 则走事务连接，否则走原来的 CachedConn
func (m *customApplyFriendModel) DeleteTx(ctx context.Context, applyFriendId int64) error {
	if session := txctx.GetSession(ctx); session != nil {
		query := fmt.Sprintf("delete from %s where `apply_friend_id` = ?", m.table)
		_, err := session.ExecCtx(ctx, query, applyFriendId)
		return err
	}
	return m.Delete(ctx, applyFriendId)
}
