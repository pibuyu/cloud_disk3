package handler

import (
	"net/http"

	"cloud_disk3/core/internal/logic"
	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDeatilHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDeatilRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserDeatilLogic(r.Context(), svcCtx)
		resp, err := l.UserDeatil(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
