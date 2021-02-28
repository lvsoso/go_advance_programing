package handler

import (
	"net/http"

	"chat_demo/user/api/internal/logic"
	"chat_demo/user/api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func userInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("x-user-id")
		l := logic.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo(userId)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
