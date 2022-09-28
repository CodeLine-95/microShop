package errorx

const defaultCode = 1001

type CodeError struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int64, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
