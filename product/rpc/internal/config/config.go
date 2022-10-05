package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		Host    string
		User    string
		Pass    string
		Data    string
		Charset string
	}
	Cache cache.CacheConf
}
