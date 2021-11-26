package svc

import (
	"github.com/blog/v1/api/user-api/internal/config"
	"github.com/blog/v1/rpc/user-rpc/users"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// users.Users 是 user rpc 服务对外暴露的接口
	Users users.Users
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//  zrpc.MustNewClient(c.User) 创建了一个 grpc 客户端
		Users: users.NewUsers(zrpc.MustNewClient(c.User)),
	}
}
