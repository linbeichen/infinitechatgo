package service

import (
	"context"
	"fmt"
	"infinitechatgo/common/common"
	"infinitechatgo/common/model/constants"
	"infinitechatgo/common/utils"
	constants2 "infinitechatgo/service/user/api/internal/types/constants"
	"infinitechatgo/service/user/api/internal/types/dto"
	"strconv"
	"time"
)

/*
 * 通知推送服务
 *
 * 功能说明：
 * - 使用Kafka异步发送系统通知消息
 * - 替代原有的HTTP同步调用方式
 * - 符合IM 项目通知消息设计方式
 */
type NotificationService struct {
	PusherManager *PusherManager
}

func NewNotificationService(pusherManager *PusherManager) *NotificationService {
	return &NotificationService{
		PusherManager: pusherManager,
	}
}

/*
 * 推送新会话通知
 */
func (ns *NotificationService) PushNewSession(senderId, userId, sessionId int64, sessionType int, notification dto.NewSessionNotificationDTO) {
	defer func() {
		if r := recover(); r != nil {
			common.Errorf("发送新会话失败，用户ID，会话ID：%d，错误: %v", userId, sessionId, r)
		}
	}()

	message := dto.SystemNotificationMessage{
		MessageId:   generateMessageId(),
		SessionId:   &sessionId,
		SenderId:    &senderId,
		ReceiverId:  userId,
		Type:        constants2.TypeSystemNewSession, // 102
		SessionType: &sessionType,
		Timestamp:   time.Now().UnixMilli(),
		Body: map[string]interface{}{
			"sessionName": notification.SessionName,
			"avatar":      notification.Avatar,
		},
	}
	ns.sendNotification(message, "新会话通知")
}

/*
 * 推送好友申请通知
 */
func (ns *NotificationService) PushNewApply(userId int64, notification dto.FriendApplicationNotificationDTO) {
	defer func() {
		if r := recover(); r != nil {
			common.Errorf("发送好友申请通知失败，用户ID: %d, 错误: %v", userId, r)
		}
	}()

	senderId := notification.ApplyUserId
	message := dto.SystemNotificationMessage{
		MessageId:   generateMessageId(),
		SessionId:   nil,
		SenderId:    &senderId,
		ReceiverId:  userId,
		Type:        constants2.TypeSystemNewApply, // 101
		SessionType: nil,
		Timestamp:   time.Now().UnixMilli(),
		Body: map[string]interface{}{
			"nickname": notification.ApplyUserName,
			"avatar":   notification.ApplyFriendAvatar,
			"msg":      notification.Message,
		},
	}

	ns.sendNotification(message, "好友申请通知")
}

/*
 * 推送新群聊会话通知
 */
func (ns *NotificationService) PushGroupNewSession(userId, sessionId int64, notification dto.NewGroupSessionNotificationDTO) {
	defer func() {
		if r := recover(); r != nil {
			common.Errorf("发送新群聊会话通知失败，用户ID: %d, 会话ID: %d, 错误: %v", userId, sessionId, r)
		}
	}()

	groupType := constants2.SessionTypeGroup
	message := dto.SystemNotificationMessage{
		MessageId:   generateMessageId(),
		SessionId:   &sessionId,
		SenderId:    nil,
		ReceiverId:  userId,
		Type:        constants2.TypeSystemNewGroupSession, // 103
		SessionType: &groupType,
		Timestamp:   time.Now().UnixMilli(),
		Body: map[string]interface{}{
			"sessionName":  notification.SessionName,
			"avatar":       notification.Avatar,
			"creatorId":    notification.CreatorId,
			"membersCount": notification.MembersCount,
		},
	}

	ns.sendNotification(message, "群组邀请通知")
}

/*
 * 推送群聊踢出通知
 */
func (ns *NotificationService) PushGroupKickNotification(receiverId, sessionId int64, notification dto.GroupKickNotificationDTO) {
	defer func() {
		if r := recover(); r != nil {
			common.Errorf("发送群聊踢出通知失败，用户ID: %d, 会话ID: %d, 错误: %v", receiverId, sessionId, r)
		}
	}()

	groupType := constants2.SessionTypeGroup
	message := dto.SystemNotificationMessage{
		MessageId:   generateMessageId(),
		SessionId:   &sessionId,
		SenderId:    nil,
		ReceiverId:  receiverId,
		Type:        constants2.TypeSystemGroupKick, // 104
		SessionType: &groupType,
		Timestamp:   time.Now().UnixMilli(),
		Body: map[string]interface{}{
			"memberIds":  notification.MemberIds,
			"operatorId": notification.OperatorId,
		},
	}

	ns.sendNotification(message, "群聊踢出通知")
}

/*
 * 发送消息到Kafka
 */
func (ns *NotificationService) sendNotification(message dto.SystemNotificationMessage, notificationName string) {
	if ns.PusherManager == nil {
		common.Errorf("PusherManager 未初始化，跳过 %s", notificationName)
		return
	}

	key := strconv.FormatInt(message.ReceiverId, 10)

	err := ns.PusherManager.KPush(
		context.Background(),
		constants.KafkaSystemNotificationTopic,
		key,
		message,
	)

	if err != nil {
		common.Errorf("%s发送失败，messageId: %s，用户ID: %d, type: %d, 错误: %s", notificationName, message.MessageId, message.ReceiverId, message.Type, err.Error())
	} else {
		common.Infof("%s发送成功, messageId: %s, 用户ID: %d, type: %d", notificationName, message.MessageId, message.ReceiverId, message.Type)
	}
}

func generateMessageId() string {
	return fmt.Sprintf("msg_%d_%d", time.Now().UnixMilli(), utils.NextId())
}
