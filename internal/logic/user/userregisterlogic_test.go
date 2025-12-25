package user

import (
	"context"
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

func TestUserRegisterLogic_UserRegister(t *testing.T) {
	tests := []struct {
		name    string
		req     *types.UserRegisterRequest
		setup   func(m *mocks.MockUserModel)
		wantErr bool
	}{
		{
			name: "Success",
			req: &types.UserRegisterRequest{
				UserAccount:   "testuser",
				UserPassword:  "password123",
				CheckPassword: "password123",
			},
			setup: func(m *mocks.MockUserModel) {
				m.EXPECT().FindOneByUserAccount(gomock.Any(), "testuser").Return(nil, model.ErrNotFound)
				m.EXPECT().Insert(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, u *model.User) {
					u.ID = 1
				}).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "AccountAlreadyExists",
			req: &types.UserRegisterRequest{
				UserAccount:   "existing",
				UserPassword:  "password123",
				CheckPassword: "password123",
			},
			setup: func(m *mocks.MockUserModel) {
				m.EXPECT().FindOneByUserAccount(gomock.Any(), "existing").Return(&model.User{ID: 1, UserAccount: "existing"}, nil)
			},
			wantErr: true,
		},
		{
			name: "PasswordMismatch",
			req: &types.UserRegisterRequest{
				UserAccount:   "testuser",
				UserPassword:  "password123",
				CheckPassword: "password456",
			},
			setup:   func(m *mocks.MockUserModel) {},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockModel := mocks.NewMockUserModel(ctrl)
			if tt.setup != nil {
				tt.setup(mockModel)
			}

			val, trans := valInit.Init()
			svcCtx := &svc.ServiceContext{
				Config:    config.Config{},
				UserModel: mockModel,
				Validator: val,
				Trans:     trans,
				ZapLogger: zap.NewNop(), // Use Nop for tests
			}
			l := NewUserRegisterLogic(context.Background(), svcCtx)
			// Disable logging for tests
			l.Logger = logx.WithContext(context.Background())

			_, err := l.UserRegister(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRegisterLogic.UserRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
