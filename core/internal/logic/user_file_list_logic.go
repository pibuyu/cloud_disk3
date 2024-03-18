package logic

import (
	"cloud_disk3/core/define"
	"cloud_disk3/core/models"
	"context"
	"time"

	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//Id                 int
//Identity           string
//UserIdentity       string
//ParentId           int64
//RepositoryIdentity string
//Ext                string
//Name               string

//	Name               string `json:"name"`
//	Size               int    `json:"size"`
//	Ext                string `json:"ext"`
//	Path               string `json:"path"`

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {

	uf := make([]*types.UserFile, 0)
	var count int64
	pageSize := req.Size
	if pageSize == 0 {
		pageSize = define.PAGE_SIZE
	}
	pageNum := req.Page
	if pageNum == 0 {
		pageNum = define.PAGE_NUM
	}
	offset := (pageNum - 1) * pageSize
	resp = new(types.UserFileListReply)

	//查询所有文件信息
	err = l.svcCtx.Engine.Table("user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id,user_repository.identity,user_repository.repository_identity,user_repository.ext,user_repository.name,"+
																	"repository_pool.path,repository_pool.size").
		Where("user_repository.deleted_at= ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.TimeFormat)). //原生的sql里没有排除掉软删除记录，手动加上
		Join("LEFT", "repository_pool", "repository_pool.identity = user_repository.repository_identity").
		Limit(define.PAGE_SIZE, offset).Find(&uf)
	if err != nil {
		return
	}
	//统计文件数目
	count, err = l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).Count(new(models.UserRepository))
	if err != nil {
		return
	}

	resp.List = uf
	resp.Count = int(count)
	return
}
