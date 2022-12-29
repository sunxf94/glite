package errors

var (
	ErrNone         = New(0, "success")
	ErrInvalidParam = New(1, "invalid param")
	ErrServer       = New(2, "internal server error")
)

func New(code int, msg string) *Err {
	return &Err{code: code, msg: msg}
}

type Err struct {
	code int
	msg  string
}

func (e *Err) HasErr() bool {
	return e.code != ErrNone.code
}

func (e *Err) Message() string {
	return e.msg
}

func (e *Err) Code() int {
	return e.code
}
