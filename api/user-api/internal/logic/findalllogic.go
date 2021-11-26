package logic

import (
	"context"

	"github.com/blog/v1/api/user-api/internal/svc"
	"github.com/blog/v1/api/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindAllLogic {
	return FindAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAllLogic) FindAll() ([]types.RespUser, error) {
	// todo: add your logic here and delete this line

	return []types.RespUser{}, nil
}
