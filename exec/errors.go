package exec

import "errors"

var (
	errNoVersion          = errors.New("bytecode contains no version")
	errNoFrames           = errors.New("no frames to execute")
	errUnexpectedEOP      = errors.New("unexpected end of program")
	errOpRet              = errors.New("operation return")
	errEmptyDelegatorName = errors.New("empty delegator name")
	errUnexpectedEOF      = errors.New("unexpected end of frames")
	errNoOperation        = errors.New("no operation in frame")
)
