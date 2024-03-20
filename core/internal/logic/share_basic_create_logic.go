package logic

import (
	"cloud_disk3/core/helper"
	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"
	"cloud_disk3/core/models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	resp = new(types.ShareBasicCreateReply)

	//如果这个identity的文件已经被分享过了，就执行更新操作，click_num+1
	//sb:=&models.ShareBasic{}
	//_, err = l.svcCtx.Engine.Where("repository_identity = ? AND user_identity = ?", req.RepositoryIdentity, userIdentity).Get(sb)
	//if err != nil {
	//	return nil,errors.New("查询是否存在文件分享记录失败")
	//}
	//if sb.ClickNum != 0 {
	//	sb.ClickNum += 1
	//	resp.Message="文件分享成功，点击数+1"
	//	return
	//}

	//不然就执行create操作
	identity := helper.GetUUID()
	data := &models.ShareBasic{
		Identity:           identity,
		UserIdentity:       userIdentity,
		RepositoryIdentity: req.RepositoryIdentity,
		ExpireTime:         req.ExpireTime,
	}
	logx.Info(data)
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return nil, err
	}
	resp.Message = "文件分享成功"
	resp.Identity = identity
	return
}
