package logic

import (
	"context"
	"github.com/blog/v1/model"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/blog/v1/rpc/user-rpc/internal/svc"
	user "github.com/blog/v1/rpc/user-rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.ReqCreate) (*user.RespUser, error) {


	first := l.svcCtx.UserModel.Where("username=?", in.Username).First(&model.Users{})
	if first.RowsAffected != 0 {
		return nil, status.Error(http.StatusConflict, "用户名已存在")
	}

	l.svcCtx.UserModel.Create(&model.Users{
		Username: in.Username,
		Password: in.Password,
		Avatar:   in.Avatar,
		Name:     in.Name,
		Phone:    in.Phone,
		Address:  in.Address,
		Birthday: in.Birthday,
	})

	return &user.RespUser{
		Username: in.Username,
		Avatar: in.Avatar,
		Name: in.Name,
		Phone: in.Phone,
		Address: in.Address,
		Birthday: in.Birthday,
	}, nil
}
