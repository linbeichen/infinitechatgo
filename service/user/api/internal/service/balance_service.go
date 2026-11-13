package service

import (
	"context"
	"errors"
	"time"

	"infinitechatgo/common/common"
	"infinitechatgo/service/user/model/user_balance"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type BalanceService struct {
	UserBalanceModel user_balance.UserBalanceModel
	rds              *redis.Redis
}

func NewBalanceService(userBalanceModel user_balance.UserBalanceModel, rds *redis.Redis) *BalanceService {
	return &BalanceService{
		UserBalanceModel: userBalanceModel,
		rds:              rds,
	}
}

func (bs *BalanceService) InitializeBalance(ctx context.Context, userId int64, initialBalance int64) {
	common.ThrowIfWithMsg(userId <= 0, common.ParamsError, "用户ID不能为空")
	common.ThrowIfWithMsg(initialBalance <= 0, common.ParamsError, "初始余额不能为负数")

	// 检查是否已存在余额记录
	userBalance, err := bs.UserBalanceModel.FindOne(ctx, userId)
	// 出现非正常情况的错误
	common.ThrowIf(err != nil && !errors.Is(err, sqlc.ErrNotFound), common.SystemError, err)
	common.ThrowIfWithMsg(userBalance != nil, common.OperationError, "用户余额记录已存在")

	newUserBalance := &user_balance.UserBalance{
		UserId:      userId,
		Balance:     initialBalance,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	_, err = bs.UserBalanceModel.InsertTx(ctx, newUserBalance)
	common.ThrowIfWithMsg(err != nil, common.SystemError, "初始化用户余额失败", err)
}
