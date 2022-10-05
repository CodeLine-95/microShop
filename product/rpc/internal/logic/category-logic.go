package logic

import (
	"context"
	"encoding/json"
	"microShop/product/rpc/internal/svc"
	"microShop/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryLogic {
	return &CategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryLogic) Category(in *product.GetCateoryReq) (*product.CommonResply, error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.Category.FindPids(l.ctx, in.Pid)
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
