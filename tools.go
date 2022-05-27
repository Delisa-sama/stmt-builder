//go:build tools

// How can I track tool dependencies for a module?
// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package query

import (
	_ "golang.org/x/lint/golint"
)
