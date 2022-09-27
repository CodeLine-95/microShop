package logic

import (
	"context"

	"microShop/home/api/internal/svc"
	"microShop/home/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) RecommendLogic {
	return RecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendLogic) Recommend(req types.RecommendReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
