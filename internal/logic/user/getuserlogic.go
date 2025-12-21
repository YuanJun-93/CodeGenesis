package user

import (
	"context"

	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.IdPathRequest) (resp *types.BaseResponseUserResponse, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.BaseResponseUserResponse{
		Code: 0,
		Msg:  "ok",
		Data: types.UserResponse{
			Id:          user.Id,
			UserAccount: user.UserAccount,
			UserName:    user.UserName.String,
			UserAvatar:  user.UserAvatar.String,
			UserProfile: user.UserProfile.String,
			UserRole:    user.UserRole,
			CreateTime:  user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  user.UpdateTime.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
