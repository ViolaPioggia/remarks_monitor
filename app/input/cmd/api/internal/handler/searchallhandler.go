package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"remarks_monitor/app/input/cmd/api/internal/logic"
	"remarks_monitor/app/input/cmd/api/internal/svc"
	"remarks_monitor/app/input/cmd/api/internal/types"
)

func searchAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchAllRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSearchAllLogic(r.Context(), svcCtx)
		resp, err := l.SearchAll(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
