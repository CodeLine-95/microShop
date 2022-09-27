package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	SecretKey string
	Mysql     struct {
		Host    string
		User    string
		Pass    string
		Data    string
		Charset string
	}
	CacheRedis cache.CacheConf
	BizRedis   redis.RedisConf

	Rpc zrpc.RpcClientConf
}
