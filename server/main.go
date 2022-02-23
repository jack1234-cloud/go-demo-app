package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/jack":
		fmt.Fprint(w, "Hello Jack, how are you today?")
	default:
		fmt.Fprint(w, "Big Error!")
	}
	fmt.Printf("Handling function with %s request\n", r.Method)
}

func htlmVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello World<h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout Attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did *not* timeout!")
}

func helloWorldJackMode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloWorldJackMode")
	fmt.Fprint(w, "<h1 style=\"background-color:blue;\">helloWorldJackMode<h1>")
}

func main() {
	http.HandleFunc("/", htlmVsPlain)
	http.HandleFunc("/jack", helloWorldPage)
	// http.ListenAndServe("", nil)

	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	var muxJackMode http.ServeMux
	server.Handler = &muxJackMode
	muxJackMode.HandleFunc("/", helloWorldJackMode)

	server.ListenAndServe()
}
