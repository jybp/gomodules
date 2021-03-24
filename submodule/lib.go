package submodule

import (
	"fmt"
	"runtime/debug"

	"github.com/jybp/errors"
	gopkgin "gopkg.in/jybp/errors.v0"
)

// V3

func Stack() string {
	return string(debug.Stack())
}

func Panic() {
	panic("panic")
}

func GopkginStack() string {
	return fmt.Sprintf("%+v", gopkgin.Stack())
}

func ErrorsStack() string {
	return fmt.Sprintf("%+v", errors.Stack())
}
