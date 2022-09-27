package logic

import (
	"context"

	"microShop/home/rpc/internal/svc"
	"microShop/home/rpc/types/home"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendLogic {
	return &RecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecommendLogic) Recommend(in *home.RecommendReq) (*home.CommonResply, error) {
	// todo: add your logic here and delete this line

	return &home.CommonResply{}, nil
}
