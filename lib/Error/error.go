package Error

import (
	"fmt"
	"order_flume/lib/beelog"
	"os"

	"github.com/go-ini/ini"
)

/*
 *DefineError错误结构体，注意不要改变结构体中字段顺序
 */
type DefineError struct {
	Code int32
	Msg  string
	Desc string
}

func (e DefineError) Error() string {
	return e.Msg
}

var cfg *ini.File

/*
 *Title:LoadErrorFile
 *Fun: 加载错误码文件
 *Param:
 *Ret:
 *Att:
 */
func LoadErrorFile() {
	var err error
	cfg, err = ini.Load("./Error/error.ini")

	if err != nil {
		println("can not find error.ini")
		os.Exit(1)
	}
}

/*
 *Name:GetDefError
 *Func:获取自定义的错误
 *Param:code 是错误码，args中第一个参数是错误描述,之所以用...实现是因为go不支持函数重载，为了减少函数的数量，才用...，通过判断args长度来确定描述是否为空
 *Ret:
 *Attention:
 */
func GetDefError(code int32, args ...string) DefineError {
	if nil == cfg {
		LoadErrorFile()
	}
	selection, err := cfg.GetSection("code")
	if err != nil {
		println("get selection error")
		beelog.Error(fmt.Sprintf("can not find error code:%v in the error file", code))
		return DefineError{-1, "undefine code", fmt.Sprintf("can not find error code:%v in the error file", code)}
	}
	str := fmt.Sprintf("0X%08X", code)
	key, errKey := selection.GetKey(str)
	if errKey != nil {
		println("can not find key" + str)
		beelog.Error(fmt.Sprintf("can not find error code:%v in the error file", code))
		return DefineError{-1, "undefine code", fmt.Sprintf("can not find error code:%v in the error file", code)}
	}
	count := 0
	value := ""
	for _, msg := range args {
		count++
		if count == 1 {
			value = msg
		}
	}

	msg := key.String()

	return DefineError{code, msg, value}
}
