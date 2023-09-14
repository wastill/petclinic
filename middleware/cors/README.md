# http 跨域中间件

## example
```go
opts = append(opts, http.Filter(cors.CORS(
		cors.AllowedHeaders([]string{"*"}),
		cors.AllowedOrigins([]string{"*"}),
		cors.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"}),
	)))
```