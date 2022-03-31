package bar

import (
	"foo" // want "found dependency on foo in foo/bar package"
)

func Hello() string {
	return foo.Hello()
}
