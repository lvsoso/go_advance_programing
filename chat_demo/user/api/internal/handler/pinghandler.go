package handler

import (
	"net/http"

	"chat_demo/user/api/internal/logic"
	"chat_demo/user/api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func pingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewPingLogic(r.Context(), ctx)
		err := l.Ping()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
