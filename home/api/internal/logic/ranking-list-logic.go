package logic

import (
	"context"

	"microShop/home/api/internal/svc"
	"microShop/home/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RankingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRankingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) RankingListLogic {
	return RankingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RankingListLogic) RankingList(req types.RankingListReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
