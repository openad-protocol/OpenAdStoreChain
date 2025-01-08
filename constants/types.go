// 公共错误库
package constants

import (
	"AdServerCollector/core/errors"
)

// 0-199	预留
// 200-299  参数错误
// 300-399	类型错误
// 400-499  中间件问题
// 500-599  内部错误
// 600-699  操作系统错误
// 700-699  协议错误

var ()

var (
	ErrRuntimePanic = errors.New("运行时错误", 500)
)
