package logic

import (
	"cloud_disk3/core/define"
	"cloud_disk3/core/helper"
	models2 "cloud_disk3/core/models"

	"context"
	"errors"
	"time"

	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeRegisterLogic {
	return &SendCodeRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeRegisterLogic) SendCodeRegister(req *types.SendCodeRequest) (resp *types.SendCodeReply, err error) {
	//检查该邮箱是否已被注册
	count, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(&models2.UserBasic{})
	if err != nil {
		return nil, errors.New("查询邮箱是否已被注册时出错")
	}
	if count > 0 {
		return nil, errors.New("该邮箱已被注册")
	}
	//没被注册，走正常注册流程
	resp = new(types.SendCodeReply)
	toEmail := req.Email
	code := helper.GenerateCode()

	l.svcCtx.RDB.Set(l.ctx, toEmail, code, time.Duration(define.CODE_EXPIRE)*time.Second) //验证码写入redis
	err = helper.SendVerifyCode(toEmail, code)
	if err != nil {
		resp.Message = "发送邮件失败"
		return nil, errors.New("发送邮件失败")
	}

	resp.Message = "发送邮件成功,目的邮箱：" + toEmail
	return
}
