package greet

import (
	"strings"
)

func Say(names []string) string {

	return "Hello, " + strings.Join(names, ", ") + "!"
}
