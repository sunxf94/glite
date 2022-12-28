package errors

var (
	ErrNone         = New(0, "success")
	ErrInvalidParam = New(1, "invalid param")
	ErrServer       = New(2, "internal server error")
)

func New(code int, msg string) *Err {
	return &Err{Code: code, Msg: msg}
}

type Err struct {
	Code int
	Msg  string
}

func (e *Err) HasErr() bool {
	return ErrNone.Code == e.Code
}

func (e *Err) Message() string {
	return e.Msg
}
