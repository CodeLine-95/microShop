package logic

import (
	"context"
	"encoding/json"
	"microShop/cart/model"
	"microShop/cart/rpc/internal/svc"
	"microShop/cart/rpc/types/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCartLogic) AddCart(in *cart.AddCartReq) (*cart.CommonResply, error) {
	res, err := l.svcCtx.Cart.Insert(l.ctx, &model.Cart{
		UserId:    uint64(in.UserId),
		ProductId: uint64(in.ProductId),
		BuyCount:  in.BuyCount,
		Checked:   in.Checked,
	})
	if err != nil {
		return nil, err
	}
	jsonRes, _ := json.Marshal(res)
	return &cart.CommonResply{
		Code:    200,
		Message: "成功",
		Data:    string(jsonRes),
	}, nil
}
