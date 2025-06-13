package play

import (
	"fmt"
	"log"
	"net/http"
)

type HttpHandlerDecorator func(http.HandlerFunc) http.HandlerFunc

func Handler(h http.HandlerFunc, decors ...HttpHandlerDecorator) http.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		h = d(h)
	}
	return h
}

func HttpServer() {
	http.HandleFunc("/hello", Handler(hello, WithServerHeader, WithBasicAuth))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func WithServerHeader(fn http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("--->WithServerHeader()")
		writer.Header().Set("Server", "HelloServer v0.0.1")
		fn(writer, request)
	}
}

func WithBasicAuth(fn http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("--->WithAuthCookie()")
		cookie := &http.Cookie{Name: "Auth", Value: "Pass", Path: "/"}
		http.SetCookie(writer, cookie)
		fn(writer, request)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recieved Request %s from %s\n", r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
}
