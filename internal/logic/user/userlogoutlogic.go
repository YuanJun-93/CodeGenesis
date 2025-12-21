package user

import (
	"context"

	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogoutLogic {
	return &UserLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogoutLogic) UserLogout() (resp *types.BaseResponseBoolean, err error) {
	return &types.BaseResponseBoolean{
		Code: 0,
		Data: true,
		Msg:  "ok",
	}, nil
}
