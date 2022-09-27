package logic

import (
	"context"

	"microShop/orders/api/internal/svc"
	"microShop/orders/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderDetailLogic {
	return OrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDetailLogic) OrderDetail(req types.OrderReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
