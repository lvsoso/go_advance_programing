package main

import (
	"bytes"
	"net/http"
)

func Post(url string, data []byte) (*[]byte,error) {
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	recvData := []byte{}
	_, err = resp.Body.Read(recvData)
	if err != nil {
		return nil, err
	}
	return &recvData, nil
}
