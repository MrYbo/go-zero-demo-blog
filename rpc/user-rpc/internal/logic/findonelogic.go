package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/blog/v1/rpc/user-rpc/internal/svc"
	user "github.com/blog/v1/rpc/user-rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneLogic {
	return &FindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOneLogic) FindOne(in *user.UserId) (*user.BaseUser, error) {
	var fUser user.BaseUser
	if row := l.svcCtx.UserModel.Where("id=?", in.Id).First(&fUser).RowsAffected; row == 0 {
		return nil, status.Error(http.StatusNotFound, "not found")
	}
	return &user.BaseUser{
		Id:       fUser.Id,
		Username: fUser.Username,
		Avatar:   fUser.Avatar,
		Name:     fUser.Name,
		Phone:    fUser.Phone,
		Address:  fUser.Address,
		Birthday: fUser.Birthday,
	}, nil
}
