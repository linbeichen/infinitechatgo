package service

import (
	"context"
	"encoding/json"
	"infinitechatgo/common/common"

	"github.com/zeromicro/go-queue/kq"
)

type PusherManager struct {
	pushers map[string]*kq.Pusher
}

func NewPusherManager(brokers []string, topics []string) *PusherManager {
	pm := &PusherManager{
		pushers: make(map[string]*kq.Pusher),
	}

	for _, topic := range topics {
		pusher := kq.NewPusher(brokers, topic)
		pm.pushers[topic] = pusher
	}
	return pm
}

// 发送消息到指定topic （无key)
func (pm *PusherManager) Push(ctx context.Context, topic string, message interface{}) error {
	pusher, ok := pm.pushers[topic]
	if !ok {
		common.Errorf("Kafka Pusher 未找到，topic: %s", topic)
		return nil
	}

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		common.Errorf("消息序列化失败：%s", err.Error())
		return err
	}

	if err := pusher.Push(ctx, string(jsonBytes)); err != nil {
		common.Errorf("Kafka消息失败，topic: %s，error: %s", topic, err.Error())
	}

	common.Infof("Kafka消息发送成功，topic: %s", topic)
	return nil
}

/*
 */
func (pm *PusherManager) KPush(ctx context.Context, topic string, key string, message interface{}) error {
	pusher, ok := pm.pushers[topic]
	if !ok {
		common.Errorf("kafka Pusher 未找到，topic: %s", topic)
		return nil
	}

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		common.Errorf("消息序列化失败：%s", err.Error())
		return err
	}

	if err := pusher.KPush(ctx, key, string(jsonBytes)); err != nil {
		common.Errorf("kafka消息失败，topic: %s，key: %s，error: %s", topic, key, err.Error())
		return err
	}

	common.Infof("kafka消息发送成功，topic: %s，key: %s", topic, key)
	return nil
}

// 关闭所有 Pusher
func (pm *PusherManager) Close() {
	for topic, pusher := range pm.pushers {
		if err := pusher.Close(); err != nil {
			common.Errorf("关闭Kafka Pusher 失败，topic: %s，error: %s", topic, err.Error())
		}
	}
}
