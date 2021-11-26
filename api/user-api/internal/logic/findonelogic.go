package logic

import (
	"context"

	"github.com/blog/v1/api/user-api/internal/svc"
	"github.com/blog/v1/api/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindOneLogic {
	return FindOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOneLogic) FindOne(req types.ReqId) (*types.RespUser, error) {
	// todo: add your logic here and delete this line

	return &types.RespUser{}, nil
}
