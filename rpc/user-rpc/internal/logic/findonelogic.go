package logic

import (
	"context"

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

func (l *FindOneLogic) FindOne(in *user.ReqId) (*user.RespUser, error) {
	// todo: add your logic here and delete this line

	return &user.RespUser{}, nil
}
