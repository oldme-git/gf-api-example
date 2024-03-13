package packed

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"strings"
)

type pErr struct {
	maps map[int]string
}

var Err = &pErr{
	maps: map[int]string{
		0:     "请求成功",
		10001: "用户名或密码错误",
		10002: "用户不存在",
		99999: "未知错误",
	},
}

// GetMsg 获取code码对应的msg
func (c *pErr) GetMsg(code int) string {
	return c.maps[code]
}

// Skip 抛出一个业务级别的错误，不会打印错误堆栈信息
func (c *pErr) Skip(code int, msg ...string) (err error) {
	var msgStr string
	if len(msg) == 0 {
		msgStr = c.GetMsg(code)
	} else {
		msg = append([]string{c.GetMsg(code)}, msg...)
		msgStr = strings.Join(msg, ", ")
	}
	return gerror.NewWithOption(gerror.Option{
		Stack: false,
		Text:  msgStr,
		Code:  gcode.New(code, "", nil),
	})
}

// Sys 抛出一个系统级别的错误，使用code码：99999，会打印错误堆栈信息
// msg 接受string和error类型
// !!! 使用该方法传入error类型时，一定要注意不要泄露系统信息
func (c *pErr) Sys(msg ...interface{}) error {
	var (
		code     = 99999
		msgSlice = []string{
			c.GetMsg(code),
		}
	)

	if len(msg) != 0 {
		for _, v := range msg {
			switch a := v.(type) {
			case error:
				msgSlice = append(msgSlice, a.Error())
			case string:
				msgSlice = append(msgSlice, a)
			}
		}
	}

	msgStr := strings.Join(msgSlice, ", ")
	return gerror.NewCode(gcode.New(code, "", nil), msgStr)
}
