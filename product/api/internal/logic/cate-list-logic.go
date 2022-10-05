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

type CateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CateListLogic {
	return CateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CateListLogic) CateList(req types.CateListReq) (resp *types.CategoryResply, err error) {
	// todo: add your logic here and delete this line

	fmt.Println(req.Pid)

	cnt, cntErr := l.svcCtx.Rpc.Category(l.ctx, &product.GetCateoryReq{Pid: req.Pid})
	if cntErr != nil {
		return nil, errorx.NewDefaultError(cntErr.Error())
	}

	category := []*types.CategoryItem{}

	jsonErr := json.Unmarshal([]byte(cnt.Data), &category)
	if jsonErr != nil {
		return nil, errorx.NewDefaultError(jsonErr.Error())
	}

	return &types.CategoryResply{
		Code:    http.StatusOK,
		Message: cnt.Message,
		Data:    category,
	}, nil
}
