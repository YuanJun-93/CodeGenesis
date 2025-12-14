package system

import (
	"net/http"

	"github.com/YuanJun-93/CodeGenesis/internal/logic/system"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthzHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := system.NewHealthzLogic(r.Context(), svcCtx)
		resp, err := l.Healthz()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
