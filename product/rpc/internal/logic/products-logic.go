package logic

import (
	"context"
	"encoding/json"
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

	res, err := l.svcCtx.Product.FindPaginations(l.ctx, in.Search, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}

	jsonRes, _ := json.Marshal(res)

	return &product.CommonResply{
		Code:    200,
		Message: "成功",
		Data:    string(jsonRes),
	}, nil
}
