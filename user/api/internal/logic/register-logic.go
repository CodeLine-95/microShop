package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"microShop/user/api/internal/svc"
	"microShop/user/api/internal/types"
	"microShop/user/rpc/account/user"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (resp *types.CommonResply, err error) {
	cnt, cntErr := l.svcCtx.Rpc.Register(l.ctx, &user.RegisterReq{
		UserName:  req.UserName,
		PassWord:  req.PassWord,
		UserNick:  req.UserNick,
		UserFace:  req.UserFace,
		UserSex:   req.UserSex,
		UserEmail: req.UserEmail,
	})
	if cntErr != nil {
		return nil, cntErr
	}

	return &types.CommonResply{
		Code:    cnt.Code,
		Message: cnt.Message,
	}, nil
}
