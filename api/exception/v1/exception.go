package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 模拟业务异常
type BusinessReq struct {
	g.Meta `path:"/business" method:"get"`
}

type BusinessRes struct {
	Name string
	Age  int
}

// 模拟系统异常
type SystemReq struct {
	g.Meta `path:"/system" method:"get"`
}

type SystemRes struct {
	Name string
	Age  int
}
