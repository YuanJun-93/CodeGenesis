package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/YuanJun-93/CodeGenesis/internal/model"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginUserLogic {
	return &GetLoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginUserLogic) GetLoginUser() (resp *types.BaseResponseLoginUserResponse, err error) {
	userId := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	id, err := userId.Int64()
	if err != nil {
		return nil, err
	}

	// 2. Query User
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// 3. Construct Response
	return &types.BaseResponseLoginUserResponse{
		Code: 0,
		Msg:  "success",
		Data: types.LoginUserResponse{
			Id:          user.ID,
			UserAccount: user.UserAccount,
			UserName:    user.UserName,
			UserAvatar:  user.UserAvatar,
			UserProfile: user.UserProfile,
			UserRole:    user.UserRole,
			CreateTime:  user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  user.UpdateTime.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
