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

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateLogic {
	return UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req types.UpdateUser) (*types.BaseUser, error) {
	respUser, err := l.svcCtx.Users.Update(l.ctx, &user.UpdateUser{
		Id: int64(req.Id),
		Avatar: req.Avatar,
		Name: req.Name,
		Phone: req.Phone,
		Address: req.Address,
	})

	if err != nil {
		return nil, errorx.NewCodeError(http.StatusNotFound, "not found")
	}

	return &types.BaseUser{
		Id:        int(respUser.Id),
		Username:  respUser.Username,
		Avatar:    respUser.Avatar,
		Phone:     respUser.Phone,
		Name:      respUser.Name,
		Address:   respUser.Address,
		Birthday:  respUser.Birthday,
		CreatedAt: respUser.CreatedAt,
		UpdatedAt: respUser.UpdatedAt,
	}, nil
}
