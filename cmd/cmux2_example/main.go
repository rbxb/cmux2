package main

import (
	"net/http"
	"github.com/rbxb/cmux2"
)

func main() {
	mux := cmux2.Branch(cmux2.Routes{
		"": defaultHandler,
		"page": cmux2.Branch(cmux2.Routes{
			"": pageHandler,
			"foo": pageFooHandler,
			"bar": pageBarHandler,
		}),
	})
	http.ListenAndServe(":8080", mux)
}

func defaultHandler(w http.ResponseWriter, req * http.Request) {
	w.Write([]byte("Hello!"))
}

func pageHandler(w http.ResponseWriter, req * http.Request) {
	w.Write([]byte("<a href='./foo'>foo</a><br><a href='./bar'>bar</a>"))
}

func pageFooHandler(w http.ResponseWriter, req * http.Request) {
	w.Write([]byte("foo"))
}

func pageBarHandler(w http.ResponseWriter, req * http.Request) {
	w.Write([]byte("bar"))
}