package logic

import (
	"context"
	"encoding/json"
	"microShop/comm/errorx"
	"microShop/product/api/internal/svc"
	"microShop/product/api/internal/types"
	"microShop/product/model"
	"microShop/product/rpc/types/product"

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

	cnt, cntErr := l.svcCtx.Rpc.Product(l.ctx, &product.ProductReq{ProductId: req.PorductId})
	if cntErr != nil {
		return nil, errorx.NewDefaultError(cntErr.Error())
	}

	productItem := model.Product{}

	jsonErr := json.Unmarshal([]byte(cnt.Data), &productItem)
	if jsonErr != nil {
		return nil, errorx.NewDefaultError(jsonErr.Error())
	}
	FormatString := "2006-01-02 15:04:05"
	productData := types.ProductItem{
		ID:          int64(productItem.Id),
		Title:       productItem.Title,
		CoverPic:    productItem.CoverPic,
		Property:    productItem.Property,
		MtPrice:     productItem.MtPrice,
		DisPrice:    productItem.DisPrice,
		Stock:       int64(productItem.Stock),
		State:       int64(productItem.State),
		SalesVolume: productItem.SalesVolume,
		Images:      productItem.Images,
		Detail:      productItem.Detail,
		CreateTime:  productItem.CreateTime.Format(FormatString),
		UpdateTime:  productItem.UpdateTime.Format(FormatString),
	}

	return &types.DetailResply{
		Code:    cnt.Code,
		Message: cnt.Message,
		Data:    productData,
	}, nil
}
