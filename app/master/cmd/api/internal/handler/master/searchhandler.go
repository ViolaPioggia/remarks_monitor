package master

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"remarks_monitor/app/master/cmd/api/internal/logic/master"
	"remarks_monitor/app/master/cmd/api/internal/svc"
	"remarks_monitor/app/master/cmd/api/internal/types"
)

func SearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := master.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
