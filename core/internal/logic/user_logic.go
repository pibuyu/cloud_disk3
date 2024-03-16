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

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LoginRequest) (resp *types.LoginReply, err error) {

	//从数据库查询当前用户
	user := new(models2.UserBasic)
	//Get后面指定查哪张表
	has, err := l.svcCtx.Engine.Where("name = ? and password = ?", req.Name, req.Password).Get(&models2.UserBasic{})
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
