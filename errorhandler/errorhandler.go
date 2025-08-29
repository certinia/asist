package errorhandler

import (
	"errors"
	"log"
	"os"
	"runtime"

	"github.com/certinia/asist/message"
)

type ExitCode int

const (
	ExitCodeSuccess       ExitCode = 0
	ExitCodeOccurrence    ExitCode = 1
	ExitCodeInternalError ExitCode = 3
	ExitCodeUserError     ExitCode = 4
)

/**
 * ExitWithCode - method used to exit the current running process with given exit code
 * It will also print stack trace in case of internal error
 */
func ExitWithCode(msg string, exitCode ExitCode) {
	log.Print(message.SetLogType("Error", msg))

	if exitCode == ExitCodeInternalError {
		// Log the stack trace
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, true)
		log.Printf("%s", buf)
	}

	os.Exit(int(exitCode))
}

func ExitWithError(err error) {
	var userErr *UserError
	var internalErr *InternalError
	if errors.As(err, &userErr) {
		ExitWithCode(err.Error(), ExitCodeUserError)
	}

	if errors.As(err, &internalErr) {
		ExitWithCode(err.Error(), ExitCodeInternalError)
	}
}

func NewUserError(msg string) *UserError {
	return &UserError{Message: msg}
}

func NewInternalError(msg string) *InternalError {
	return &InternalError{Message: msg}
}
