package main

type RegisterReq struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserReply struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	Mobile   string `json:"mobile"`
	JwtToken
}

type JwtToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	AccessExpire int64  `json:"accessExpire,omitempty"`
	RefreshAfter int64  `json:"refreshAfter,omitempty"`
}
