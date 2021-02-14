package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

func helloHandler(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

func showInfoHandler(wr http.ResponseWriter, r *http.Request) {
	// ...
}

func showEmailHandler(wr http.ResponseWriter, r *http.Request) {
	// ...
}

func showFriendsHandler(wr http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	wr.Write([]byte("your friends is tom and alex"))
	timeElapsed := time.Since(timeStart)
	logger.Println(timeElapsed)
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		logger.Println(timeElapsed)
	})
}

func main() {
	http.HandleFunc("/", timeMiddleware(http.HandleFunc(helloHandler)))
	http.HandleFunc("/info/show", showInfoHandler)
	http.HandleFunc("/email/show", showEmailHandler)
	http.HandleFunc("/friends/show", showFriendsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
