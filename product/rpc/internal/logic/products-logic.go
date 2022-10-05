package logic

import (
	"context"

	"microShop/product/rpc/internal/svc"
	"microShop/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.GetProductsReq) (*product.CommonResply, error) {
	// todo: add your logic here and delete this line

	return &product.CommonResply{}, nil
}
