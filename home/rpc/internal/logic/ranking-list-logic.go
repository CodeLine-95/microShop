package logic

import (
	"context"

	"microShop/home/rpc/internal/svc"
	"microShop/home/rpc/types/home"

	"github.com/zeromicro/go-zero/core/logx"
)

type RankingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRankingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RankingListLogic {
	return &RankingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RankingListLogic) RankingList(in *home.RankingListReq) (*home.CommonResply, error) {
	// todo: add your logic here and delete this line

	return &home.CommonResply{}, nil
}
