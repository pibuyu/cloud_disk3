package handler

import (
	"net/http"

	"cloud_disk3/core/internal/logic"
	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRepositorySaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRepositorySaveRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserRepositorySaveLogic(r.Context(), svcCtx)

		//todo:r.Header.Get("UserIdentity") 没拿到UserIdentity，不清楚什么原因
		resp, err := l.UserRepositorySave(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
