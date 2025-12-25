package user

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/YuanJun-93/CodeGenesis/internal/model"
	"github.com/YuanJun-93/CodeGenesis/internal/svc"
	"github.com/YuanJun-93/CodeGenesis/internal/types"
	"github.com/go-playground/validator/v10"

	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.BaseResponseLong, err error) {
	// 1. Validate params using Validator
	if err := l.svcCtx.Validator.Struct(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return nil, fmt.Errorf("params error: %s", errs[0].Translate(l.svcCtx.Trans))
		}
		return nil, fmt.Errorf("params error: %v", err)
	}

	// Log with Zap
	l.svcCtx.ZapLogger.Info("UserRegister called", zap.String("account", req.UserAccount))

	if req.UserPassword != req.CheckPassword {
		return nil, errors.New("params error: two passwords are inconsistent")
	}

	// 2. Check if account exists
	_, err = l.svcCtx.UserModel.FindOneByUserAccount(l.ctx, req.UserAccount)
	if err == nil {
		return nil, errors.New("params error: account already exists")
	} else if err != model.ErrNotFound {
		return nil, err
	}

	// 3. Encrypt password (Simple MD5 for demo parity, ideally use bcrypt)
	// For strict parity with Java "DigestUtils.md5DigestAsHex(SALT + password)"
	// We will assume a simple encryption here or implement a helper.
	// Using a simple salt here.
	const salt = "Louis"
	encryptedPassword := fmt.Sprintf("%x", md5.Sum([]byte(salt+req.UserPassword)))

	// 4. Insert User
	newUser := &model.User{
		UserAccount:  req.UserAccount,
		UserPassword: encryptedPassword,
		UserName:     "User" + req.UserAccount,
		UserRole:     "user",
	}

	err = l.svcCtx.UserModel.Insert(l.ctx, newUser)
	if err != nil {
		return nil, err
	}

	userId := newUser.ID

	return &types.BaseResponseLong{
		Code: 0,
		Data: userId,
		Msg:  "ok",
	}, nil
}
