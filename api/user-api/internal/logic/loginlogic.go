package logic

import (
	"context"
	"github.com/blog/v1/common/errorx"
	user "github.com/blog/v1/rpc/user-rpc/user"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"

	"github.com/blog/v1/api/user-api/internal/svc"
	"github.com/blog/v1/api/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) IssuesToken(userId int64) (string, error) {
	type Claims struct {
		UserId int64
		jwt.StandardClaims
	}

	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: l.svcCtx.Config.Auth.AccessExpire,
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
}

func (l *LoginLogic) Login(req types.ReqLogin) (*types.RespLogin, error) {
	resp, err := l.svcCtx.Users.Login(l.ctx, &user.ReqLogin{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorx.NewCodeError(http.StatusUnauthorized, err.Error())
	}
	token, _ := l.IssuesToken(resp.Id)
	return &types.RespLogin{Token: token}, nil
}
