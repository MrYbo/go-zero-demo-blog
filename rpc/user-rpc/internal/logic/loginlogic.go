package logic

import (
	"context"
	"github.com/blog/v1/model"
	"github.com/blog/v1/utils"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"

	"github.com/blog/v1/rpc/user-rpc/internal/svc"
	user "github.com/blog/v1/rpc/user-rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.ReqLogin) (*user.BaseUser, error) {
	if len(strings.TrimSpace(in.Username)) == 0 || len(strings.TrimSpace(in.Password)) == 0 {
		return nil, status.Error(http.StatusBadRequest, "用户名或者密码错误")
	}
	var isUser model.Users
	first := l.svcCtx.UserModel.Where("username=?", in.Username).First(&isUser)
	if first.RowsAffected == 0 || !utils.CheckPassword(in.Password, isUser.Password) {
		return nil, status.Error(http.StatusUnauthorized, "用户名或者密码错误")
	}
	return &user.BaseUser{
		Id:       int64(isUser.ID),
		Username: isUser.Username,
		Avatar:   isUser.Avatar,
		Name:     isUser.Name,
	}, nil
}
