package gomodules

import (
	"fmt"
	"runtime/debug"

	errors "github.com/jybp/errors"
	gopkgin "gopkg.in/jybp/errors.v0"
)

// V4

func Stack() string {
	return string(debug.Stack())
}

func GopkginStack() string {
	return fmt.Sprintf("%+v", gopkgin.Stack())
}

func ErrorsStack() string {
	return fmt.Sprintf("%+v", errors.Stack())
}
