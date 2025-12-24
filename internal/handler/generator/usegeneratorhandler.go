package generator

import (
	"net/http"

	"github.com/YuanJun-93/CodeGenesis/internal/logic/generator"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UseGeneratorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GeneratorRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := generator.NewUseGeneratorLogic(r.Context(), svcCtx)
		resp, err := l.UseGenerator(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
