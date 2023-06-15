package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"remarks_monitor/app/usercenter/cmd/api/internal/logic/user"
	"remarks_monitor/app/usercenter/cmd/api/internal/svc"
	"remarks_monitor/app/usercenter/cmd/api/internal/types"
)

func AddAdminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddAdminReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewAddAdminLogic(r.Context(), svcCtx)
		resp, err := l.AddAdmin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
