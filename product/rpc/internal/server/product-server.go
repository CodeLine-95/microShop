// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package server

import (
	"context"

	"microShop/product/rpc/internal/logic"
	"microShop/product/rpc/internal/svc"
	"microShop/product/rpc/types/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) Products(ctx context.Context, in *product.GetProductsReq) (*product.CommonResply, error) {
	l := logic.NewProductsLogic(ctx, s.svcCtx)
	return l.Products(in)
}

func (s *ProductServer) Product(ctx context.Context, in *product.ProductReq) (*product.CommonResply, error) {
	l := logic.NewProductLogic(ctx, s.svcCtx)
	return l.Product(in)
}

func (s *ProductServer) Category(ctx context.Context, in *product.GetCateoryReq) (*product.CommonResply, error) {
	l := logic.NewCategoryLogic(ctx, s.svcCtx)
	return l.Category(in)
}
