// +build !appengin

package sentry

import (
	"runtime"
	"strings"

	"github.com/matishsiao/goInfo"
)

func init() {
	gi := goInfo.GetInfo()

	AddDefaultOptions(
		RuntimeContext("go", strings.TrimPrefix(runtime.Version(), "go")),
		OSContext(&OSContextInfo{
			Name:          gi.GoOS,
			Version:       gi.OS,
			KernelVersion: gi.Core,
		}),
		DeviceContext(&DeviceContextInfo{
			Architecture: gi.Platform,
			Family:       gi.Kernel,
			Model:        "Unknown",
		}),
	)
}
