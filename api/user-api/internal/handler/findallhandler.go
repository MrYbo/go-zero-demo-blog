package handler

import (
	"net/http"

	"github.com/blog/v1/api/user-api/internal/logic"
	"github.com/blog/v1/api/user-api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func FindAllHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFindAllLogic(r.Context(), ctx)
		resp, err := l.FindAll()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
