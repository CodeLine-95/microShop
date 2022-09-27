package logic

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"microShop/comm"
	"microShop/user/rpc/account/internal/svc"
	"microShop/user/rpc/account/model"
	"microShop/user/rpc/account/types/account"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *account.RegisterReq) (*account.CommonResply, error) {
	// todo: add your logic here and delete this line
	userData, _ := l.svcCtx.UserModel.FindOneByUserPhone(l.ctx, in.UserName)
	if userData != nil {
		return nil, errors.New("用户名已注册")
	}

	if len(in.PassWord) < 6 || len(in.PassWord) > 16 {
		return nil, errors.New("密码之能是6~16位")
	}

	if ok := comm.CheckPassword(in.PassWord); !ok {
		return nil, errors.New("密码格式错误")
	}

	if ok := comm.CheckUsername(in.UserName); !ok {
		return nil, errors.New("用户名格式错误")
	}

	enPassWord, enPassErr := comm.Encrypt(in.PassWord, []byte(l.svcCtx.Config.SecretKey))
	if enPassErr != nil {
		return nil, enPassErr
	}

	newUUID, _ := uuid.NewUUID()
	user := model.User{
		UserIdentity: newUUID.String(),
		UserName:     in.UserName,
		PassWord:     enPassWord,
		UserNick:     in.UserNick,
		UserFace:     in.UserFace,
		UserSex:      in.UserSex,
		UserEmail:    in.UserEmail,
		UserPhone:    "",
		LoginAddress: "",
	}
	cnt, cntErr := l.svcCtx.UserModel.Insert(l.ctx, &user)

	if cntErr != nil {
		return nil, cntErr
	}

	rowsAffected, rowsErr := cnt.RowsAffected()
	if rowsErr != nil {
		return nil, rowsErr
	}

	if rowsAffected == 0 {
		return nil, errors.New("注册失败")
	}

	return &account.CommonResply{
		Code:    0,
		Message: "注册成功",
	}, nil
}
