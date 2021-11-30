package logic

import (
	"context"
	"github.com/blog/v1/api/user-api/internal/svc"
	"github.com/blog/v1/api/user-api/internal/types"
	user "github.com/blog/v1/rpc/user-rpc/user"

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

func (l *FindAllLogic) FindAll(req types.SelectParameters) ([]types.BaseUser, error) {
	resp, _ := l.svcCtx.Users.FindAll(l.ctx, &user.SelectParameters{Page: int64(req.Page), PageSize: int64(req.PageSize)})

	var allUser = make([]types.BaseUser, len(resp.Users))
	for i := 0; i < len(resp.Users); i++ {
		item := resp.Users[i]
		allUser[i] = types.BaseUser{
			Id:        int(item.Id),
			Username:  item.Username,
			Avatar:    item.Avatar,
			Name:      item.Name,
			Phone:     item.Phone,
			Address:   item.Address,
			Birthday:  item.Birthday,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
	}
	return allUser, nil
}
