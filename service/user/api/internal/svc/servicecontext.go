// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"infinitechatgo/service/user/api/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(c.Redis)
	return &ServiceContext{
		Config: c,
		Redis:  rds,
	}
}
