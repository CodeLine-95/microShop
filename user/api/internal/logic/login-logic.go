package logic

import (
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"microShop/user/api/internal/svc"
	"microShop/user/api/internal/types"
	"microShop/user/rpc/account/model"
	"microShop/user/rpc/account/user"
	"time"
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

	var userData model.User

	userErr := json.Unmarshal([]byte(cnt.Data), &userData)
	if userErr != nil {
		return nil, userErr
	}

	// jwt
	payloads := make(map[string]any)
	payloads["userIdentity"] = userData.UserIdentity

	accessToken, tokenErr := l.GetToken(time.Now().Unix(), l.svcCtx.Config.Auth.AccessSecret, payloads, l.svcCtx.Config.Auth.AccessExpire)
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &types.CommonResply{
		Code:    cnt.Code,
		Message: cnt.Message,
		Data:    accessToken,
	}, nil
}

func (l *LoginLogic) GetToken(iat int64, secretKey string, payloads map[string]any, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["expTime"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
