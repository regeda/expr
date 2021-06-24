package exec

import "errors"

var (
	errNoVersion          = errors.New("bytecode contains no version")
	errNoFrames           = errors.New("no frames to execute")
	errEmptyDelegatorName = errors.New("empty delegator name")
	errUnexpectedEOF      = errors.New("unexpected end of frames")
	errNoOperation        = errors.New("no operation in frame")
	errDivByZero          = errors.New("division by zero")
)
