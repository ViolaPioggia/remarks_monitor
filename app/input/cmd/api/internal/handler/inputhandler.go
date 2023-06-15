package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"remarks_monitor/app/input/cmd/api/internal/logic"
	"remarks_monitor/app/input/cmd/api/internal/svc"
	"remarks_monitor/app/input/cmd/api/internal/types"
)

func inputHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InputRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewInputLogic(r.Context(), svcCtx)
		resp, err := l.Input(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
