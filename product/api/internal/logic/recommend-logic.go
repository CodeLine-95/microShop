package logic

import (
	"context"
	"encoding/json"
	"microShop/comm/errorx"
	"microShop/product/rpc/types/product"
	"net/http"
	"strconv"
	"strings"

	"microShop/product/api/internal/svc"
	"microShop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) RecommendLogic {
	return RecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendLogic) Recommend(req types.RecommendReq) (resp *types.RecommendResply, err error) {
	// 获取推荐表中的推荐数据
	RecommendRes, RecommendErr := l.svcCtx.Recommend.FindAll(l.ctx)
	if RecommendErr != nil {
		return nil, errorx.NewDefaultError(RecommendErr.Error())
	}
	// 处理推荐表的全部商品ID
	ProductIds := ""
	for _, v := range RecommendRes {
		ProductIds = ProductIds + strconv.FormatUint(v.ProductId, 10) + ","
	}
	ProductIds = strings.TrimRight(ProductIds, ",")
	// 传入推荐表中处理的商品ID，获取指定ID的商品
	cnt, cntErr := l.svcCtx.Rpc.Product(l.ctx, &product.ProductReq{ProductIds: ProductIds})
	if cntErr != nil {
		return nil, errorx.NewDefaultError(cntErr.Error())
	}

	products := []*types.ProductItem{}

	jsonErr := json.Unmarshal([]byte(cnt.Data), &products)
	if jsonErr != nil {
		return nil, errorx.NewDefaultError(jsonErr.Error())
	}

	return &types.RecommendResply{
		Code:    http.StatusOK,
		Message: cnt.Message,
		Data:    products,
	}, nil
}
