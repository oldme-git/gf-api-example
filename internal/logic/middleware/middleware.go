package middleware

import "myapp/internal/service"

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}
