package exception

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"myapp/internal/packed"
	"myapp/internal/service"
)

type sException struct {
}

func init() {
	service.RegisterException(New())
}

func New() *sException {
	return &sException{}
}

func (s *sException) Business() error {
	return packed.Err.Skip(10001)
}

// System 这里我们对 gjson.Decode() 传入错误数据，用来模拟组件内部抛出err
func (s *sException) System() error {
	_, err := gjson.Decode("")
	if err != nil {
		return packed.Err.Sys("可选自定义信息")
	}
	return nil
}
