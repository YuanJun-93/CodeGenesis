package user

import (
	"net/http"

	"github.com/YuanJun-93/CodeGenesis/internal/logic/user"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListUserByPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserQueryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewListUserByPageLogic(r.Context(), svcCtx)
		resp, err := l.ListUserByPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
