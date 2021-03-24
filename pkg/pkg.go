package pkg

import "runtime/debug"

func Stack() string {
	return string(debug.Stack())
}
