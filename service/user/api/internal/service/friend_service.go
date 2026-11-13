package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"infinitechatgo/common/common"
	"infinitechatgo/common/txctx"
	"infinitechatgo/common/utils"

	"infinitechatgo/service/user/api/internal/types/constants"
	"infinitechatgo/service/user/api/internal/types/dto"
	"infinitechatgo/service/user/model/apply_friend"
	"infinitechatgo/service/user/model/friend"
	"infinitechatgo/service/user/model/session"
	"infinitechatgo/service/user/model/user"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type FriendService struct {
	IUserService        *UserService
	ISessionService     *SessionService
	IUserSessionService *UserSessionService
	FriendModel         friend.FriendModel
	ApplyFriendModel    apply_friend.ApplyFriendModel
	rds                 *redis.Redis
	SqlConn             sqlx.SqlConn
	NotificationService *NotificationService
}

func NewFriendService(iUserService *UserService, iSessionService *SessionService, iUserSessionService *UserSessionService, friendModel friend.FriendModel, applyFriendModel apply_friend.ApplyFriendModel, rds *redis.Redis, sqlConn sqlx.SqlConn, notificationService *NotificationService) *FriendService {
	return &FriendService{
		IUserService:        iUserService,
		ISessionService:     iSessionService,
		IUserSessionService: iUserSessionService,
		FriendModel:         friendModel,
		ApplyFriendModel:    applyFriendModel,
		rds:                 rds,
		SqlConn:             sqlConn,
		NotificationService: notificationService,
	}
}

/**
 * 添加好友关系
 *
 * 处理流程：
 * 1. 验证用户存在性
 * 2. 检查是否已是好友关系
 * 3. 创建双向好友关系
 * 4. 创建会话和用户会话关系
 * 5. 发送Kafka通知
 *
 *  recipient 接收好友请求的用户 【to】
 *  friendId  申请添加的好友ID 【from】
 *  ModifyFriendApplicationResponse 响应对象
 */
func (fs *FriendService) AddFriend(ctx context.Context, recipient *user.User, friendId int64) dto.ModifyFriendApplicationResponse {
	common.ThrowIfWithMsg(friendId <= 0, common.ParamsError, "好友申请的发起者用户ID不存在")
	common.ThrowIfWithMsg(recipient.UserId <= 0, common.NotFoundError, "好友申请的接收者用户ID不存在")

	applicant, err1 := fs.IUserService.GetUserById(friendId)
	common.ThrowIfWithMsg(err1 != nil && !errors.Is(err1, sqlx.ErrNotFound), common.SystemError, "获取好友申请者信息失败")
	common.ThrowIfWithMsg(applicant == nil, common.NotFoundError, "好友申请者不存在")

	//检查是否已是好友关系
	friendRelation, err2 := fs.FriendModel.FindOneByUserIdFriendId(ctx, recipient.UserId, friendId)
	common.ThrowIfWithMsg(err2 != nil && !errors.Is(err2, friend.ErrNotFound), common.SystemError, "获取好友关系失败")
	common.ThrowIfWithMsg(friendRelation != nil, common.OperationError, "已经是好友关系")

	// ========== 事务开始 ==========
	sessionId := utils.NextId()
	err := txctx.WithTransaction(ctx, fs.SqlConn, func(ctx context.Context) error {
		//创建双向好友关系
		fs.createFriendRelations(ctx, recipient.UserId, friendId)
		//创建会话
		fs.ISessionService.CreateSession(ctx, &session.Session{
			SessionId:   sessionId,
			Name:        "",
			Type:        constants.SignalType,
			Status:      constants.StatusSuccess,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		})

		//创建用户会话关系
		fs.IUserSessionService.CreateUserSessions(ctx, recipient.UserId, friendId, sessionId)
		return nil
	})

	// ========== 事务结束 ==========
	common.ThrowIfWithMsg(err != nil, common.SystemError, "添加好友失败")

	// 发送Kafka通知给申请者

	fs.sendNewSessionNotification(friendId, recipient, sessionId)

	//构建响应对象
	return dto.ModifyFriendApplicationResponse{
		UserId:      strconv.FormatInt(applicant.UserId, 10),
		SessionId:   strconv.FormatInt(sessionId, 10),
		SessionType: constants.SignalType,
		SessionName: applicant.Nickname,
		Avatar:      applicant.Avatar,
	}
}

/*
 * 发送新会话通知给申请者
 *
 * 好友申请被通过后，通知申请者有新的会话创建
 * 通知内容包含接收者（recipient）的昵称和头像，用于申请者的聊天列表展示
 */
func (fs *FriendService) sendNewSessionNotification(friendId int64, recipient *user.User, sessionId int64) {
	if fs.NotificationService == nil {
		common.Errorf("NotificationService 未初始化，跳过发送新会话通知")
		return
	}

	notification := dto.NewSessionNotificationDTO{
		SessionName: recipient.Nickname,
		Avatar:      recipient.Avatar,
	}

	// goroutine 异步发送
	go func() {
		defer func() {
			if r := recover(); r != nil {
				common.Errorf("发送新会话通知异常，申请者ID: %d, 会话ID: %d, 错误: %v", friendId, sessionId, r)
			}
		}()

		fs.NotificationService.PushNewSession(
			recipient.UserId,            // senderId
			friendId,                    // userId（接收通知的人）
			sessionId,                   // 新会话ID
			constants.SessionTypeSignal, // 单聊
			notification,
		)

		common.Infof("发送新会话通知成功，申请者ID: %d, 会话ID: %d", friendId, sessionId)
	}()
}

/*
创建双向好友关系
userId 当前用户ID
friendId 好友用户ID
*/
func (fs *FriendService) createFriendRelations(ctx context.Context, userId int64, friendId int64) {
	friend1 := &friend.Friend{
		UserId:      userId,
		FriendId:    friendId,
		Status:      constants.Normal.Code,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	friend2 := &friend.Friend{
		UserId:      friendId,
		FriendId:    userId,
		Status:      constants.Normal.Code,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	err := txctx.WithTransaction(ctx, fs.SqlConn, func(ctx context.Context) error {
		_, err1 := fs.FriendModel.InsertTx(ctx, friend1)
		_, err2 := fs.FriendModel.InsertTx(ctx, friend2)
		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}
		return nil
	})

	common.ThrowIfWithMsg(err != nil, common.SystemError, "添加好友关系失败")
}
