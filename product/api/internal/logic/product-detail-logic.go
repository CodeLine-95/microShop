package logic

import (
	"context"
	"encoding/json"
	"microShop/comm/errorx"
	"microShop/product/api/internal/svc"
	"microShop/product/api/internal/types"
	"microShop/product/model"
	product2 "microShop/product/rpc/types/product"

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

	cnt, cntErr := l.svcCtx.Rpc.Product(l.ctx, &product2.ProductReq{ProductId: req.PorductId})
	if cntErr != nil {
		return nil, errorx.NewDefaultError(cntErr.Error())
	}

	product := model.Product{}

	jsonErr := json.Unmarshal([]byte(cnt.Data), &product)
	if jsonErr != nil {
		return nil, errorx.NewDefaultError(jsonErr.Error())
	}
	FormatString := "2006-01-02 15:04:05"
	productData := types.ProductItem{
		ID:          int64(product.Id),
		Title:       product.Title,
		CoverPic:    product.CoverPic,
		Property:    product.Property,
		MtPrice:     product.MtPrice,
		DisPrice:    product.DisPrice,
		Stock:       int64(product.Stock),
		State:       int64(product.State),
		SalesVolume: product.SalesVolume,
		Images:      product.Images,
		Detail:      product.Detail,
		CreateTime:  product.CreateTime.Format(FormatString),
		UpdateTime:  product.UpdateTime.Format(FormatString),
	}

	return &types.DetailResply{
		Code:    cnt.Code,
		Message: cnt.Message,
		Data:    productData,
	}, nil
}
