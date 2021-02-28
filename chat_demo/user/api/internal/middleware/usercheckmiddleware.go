package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

var (
	errorUserInfo = errors.New("用户信息获取失败")
	authDeny      = errors.New("用户信息不一致")
)

const (
	userKey = `x-user-id`
)

type UsercheckMiddleware struct {
}

func NewUsercheckMiddleware() *UsercheckMiddleware {
	return &UsercheckMiddleware{}
}

func (m *UsercheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get(userKey)
		jwtUserId := r.Context().Value("userId")
		userInt, err := json.Number(userId).Int64()
		if err != nil {
			httpx.Error(w, errorUserInfo)
			return
		}

		jwtInt, err := json.Number(fmt.Sprintf("%v", jwtUserId)).Int64()
		if err != nil {
			httpx.Error(w, errorUserInfo)
			return
		}

		if jwtInt != userInt {
			httpx.Error(w, authDeny)
			return
		}

		next(w, r)
	}
}
