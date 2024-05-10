package customError

import "strconv"

type Http struct {
	code int
	msg  string
}

func New(code int, msg string) *Http {
	if code < 400 || code >= 600 {
		code = 500
	}

	return &Http{
		code: code,
		msg:  msg,
	}
}

func (e *Http) Error() string {
	return strconv.Itoa(e.code) + e.msg
}

func (e *Http) GetCode() int {
	return e.code
}

func (e *Http) GetMsg() string {
	return e.msg
}
