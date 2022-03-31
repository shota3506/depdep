package baz

import (
	"foo/bar"      // want "found dependency on foo/bar in foo/baz package"
	"foo/qux"      // want "found dependency on foo/qux in foo/baz package"
	"foo/qux/quux" // want "found dependency on foo/qux/quux in foo/baz package"
)

func Hello() string {
	return bar.Hello() + " " + qux.Hello() + " " + quux.Hello()
}
