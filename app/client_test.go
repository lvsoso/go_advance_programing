package main

import (
	"encoding/json"
	"testing"
)

func TestPost(t *testing.T) {
	register := RegisterReq{
		Username: "ccc",
		Nickname: "bbb",
		Gender:   "male",
		Mobile:   "13190090119",
		Password: "222222",
	}
	data, err := json.Marshal(register)
	if err != nil {
	    t.Fatal(err)
	}
	url := serverURL + "/user/register"
	resp, err := Post(url, data)
	if err != nil && err.Error()!= "EOF"{
		t.Fatal(err)
	}
	t.Logf("%v", resp)
}