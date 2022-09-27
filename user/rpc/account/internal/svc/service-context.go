package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"microShop/user/rpc/account/internal/config"
	"microShop/user/rpc/account/model"
	"net/url"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true&loc=%s",
		c.Mysql.User, c.Mysql.Pass, c.Mysql.Host, c.Mysql.Data, c.Mysql.Charset, url.QueryEscape("Asia/Shanghai"))
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(DataSource), c.Cache),
	}
}
