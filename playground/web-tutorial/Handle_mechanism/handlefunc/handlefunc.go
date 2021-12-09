package main

import (
	"fmt"
	"net/http"
)

func RootHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("myHandler\n"))
	// fmt.Fprintln(w, "myHandler") // 本质依然是调用 w.Write
}

func WelcomeHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WelcomeHandler") // 本质依然是调用 w.Write
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/", RootHandleFunc)

	http.HandleFunc("/welcome", WelcomeHandleFunc)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
