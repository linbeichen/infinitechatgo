// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Redis      redis.RedisConf
	DataSource string
	// go-zero model 层缓存配置
	Cache cache.CacheConf

	Etcd struct {
		Endpoints []string
		Prefix    string
	}

	Minio struct {
		Url       string
		AccessKey string
		SecretKey string
	}

	Kafka struct {
		// optional
		Brokers []string `json:",optional"` // kafka broker 地址列表
	}
}
