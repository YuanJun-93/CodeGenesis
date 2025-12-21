package user

import (
	"context"
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/YuanJun-93/CodeGenesis/internal/config"
	"github.com/YuanJun-93/CodeGenesis/internal/mocks"
	"github.com/YuanJun-93/CodeGenesis/internal/model"
	valInit "github.com/YuanJun-93/CodeGenesis/internal/pkg/validator"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"
	"github.com/golang/mock/gomock"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

func TestUserLoginLogic_UserLogin(t *testing.T) {
	const salt = "Louis"

	tests := []struct {
		name    string
		req     *types.UserLoginRequest
		setup   func(m *mocks.MockUserModel)
		wantErr bool
	}{
		{
			name: "Success",
			req: &types.UserLoginRequest{
				UserAccount:  "testuser",
				UserPassword: "password123",
			},
			setup: func(m *mocks.MockUserModel) {
				encrypted := fmt.Sprintf("%x", md5.Sum([]byte(salt+"password123")))
				m.EXPECT().FindOneByUserAccount(gomock.Any(), "testuser").Return(&model.User{
					Id:           1,
					UserAccount:  "testuser",
					UserPassword: encrypted,
					UserRole:     "user",
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "UserNotFound",
			req: &types.UserLoginRequest{
				UserAccount:  "unknown",
				UserPassword: "password123",
			},
			setup: func(m *mocks.MockUserModel) {
				m.EXPECT().FindOneByUserAccount(gomock.Any(), "unknown").Return(nil, model.ErrNotFound)
			},
			wantErr: true,
		},
		{
			name: "WrongPassword",
			req: &types.UserLoginRequest{
				UserAccount:  "testuser",
				UserPassword: "wrongpassword",
			},
			setup: func(m *mocks.MockUserModel) {
				encrypted := fmt.Sprintf("%x", md5.Sum([]byte(salt+"password123")))
				m.EXPECT().FindOneByUserAccount(gomock.Any(), "testuser").Return(&model.User{
					Id:           1,
					UserAccount:  "testuser",
					UserPassword: encrypted,
				}, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			val, trans := valInit.Init()
			mockModel := mocks.NewMockUserModel(ctrl)
			if tt.setup != nil {
				tt.setup(mockModel)
			}

			svcCtx := &svc.ServiceContext{
				Config: config.Config{
					Auth: struct {
						AccessSecret string
						AccessExpire int64
					}{
						AccessSecret: "test_secret",
						AccessExpire: 3600,
					},
				},
				UserModel: mockModel,
				Validator: val,
				Trans:     trans,
				ZapLogger: zap.NewNop(),
			}
			l := NewUserLoginLogic(context.Background(), svcCtx)
			l.Logger = logx.WithContext(context.Background())

			_, err := l.UserLogin(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserLoginLogic.UserLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
