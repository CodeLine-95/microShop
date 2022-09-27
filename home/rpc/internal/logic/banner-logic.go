package logic

import (
	"context"
	"encoding/json"
	"microShop/home/rpc/internal/svc"
	"microShop/home/rpc/types/home"

	"github.com/zeromicro/go-zero/core/logx"
)

type BannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BannerLogic {
	return &BannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BannerLogic) Banner(in *home.BannerReq) (*home.CommonResply, error) {

	res, err := l.svcCtx.BannerModel.FindAll(l.ctx, "desc")
	if err != nil {
		return nil, err
	}

	jsonRes, _ := json.Marshal(res)

	return &home.CommonResply{
		Code:    200,
		Message: "成功",
		Data:    string(jsonRes),
	}, nil
}
