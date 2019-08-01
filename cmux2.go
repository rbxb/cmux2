package cmux2

import (
	"net/http"
	"strings"
)

type Routes map[string]http.HandlerFunc

func Branch(ro Routes) http.HandlerFunc {
	return func(w http.ResponseWriter, req * http.Request) {
		p := strings.Trim(req.URL.Path, "/")
		v := strings.Split(p, "/")
		if h := ro[v[0]]; h != nil {
			req.URL.Path = p[len(v[0]):]
			h(w,req)
		} else {
			http.Error(w, "Not found.", 404)
		}
	}
}