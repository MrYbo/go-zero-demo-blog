package svc

import (
	"fmt"
	"github.com/blog/v1/model"
	"github.com/blog/v1/rpc/user-rpc/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("数据库初始化失败，%s", err))
	}

	db.AutoMigrate(&model.Users{})
	return &ServiceContext{
		Config: c,
		UserModel: db,
	}
}
