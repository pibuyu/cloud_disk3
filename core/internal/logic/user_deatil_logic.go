package logic

import (
	models2 "cloud_disk3/core/models"

	"context"
	"errors"

	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeatilLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeatilLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeatilLogic {
	return &UserDeatilLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeatilLogic) UserDeatil(req *types.UserDeatilRequest) (resp *types.UserDeatilReply, err error) {

	identity := req.Identity
	userBasic := new(models2.UserBasic)

	has, err := l.svcCtx.Engine.Where("identity = ?", identity).Get(userBasic)
	if err != nil {
		return nil, errors.New("查询用户出错")
	}
	if !has {
		return nil, errors.New("用户不存在")
	}

	resp = new(types.UserDeatilReply)
	resp.Email = userBasic.Email
	resp.Name = userBasic.Name
	return
}
