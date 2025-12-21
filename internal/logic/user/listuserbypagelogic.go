package user

import (
	"context"

	"github.com/YuanJun-93/CodeGenesis/internal/model"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserByPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserByPageLogic {
	return &ListUserByPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserByPageLogic) ListUserByPage(req *types.UserQueryRequest) (resp *types.BaseResponseUserPage, err error) {
	db := l.svcCtx.SqlConn.Model(&model.User{})

	// Filter
	if req.UserName != "" {
		db = db.Where("userName LIKE ?", "%"+req.UserName+"%")
	}
	if req.UserProfile != "" {
		db = db.Where("userProfile LIKE ?", "%"+req.UserProfile+"%")
	}
	if req.UserRole != "" {
		db = db.Where("userRole = ?", req.UserRole)
	}

	// Count
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	// Pagination
	var users []*model.User
	offset := (req.Current - 1) * req.PageSize
	if err := db.Limit(int(req.PageSize)).Offset(int(offset)).Find(&users).Error; err != nil {
		return nil, err
	}

	// Model -> Response
	var records []types.UserResponse
	for _, user := range users {
		records = append(records, types.UserResponse{
			Id:          user.Id,
			UserAccount: user.UserAccount,
			UserName:    user.UserName.String,
			UserAvatar:  user.UserAvatar.String,
			UserProfile: user.UserProfile.String,
			UserRole:    user.UserRole,
			CreateTime:  user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  user.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.BaseResponseUserPage{
		Code: 0,
		Msg:  "ok",
		Data: types.UserPage{
			Records: records,
			Total:   total,
		},
	}, nil
}
