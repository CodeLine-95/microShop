package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"microShop/cart/model"
	"microShop/cart/rpc/internal/svc"
	"microShop/cart/rpc/types/cart"
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

	var commonResp cart.CommonResply

	commonResp.Code = 200
	commonResp.Message = "成功"

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		commonResp.Code = 400
		commonResp.Message = "失败"
	}
	return &commonResp, nil
}
