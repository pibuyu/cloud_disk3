package logic

import (
	"cloud_disk3/core/helper"
	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"
	"cloud_disk3/core/models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, userIdentity string) (resp *types.UserRepositorySaveReply, err error) {

	ur := &models.UserRepository{
		Identity:           helper.GetUUID(),
		ParentId:           req.ParentId,
		UserIdentity:       userIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
		RepositoryIdentity: req.RepositoryIdentity,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return nil, err
	}

	return
}
