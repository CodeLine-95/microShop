package logic

import (
	"context"

	"microShop/orders/api/internal/svc"
	"microShop/orders/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderLogisticsLogic {
	return OrderLogisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogisticsLogic) OrderLogistics(req types.OrderReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
