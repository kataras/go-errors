// Package errors helps you to write and design your own pre-defined errors, useful when you have a known list of errors
package errors

import (
	"fmt"
	"runtime"
)

var (
	// Prefix the error prefix, applies to each error's message
	// defaults to Error:_
	Prefix = "Error: "
	// NewLine adds a new line to the end of each error's message
	// defaults to true
	NewLine = true
)

// Error holds the error message, this message never really changes
type Error struct {
	message string
}

// New creates and returns an Error with a pre-defined user output message
// all methods below that doesn't accept a pointer receiver because actualy they are not changing the original message
func New(errMsg string) *Error {
	if NewLine {
		errMsg += "\n"
	}
	return &Error{message: Prefix + errMsg}
}

// Error returns the message of the actual error
// implements the error
func (e Error) Error() string {
	return e.message
}

// Format returns a formatted new error based on the arguments
// it does NOT change the original error's message
func (e Error) Format(a ...interface{}) error {
	e.message = fmt.Sprintf(e.message, a...)
	return e
}

// Append adds a message to the predefined error message and returns a new error
// it does NOT change the original error's message
func (e Error) Append(format string, a ...interface{}) Error {
	// eCp := *e
	if NewLine {
		format += "\n"
	}
	e.message += fmt.Sprintf(format, a...)
	return e
}

// AppendErr adds an error's message to the predefined error message and returns a new error
// it does NOT change the original error's message
func (e Error) AppendErr(err error) Error {
	return e.Append(err.Error())
}

// Panic output the message and after panics
func (e Error) Panic() {
	_, fn, line, _ := runtime.Caller(1)
	errMsg := e.message
	errMsg = "\nCaller was: " + fmt.Sprintf("%s:%d", fn, line)
	panic(errMsg)
}

// Panicf output the formatted message and after panics
func (e Error) Panicf(args ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	errMsg := e.Format(args...).Error()
	errMsg = "\nCaller was: " + fmt.Sprintf("%s:%d", fn, line)
	panic(errMsg)
}
