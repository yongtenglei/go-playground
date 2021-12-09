package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("myHandler\n"))
	// fmt.Fprintln(w, "myHandler") // 本质依然是调用 w.Write
}

type WelcomeHandler struct{}

func (wh *WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WelcomeHandler") // 本质依然是调用 w.Write
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.Handle("/", &myHandler{})

	http.Handle("/welcome", &WelcomeHandler{})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
