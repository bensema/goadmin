package ecode

import (
	"github.com/pkg/errors"
)

// New new a ecode.Codes by int value.
// NOTE: ecode must unique in global, the New will check repeat and then panic.
func New(e int, msg string) Code {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return add(e, msg)
}

func add(e int, msg string) Code {
	return Code{e, msg}
}

// Codes ecode error interface which has a code & message.
type Codes interface {
	// sometimes Error return Code in string form
	// NOTE: don't use Error in monitor report even it also work for now
	Error() string
	// Code get error code.
	Code() int
	// Message get code message.
	Message() string
	//Detail get error detail,it may be nil.
	Details() []interface{}
	// Equal for compatible.
	// Deprecated: please use ecode.EqualError.
	Equal(error) bool
}

// A Code is an int error code spec.
type Code struct {
	code int
	msg  string
}

func (e Code) Error() string {
	return e.msg
}

// Code return error code
func (e Code) Code() int { return e.code }

// Message return error message
func (e Code) Message() string {
	return e.Error()
}

// Details return details.
func (e Code) Details() []interface{} { return nil }

// Equal for compatible.
// Deprecated: please use ecode.EqualError.
func (e Code) Equal(err error) bool { return EqualError(e, err) }

// String parse code string to error.
func String(e string) Code {
	if e == "" {
		return OK
	}

	return Code{public.code, e}
}

// Cause cause from error to ecode.
func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return String(e.Error())
}

// Equal equal a and b by code int.
func Equal(a, b Codes) bool {
	if a == nil {
		a = OK
	}
	if b == nil {
		b = OK
	}
	return a.Code() == b.Code()
}

// EqualError equal error
func EqualError(code Codes, err error) bool {
	return Cause(err).Code() == code.Code()
}
