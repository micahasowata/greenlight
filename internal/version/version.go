package version

import (
	"runtime/debug"
)

func New() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}

	return bi.Main.Version
}
