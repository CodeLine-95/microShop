// Code generated by goctl. DO NOT EDIT!
// Source: account.proto

package server

import (
	"context"

	"microShop/user/rpc/account/internal/logic"
	"microShop/user/rpc/account/internal/svc"
	"microShop/user/rpc/account/types/account"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	account.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Register(ctx context.Context, in *account.RegisterReq) (*account.CommonResply, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Login(ctx context.Context, in *account.LoginReq) (*account.CommonResply, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) UserInfo(ctx context.Context, in *account.UserInfoReq) (*account.CommonResply, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}
