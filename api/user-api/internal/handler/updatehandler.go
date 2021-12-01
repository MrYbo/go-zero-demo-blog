package handler

import (
	"net/http"

	"github.com/blog/v1/api/user-api/internal/logic"
	"github.com/blog/v1/api/user-api/internal/svc"
	"github.com/blog/v1/api/user-api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func updateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUser
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateLogic(r.Context(), ctx)
		resp, err := l.Update(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
