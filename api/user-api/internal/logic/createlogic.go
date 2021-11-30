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

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateLogic {
	return CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req types.UserParams) (*types.BaseUser, error) {
	respUser, err := l.svcCtx.Users.Create(l.ctx, &user.UserParams{
		Username: req.Username,
		Password: req.Password,
		Avatar:   req.Avatar,
		Name:     req.Name,
		Phone:    req.Phone,
		Address:  req.Address,
		Birthday: req.Birthday,
	})

	if err != nil {
		return nil, errorx.NewCodeError(http.StatusConflict, "用户名已存在")
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
	}, nil
}
