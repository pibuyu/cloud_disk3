package logic

import (
	"cloud_disk3/core/helper"
	"cloud_disk3/core/models"
	"context"

	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadReply, err error) {
	rp := &models.RepositoryPool{
		Identity: helper.GetUUID(),
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     req.Path,
	}
	insert, err := l.svcCtx.Engine.Insert(rp)
	if err != nil || insert == 0 {
		return nil, err
	}

	resp = new(types.FileUploadReply)
	resp.Identity = rp.Identity
	resp.Message = "upload success!"

	return
}
