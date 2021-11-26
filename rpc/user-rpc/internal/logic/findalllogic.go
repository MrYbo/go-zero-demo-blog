package logic

import (
	"context"

	"github.com/blog/v1/rpc/user-rpc/internal/svc"
	user "github.com/blog/v1/rpc/user-rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllLogic {
	return &FindAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindAllLogic) FindAll(in *user.ReqFindAll) (*user.RespFindAll, error) {
	// todo: add your logic here and delete this line

	return &user.RespFindAll{}, nil
}
