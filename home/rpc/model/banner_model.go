package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BannerModel = (*customBannerModel)(nil)

type (
	// BannerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBannerModel.
	BannerModel interface {
		bannerModel
	}

	customBannerModel struct {
		*defaultBannerModel
	}
)

// NewBannerModel returns a model for the database table.
func NewBannerModel(conn sqlx.SqlConn) BannerModel {
	return &customBannerModel{
		defaultBannerModel: newBannerModel(conn),
	}
}
