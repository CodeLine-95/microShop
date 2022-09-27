package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"microShop/home/rpc/internal/config"
	"microShop/home/rpc/model"
	"net/url"
)

type ServiceContext struct {
	Config      config.Config
	BannerModel model.BannerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true&loc=%s",
		c.Mysql.User, c.Mysql.Pass, c.Mysql.Host, c.Mysql.Data, c.Mysql.Charset, url.QueryEscape("Asia/Shanghai"))
	return &ServiceContext{
		Config:      c,
		BannerModel: model.NewBannerModel(sqlx.NewMysql(DataSource)),
	}
}
