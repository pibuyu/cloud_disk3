package logic

import (
	"cloud_disk3/core/helper"
	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"
	"cloud_disk3/core/models"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	//从数据库查询当前用户
	user := &models.UserBasic{}
	has, err := l.svcCtx.Engine.Where("name = ? and password = ?", req.Name, req.Password).Get(user)
	if err != nil {
		return nil, errors.New("查询用户出错")
	}
	if !has {
		return nil, errors.New("用户不存在")
	}

	//返回token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, errors.New("生成token出错")
	}

	resp = new(types.LoginReply)
	resp.Token = token
	return
}
