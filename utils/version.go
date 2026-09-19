package utils

import (
	"fmt"
	"runtime/debug"
	"strings"
)

var Version = "dev"

func GetVersion() string {
	if Version != "dev" && Version != "" {
		if !strings.HasPrefix(Version, "v") {
			return "v" + Version
		}
		return Version
	}

	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		v := info.Main.Version
		if !strings.HasPrefix(v, "v") {
			return "v" + v
		}
		return v
	}

	return "dev"
}

func PrintVersion() {
	fmt.Printf("fileWatcher %s\n", GetVersion())
}
