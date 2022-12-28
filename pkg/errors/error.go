package errors

type Error interface {
	HasErr() bool
	Message() string
}
