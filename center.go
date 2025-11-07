package goerr

import (
	"github.com/puzpuzpuz/xsync"
)

var codeMap *xsync.MapOf[int, ErrCode]

var defaultErrCode = ErrCode{
	HttpCode: 200,
}

func init() {
	codeMap = xsync.NewIntegerMapOf[int, ErrCode]()
}

func register(code ErrCode) {
	if _, ok := codeMap.Load(code.BusinessCode); ok {
		return
	}
	codeMap.Store(code.BusinessCode, code)
}

func getCode(business int) ErrCode {
	code, ok := codeMap.Load(business)
	if ok {
		return code
	}
	ret := defaultErrCode
	ret.BusinessCode = business
	return ret
}
