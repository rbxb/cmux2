package cmux2

import (
	"net/http"
	"strings"
)

type Routes map[string]http.HandlerFunc

func Branch(ro Routes) http.HandlerFunc {
	return func(w http.ResponseWriter, req * http.Request) {
		v := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
		if h := match(ro,v); h != nil {
			req.URL.Path = req.URL.Path[len(v[0]):]
			h(w,req)
		} else {
			http.Error(w, "Not found.", 404)
		}
	}
}

func match(ro Routes, v []string) http.HandlerFunc {
	if len(v) == 0 {
		if h := ro[""]; h != nil {
			return h
		}
	} else if h := ro[v[0]]; h != nil {
		return h
	}
	return nil
}