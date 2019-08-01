# cmux2

Nested muxer.  
Extremely bare-bones.

## Usage Example

From [cmux2_example](./cmd/cmux2_example/main.go).

```go
mux := cmux2.Branch(cmux2.Routes{
	"": defaultHandler,
	"page": cmux2.Branch(cmux2.Routes{
		"": pageHandler,
		"foo": pageFooHandler,
		"bar": pageBarHandler,
	}),
})
http.ListenAndServe(":8080", mux)
```

Each node trims the Request path before passing it down, so the finishing http.HandlerFunc will see a relative path in the Request.