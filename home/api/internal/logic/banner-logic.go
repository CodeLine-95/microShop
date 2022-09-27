package logic

import (
	"context"
	"encoding/json"
	"microShop/home/rpc/types/home"
	"net/http"

	"microShop/home/api/internal/svc"
	"microShop/home/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) BannerLogic {
	return BannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BannerLogic) Banner(req types.BannerReq) (resp *types.BannerResply, err error) {

	cnt, cntErr := l.svcCtx.Rpc.Banner(l.ctx, &home.BannerReq{})
	if cntErr != nil {
		return nil, cntErr
	}

	banners := []*types.BannerItem{}

	jsonErr := json.Unmarshal([]byte(cnt.Data), &banners)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &types.BannerResply{
		Code:    http.StatusOK,
		Message: "获取成功",
		Data:    banners,
	}, nil
}
