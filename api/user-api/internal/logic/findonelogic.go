package logic

import (
	"context"
	"github.com/blog/v1/common/errorx"
	user "github.com/blog/v1/rpc/user-rpc/user"
	"net/http"

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

func (l *FindOneLogic) FindOne(req types.UserId) (*types.BaseUser, error) {
	respUser, err := l.svcCtx.Users.FindOne(l.ctx, &user.UserId{Id: req.Id})

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
