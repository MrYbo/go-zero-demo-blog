package logic

import (
	"context"
	"github.com/blog/v1/common/errorx"
	user "github.com/blog/v1/rpc/user-rpc/user"
	"github.com/golang-jwt/jwt"
	"net/http"

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

func (l *LoginLogic) getJwtToken(iat, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = l.svcCtx.Config.Auth.AccessExpire
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
}

func (l *LoginLogic) Login(req types.ReqLogin) (*types.RespLogin, error) {
	resp, err := l.svcCtx.Users.Login(l.ctx, &user.ReqLogin{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorx.NewCodeError(http.StatusUnauthorized, err.Error())
	}
	token, _ := l.getJwtToken(12, resp.Id)
	return &types.RespLogin{Token: token}, nil
}
