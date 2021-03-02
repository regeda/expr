package memory

import "errors"

var (
	errOutOfBounds   = errors.New("memory: out of bounds")
	errGridOverflow  = errors.New("memory: grid limit exceeded")
	errLinksOverflow = errors.New("memory: links limit exceeded")
	errOutOfGrid     = errors.New("memory: out of grid")
)
