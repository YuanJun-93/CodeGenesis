package user

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"github.com/YuanJun-93/CodeGenesis/internal/model"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"
	"github.com/go-playground/validator/v10"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.BaseResponseLoginUserResponse, err error) {
	// 1. Params Validate
	if err := l.svcCtx.Validator.Struct(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return nil, fmt.Errorf("params error: %s", errs[0].Translate(l.svcCtx.Trans))
		}
		return nil, fmt.Errorf("params error: %v", err)
	}

	user, err := l.svcCtx.UserModel.FindOneByUserAccount(l.ctx, req.UserAccount)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if encryptPassword(req.UserPassword) != user.UserPassword {
		return nil, errors.New("wrong password")
	}

	// Generate Token
	now := time.Now().Unix()
	accessSecret := l.svcCtx.Config.Auth.AccessSecret
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := l.getJwtToken(accessSecret, now, accessExpire, user.ID)
	if err != nil {
		return nil, err
	}

	return &types.BaseResponseLoginUserResponse{
		Code: 0,
		Msg:  "Login Success",
		Data: types.LoginUserResponse{
			Id:          user.ID,
			UserAccount: user.UserAccount,
			UserName:    user.UserName,
			UserAvatar:  user.UserAvatar,
			UserProfile: user.UserProfile,
			UserRole:    user.UserRole,
			CreateTime:  user.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  user.UpdateTime.Format("2006-01-02 15:04:05"),
			Token:       token,
		},
	}, nil
}

// encryptPassword is a placeholder for the actual password encryption logic.
// This function should be defined elsewhere in the package or imported.
func encryptPassword(password string) string {
	const salt = "Louis" // This salt should ideally be unique per user or a strong global secret
	return fmt.Sprintf("%x", md5.Sum([]byte(salt+password)))
}

// getJwtToken generates a JWT token for the given user ID.
func (l *UserLoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	// Assuming user.UserRole is available and passed or retrieved here if needed for claims
	// claims["userRole"] = user.UserRole // Important for Admin Middleware

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
