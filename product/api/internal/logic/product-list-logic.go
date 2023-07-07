package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"microShop/comm/errorx"
	"microShop/product/model"
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

	products := []*model.Product{}

	jsonErr := json.Unmarshal([]byte(cnt.Data), &products)
	if jsonErr != nil {
		return nil, errorx.NewDefaultError(jsonErr.Error())
	}

	productsData, productsDataErr := mr.MapReduce(func(source chan<- any) {
		// 这里是 generate | 将列表的下标值记录到 chan
		// 传入到 mapper
		for key := range products {
			source <- key
		}
	}, func(item any, writer mr.Writer[any], cancel func(error)) {
		// 这里是 mapper | 处理存入的 chan
		key := item.(int)
		// 写入到 reducer
		writer.Write(key)
	}, func(pipe <-chan any, writer mr.Writer[any], cancel func(error)) {
		// 处理 reducer | 对 mapper 进行数据聚合
		productsData := []types.ProductItem{}
		productsDataMap := []any{}
		for p := range pipe {
			itemOne := products[p.(int)]
			FormatString := "2006-01-02 15:04:05"
			productItem := types.ProductItem{
				ID:          int64(itemOne.Id),
				Title:       itemOne.Title,
				CoverPic:    itemOne.CoverPic,
				Property:    itemOne.Property,
				MtPrice:     itemOne.MtPrice,
				DisPrice:    itemOne.DisPrice,
				Stock:       int64(itemOne.Stock),
				State:       int64(itemOne.State),
				SalesVolume: itemOne.SalesVolume,
				Images:      itemOne.Images,
				Detail:      itemOne.Detail,
				CreateTime:  itemOne.CreateTime.Format(FormatString),
				UpdateTime:  itemOne.UpdateTime.Format(FormatString),
			}
			// 插入 []数组
			productsDataMap = append(productsDataMap, productItem)
			// interface 转  []types.ProductItem{}
			productsDataJson, _ := json.Marshal(productsDataMap)
			json.Unmarshal(productsDataJson, &productsData)
		}
		// 输出到结果
		writer.Write(productsData)
	})

	if productsDataErr != nil {
		return nil, errorx.NewDefaultError(productsDataErr.Error())
	}

	productsDataList := []*types.ProductItem{}
	// 对 MapReduce 结果进行转换 | 因为 MapReduce 返回的结果是 interface{} 跟咱们返回的Data类型不一致
	// interface 转  []*types.ProductItem{}
	productsDataJson, _ := json.Marshal(productsData)
	json.Unmarshal(productsDataJson, &productsDataList)

	return &types.ProductsResply{
		Code:       http.StatusOK,
		Message:    cnt.Message,
		Page:       req.Page,
		PageSize:   req.PageSize,
		TotalCount: int64(len(productsDataList)),
		Data:       productsDataList,
	}, nil
}
