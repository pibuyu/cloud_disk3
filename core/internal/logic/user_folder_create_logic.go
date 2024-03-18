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

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {
	//创建文件夹之前需要判断一下同一个parent_id下是否存在同名文件夹，如果存在，拒绝这次创建请求
	count, err := l.svcCtx.Engine.Where("name = ? AND parent_id= ? ", req.Name, req.ParentId).Count(&models.UserRepository{})
	if err != nil {
		return
	}
	if count > 0 {
		return nil, errors.New("该文件夹下已存在同名文件")
	}
	//创建文件夹
	data := &models.UserRepository{
		Identity:     helper.GetUUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return
	}
	resp = new(types.UserFolderCreateReply)
	resp.Message = "创建文件夹成功"
	return
}
