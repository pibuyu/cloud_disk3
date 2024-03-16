package logic

import (
	"cloud_disk3/core/helper"
	models2 "cloud_disk3/core/models"
	"context"
	"errors"

	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	email := req.Email
	name := req.Name
	password := req.Password
	inputCode := req.Code

	if email == "" || name == "" || password == "" || inputCode == "" {
		return nil, errors.New("参数不能为空")
	}

	sysCode := l.svcCtx.RDB.Get(l.ctx, email).Val() //从redis中获取验证码

	if sysCode != inputCode {
		return nil, errors.New("验证码错误")
	}

	user := &models2.UserBasic{
		Identity: helper.GetUUID(),
		Name:     name,
		Password: password,
		Email:    email,
	}
	insert, err := l.svcCtx.Engine.Insert(user)
	if err != nil || insert == 0 {
		return nil, errors.New("注册失败")
	}

	resp = new(types.UserRegisterReply)
	resp.Message = "注册成功"
	return
}
