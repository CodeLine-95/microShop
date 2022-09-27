package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"microShop/user/api/internal/svc"
	"microShop/user/api/internal/types"
	"microShop/user/rpc/account/user"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.CommonResply, err error) {
	cnt, cntErr := l.svcCtx.Rpc.Login(l.ctx, &user.LoginReq{
		UserName: req.UserName,
		PassWord: req.PassWord,
	})

	if cntErr != nil {
		return nil, cntErr
	}

	return &types.CommonResply{
		Code:    cnt.Code,
		Message: cnt.Message,
		Data:    cnt.Data,
	}, nil
}
