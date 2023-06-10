package master

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	_ "net/http/pprof"
	"remarks_monitor/app/master/cmd/api/internal/logic/master"
	"remarks_monitor/app/master/cmd/api/internal/svc"
	"remarks_monitor/app/master/cmd/api/internal/types"
)

func CoordinatorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MasterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := master.NewCoordinatorLogic(r.Context(), svcCtx)
		resp, err := l.Coordinator(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
