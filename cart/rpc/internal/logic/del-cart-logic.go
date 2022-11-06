package logic

import (
	"context"

	"microShop/cart/rpc/internal/svc"
	"microShop/cart/rpc/types/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCartLogic {
	return &DelCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCartLogic) DelCart(in *cart.DelCartReq) (*cart.CommonResply, error) {
	// todo: add your logic here and delete this line

	return &cart.CommonResply{}, nil
}
