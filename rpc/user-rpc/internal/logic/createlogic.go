package logic

import (
	"context"
	"github.com/blog/v1/model"
	"google.golang.org/grpc/status"
	"net/http"
	"time"

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

func (l *CreateLogic) Create(in *user.UserParams) (*user.BaseUser, error) {

	first := l.svcCtx.UserModel.Where("username=?", in.Username).First(&model.Users{})
	if first.RowsAffected != 0 {
		return nil, status.Error(http.StatusConflict, "用户名已存在")
	}

	birthday, _ := time.ParseInLocation("2006-01-02 15:04:05", in.Birthday, time.Local)
	_user := &model.Users{
		Username: in.Username,
		Password: in.Password,
		Avatar:   in.Avatar,
		Name:     in.Name,
		Phone:    in.Phone,
		Address:  in.Address,
		Birthday: &birthday,
	}
	l.svcCtx.UserModel.Create(&_user)
	return &user.BaseUser{
		Id:        int64(_user.ID),
		Username:  in.Username,
		Avatar:    in.Avatar,
		Name:      in.Name,
		Phone:     in.Phone,
		Address:   in.Address,
		Birthday:  in.Birthday,
		CreatedAt: _user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: _user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
