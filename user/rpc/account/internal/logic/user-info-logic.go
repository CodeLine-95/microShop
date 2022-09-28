package logic

import (
	"context"
	"encoding/json"
	"microShop/comm/errorx"
	"net/http"

	"microShop/user/rpc/account/internal/svc"
	"microShop/user/rpc/account/types/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *account.UserInfoReq) (*account.CommonResply, error) {
	// 获取用户信息
	userData, _ := l.svcCtx.UserModel.FindOneByUserIdentity(l.ctx, in.UserIdentity)
	if userData == nil {
		return nil, errorx.NewDefaultError("找不到该用户")
	}

	userJsonData, JsonErr := json.Marshal(userData)
	if JsonErr != nil {
		return nil, errorx.NewDefaultError("数据转换失败")
	}
	return &account.CommonResply{
		Code:    http.StatusOK,
		Message: "获取成功",
		Data:    string(userJsonData),
	}, nil
}
