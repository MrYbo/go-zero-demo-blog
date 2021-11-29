package logic

import (
	"context"
	"github.com/blog/v1/api/user-api/internal/svc"
	"github.com/blog/v1/api/user-api/internal/types"
	"github.com/blog/v1/common/errorx"
	user "github.com/blog/v1/rpc/user-rpc/user"
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
)

type DestroyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDestroyLogic(ctx context.Context, svcCtx *svc.ServiceContext) DestroyLogic {
	return DestroyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DestroyLogic) Destroy(req types.UserId) (*types.CommonOK, error) {
	_, err := l.svcCtx.Users.Destroy(l.ctx, &user.UserId{Id: req.Id})
	if err != nil {
		return nil, errorx.NewCodeError(http.StatusNotFound, err.Error())
	}

	return &types.CommonOK{
		Code: http.StatusOK,
	}, nil
}
