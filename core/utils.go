package core

import (
	"runtime/debug"
)

func RecoverAndPrint() {
	if s := recover(); s != nil {
		debug.PrintStack()
	}
}
