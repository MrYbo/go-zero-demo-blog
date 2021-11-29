package logic

import (
	"context"
	"github.com/blog/v1/model"
	"github.com/blog/v1/rpc/user-rpc/internal/svc"
	user "github.com/blog/v1/rpc/user-rpc/user"
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
)

type DestroyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDestroyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DestroyLogic {
	return &DestroyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DestroyLogic) Destroy(in *user.UserId) (*user.CommonOK, error) {
	l.svcCtx.UserModel.Debug().Delete(&model.Users{}, in.Id)
	return &user.CommonOK{Code: http.StatusOK}, nil
}
