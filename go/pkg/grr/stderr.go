package grr

import "google.golang.org/grpc/codes"

type StdErr struct {
	err  error
	code codes.Code
}

func NewStdErr(err error, code codes.Code) *StdErr {
	return &StdErr{
		err:  err,
		code: code,
	}
}

func (e *StdErr) Error() string {
	return e.err.Error()
}

func (e *StdErr) Unwrap() error {
	return e.err
}

func (e *StdErr) Code() codes.Code {
	return e.code
}
