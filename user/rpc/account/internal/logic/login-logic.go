package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"microShop/comm"
	"microShop/user/rpc/account/internal/svc"
	"microShop/user/rpc/account/model"
	"microShop/user/rpc/account/types/account"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *account.LoginReq) (*account.CommonResply, error) {
	// 获取用户信息
	userData, _ := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.UserName)
	if userData == nil {
		return nil, errors.New("用户未注册")
	}
	// 密码解密
	dePassWord, dePassErr := comm.Decrypt(userData.PassWord, []byte(l.svcCtx.Config.SecretKey))
	if dePassErr != nil {
		return nil, dePassErr
	}
	if in.PassWord != dePassWord {
		return nil, errors.New("密码错误")
	}

	// 获取客户端IP
	ip, _ := comm.ExternalIp()

	user := model.User{
		Id:           userData.Id,
		UserIdentity: userData.UserIdentity,
		UserName:     userData.UserName,
		PassWord:     userData.PassWord,
		UserNick:     userData.UserNick,
		UserFace:     userData.UserFace,
		UserSex:      userData.UserSex,
		UserEmail:    userData.UserEmail,
		UserPhone:    userData.UserPhone,
		LoginAddress: ip.String(),
	}

	cntErr := l.svcCtx.UserModel.Update(l.ctx, &user)
	if cntErr != nil {
		return nil, cntErr
	}

	// jwt
	payloads := make(map[string]any)
	payloads["userId"] = userData.Id

	accessToken, tokenErr := l.GetToken(time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessSecret, payloads, l.svcCtx.Config.JwtAuth.AccessExpire)
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &account.CommonResply{
		Code:    0,
		Message: "登录成功",
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
