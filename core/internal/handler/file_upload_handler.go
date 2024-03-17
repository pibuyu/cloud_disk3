package handler

import (
	"cloud_disk3/core/helper"
	"cloud_disk3/core/models"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud_disk3/core/internal/logic"
	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//拿到文件的hash、name、ext、size和path并放入req中,接下来在file_upload_logic.go中处理
		file, header, err := r.FormFile("file")
		if err != nil {
			return
		}

		//计算文件的hash值，用的是md5.Sum方法
		b := make([]byte, header.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))

		//判断文件是否已存在
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			return
		}
		if has {
			httpx.OkJson(w, &types.FileUploadReply{
				Message:  "file already exists!",
				Identity: rp.Identity,
				Ext:      rp.Ext,
				Name:     rp.Name,
			})
			return
		}

		//把文件存到cos中
		cosPath, err := helper.UploadFile(r)
		req.Name = header.Filename
		req.Ext = path.Ext(header.Filename)
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
