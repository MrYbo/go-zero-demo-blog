package main

import (
	"flag"
	"fmt"
	"github.com/blog/v1/common/errorx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"

	"github.com/blog/v1/rpc/user-rpc/internal/config"
	"github.com/blog/v1/rpc/user-rpc/internal/server"
	"github.com/blog/v1/rpc/user-rpc/internal/svc"
	user "github.com/blog/v1/rpc/user-rpc/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUsersServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUsersServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()


	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
