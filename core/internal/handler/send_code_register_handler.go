package handler

import (
	"net/http"

	"cloud_disk3/core/internal/logic"
	"cloud_disk3/core/internal/svc"
	"cloud_disk3/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendCodeRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendCodeRegisterLogic(r.Context(), svcCtx)
		resp, err := l.SendCodeRegister(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
