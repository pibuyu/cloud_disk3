package logic

import (
	"context"

	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDeatilLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDeatilLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDeatilLogic {
	return &ShareBasicDeatilLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDeatilLogic) ShareBasicDeatil(req *types.ShareBasicDeatilRequest) (resp *types.ShareBasicDeatilReply, err error) {
	// 更新分享记录的点击次数+1
	_, err = l.svcCtx.Engine.Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?", req.Identity)
	if err != nil {
		return nil, err
	}

	//因为req的参数是share_basic表的repository_identity，详细信息在repository_pool中，
	//因此需要将share_basic.repository_identity与详细信息在repotisory_pool.identity连接
	resp = new(types.ShareBasicDeatilReply)
	_, err = l.svcCtx.Engine.Table("share_basic").Select("share_basic.repository_identity,repository_pool.name,repository_pool.ext,repository_pool.path,repository_pool.size").
		Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Where("share_basic.identity = ?", req.Identity).
		Get(resp)
	if err != nil {
		return nil, err
	}
	return
}
