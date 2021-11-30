package logic

import (
	"context"
	"github.com/blog/v1/model"
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

func (l *FindAllLogic) FindAll(in *user.SelectParameters) (*user.RespFindAll, error) {
	var fUsers []model.Users
	offset := (in.Page - 1) * in.Page
	l.svcCtx.UserModel.Debug().Limit(int(in.PageSize)).Offset(int(offset)).Find(&fUsers)

	var resUsers = make([]*user.BaseUser, len(fUsers))

	for i := 0; i < len(fUsers); i++ {
		var item = fUsers[i]
		resUsers[i] = &user.BaseUser{
			Id:        int64(item.ID),
			Username:  item.Username,
			Avatar:    item.Avatar,
			Name:      item.Name,
			Phone:     item.Phone,
			Address:   item.Address,
			Birthday:  item.Birthday,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: item.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return &user.RespFindAll{Users: resUsers}, nil
}
