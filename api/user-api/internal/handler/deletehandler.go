package handler

import (
	"net/http"

	"github.com/blog/v1/api/user-api/internal/logic"
	"github.com/blog/v1/api/user-api/internal/svc"
	"github.com/blog/v1/api/user-api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func DeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteLogic(r.Context(), ctx)
		resp, err := l.Delete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
