package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"microShop/home/api/internal/config"
	"microShop/home/rpc/homeclient"
	"microShop/home/rpc/model"
	"net/url"
)

type ServiceContext struct {
	Config      config.Config
	BannerModel model.BannerModel
	BizRedis    *redis.Redis
	Rpc         homeclient.Home
}

func NewServiceContext(c config.Config) *ServiceContext {
	DataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true&loc=%s",
		c.Mysql.User, c.Mysql.Pass, c.Mysql.Host, c.Mysql.Data, c.Mysql.Charset, url.QueryEscape("Asia/Shanghai"))
	return &ServiceContext{
		Config:      c,
		BannerModel: model.NewBannerModel(sqlx.NewMysql(DataSource)),
		BizRedis:    redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		Rpc:         homeclient.NewHome(zrpc.MustNewClient(c.Rpc)),
	}
}
