package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"microShop/comm/errorx"
	"microShop/product/rpc/types/product"
	"net/http"

	"microShop/product/api/internal/svc"
	"microShop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductListLogic {
	return ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req types.ProductsReq) (resp *types.ProductsResply, err error) {

	Search := "1 = 1"
	if len(req.ProductName) > 0 {
		Search = Search + fmt.Sprintf(" AND position('%s' in `title`)", req.ProductName)
	}

	if req.State > 0 {
		Search = Search + fmt.Sprintf(" AND state = %d", req.State)
	}

	if len(req.SearchTime) > 0 {
		Search = Search + fmt.Sprintf(" AND create_time between '%s' AND '%s'", req.SearchTime, req.SearchTime)
	}

	cnt, cntErr := l.svcCtx.Rpc.Products(l.ctx, &product.GetProductsReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Search:   Search,
	})
	if cntErr != nil {
		return nil, errorx.NewDefaultError(cntErr.Error())
	}

	products := []*types.ProductItem{}

	jsonErr := json.Unmarshal([]byte(cnt.Data), &products)
	if jsonErr != nil {
		return nil, errorx.NewDefaultError(jsonErr.Error())
	}

	return &types.ProductsResply{
		Code:       http.StatusOK,
		Message:    cnt.Message,
		Page:       req.Page,
		PageSize:   req.PageSize,
		TotalCount: int64(len(products)),
		Data:       products,
	}, nil
}
