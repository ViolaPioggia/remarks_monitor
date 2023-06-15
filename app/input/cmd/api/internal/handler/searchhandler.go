package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"remarks_monitor/app/input/cmd/api/internal/logic"
	"remarks_monitor/app/input/cmd/api/internal/svc"
	"remarks_monitor/app/input/cmd/api/internal/types"
)

func searchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
