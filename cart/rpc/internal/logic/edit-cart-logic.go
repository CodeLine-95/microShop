package logic

import (
	"context"

	"microShop/cart/rpc/internal/svc"
	"microShop/cart/rpc/types/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditCartLogic {
	return &EditCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditCartLogic) EditCart(in *cart.EditCartReq) (*cart.CommonResply, error) {
	// todo: add your logic here and delete this line

	return &cart.CommonResply{}, nil
}
