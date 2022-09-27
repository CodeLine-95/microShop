package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"microShop/user/api/internal/config"
	"microShop/user/api/model"
	"microShop/user/rpc/account/user"
	"net/url"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	BizRedis  *redis.Redis
	Rpc       user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	DataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true&loc=%s",
		c.Mysql.User, c.Mysql.Pass, c.Mysql.Host, c.Mysql.Data, c.Mysql.Charset, url.QueryEscape("Asia/Shanghai"))
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(DataSource), c.CacheRedis),
		BizRedis:  redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		Rpc:       user.NewUser(zrpc.MustNewClient(c.Rpc)),
	}
}
