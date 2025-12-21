package user

import (
	"net/http"

	"github.com/YuanJun-93/CodeGenesis/internal/logic/user"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetLoginUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetLoginUserLogic(r.Context(), svcCtx)
		resp, err := l.GetLoginUser()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
