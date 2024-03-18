package logic

import (
	"cloud_disk3/core/models"
	"context"
	"errors"

	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	//修改文件名称之前需要判断一下同一个parent_id下是否存在同名文件，如果存在，拒绝这次修改
	count, err := l.svcCtx.Engine.Where("name = ? AND parent_id=(SELECT parent_id FROM user_repository ur WHERE ur.identity = ?)", req.Name, req.Identity).Count(&models.UserRepository{})
	if err != nil {
		return
	}
	if count > 0 {
		return nil, errors.New("该文件夹下已存在同名文件")
	}

	//修改之前也需要按照文件的identity查询该文件是否存在
	cnt, err := l.svcCtx.Engine.Where("identity = ? ", req.Identity).Count(&models.UserRepository{})
	if err != nil {
		return
	}
	if cnt == 0 {
		return nil, errors.New("要修改的文件不存在")
	}

	//既不存在同名文件，同时要修改的文件也存在，进行修改
	data := &models.UserRepository{Name: req.Name}
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(data)
	if err != nil {
		return
	}
	resp = new(types.UserFileNameUpdateReply)
	resp.Message = "修改文件名称成功"
	return
}
