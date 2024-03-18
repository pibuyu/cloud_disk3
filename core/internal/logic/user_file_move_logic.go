package logic

import (
	"cloud_disk3/core/models"
	"context"
	"errors"

	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	//先根据文件的idnetity查询这个文件是否存在
	count, err := l.svcCtx.Engine.Where("identity = ?", req.Identity).Count(&models.UserRepository{})
	if err != nil {
		return
	}
	if count == 0 {
		return nil, errors.New("要修改的文件不存在")
	}

	//开始修改parent_id
	data := &models.UserRepository{
		ParentId: req.ParentId,
	}
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(data)
	if err != nil {
		return nil, errors.New("修改文件的parent_id失败")
	}
	resp = new(types.UserFileMoveReply)
	resp.Message = "修改文件的parent_id成功"
	return
}
