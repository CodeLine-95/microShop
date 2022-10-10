package logic

import (
	"context"

	"microShop/product/api/internal/svc"
	"microShop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductDetailLogic {
	return ProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDetailLogic) ProductDetail(req types.DetailReq) (resp *types.DetailResply, err error) {
	// todo: add your logic here and delete this line

	return
}
