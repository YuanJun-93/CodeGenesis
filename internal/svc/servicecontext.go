package svc

import (
	"github.com/YuanJun-93/CodeGenesis/internal/config"
	"github.com/YuanJun-93/CodeGenesis/internal/middleware"
	"github.com/YuanJun-93/CodeGenesis/internal/model"
	logInit "github.com/YuanJun-93/CodeGenesis/internal/pkg/log"
	valInit "github.com/YuanJun-93/CodeGenesis/internal/pkg/validator"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	SqlConn     *gorm.DB
	RedisClient *redis.Redis
	UserModel   model.UserModel
	AdminCheck  rest.Middleware
	ZapLogger   *zap.Logger
	Validator   *validator.Validate
	Trans       ut.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds, _ := redis.NewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Type: c.Redis.Type,
		Pass: c.Redis.Pass,
	})

	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Init Logger
	logger := logInit.Init(c.Log)

	// Init Validator
	val, trans := valInit.Init()

	return &ServiceContext{
		Config:      c,
		SqlConn:     db,
		RedisClient: rds,
		UserModel:   model.NewUserModel(db),
		AdminCheck:  middleware.NewAdminCheckMiddleware().Handle,
		ZapLogger:   logger,
		Validator:   val,
		Trans:       trans,
	}
}
