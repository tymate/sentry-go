// +build appengine

package sentry

import (
	"runtime"
	"strings"

	appengine "google.golang.org/appengine"
)

func init() {
	AddDefaultOptions(
		RuntimeContext("go", strings.TrimPrefix(runtime.Version(), "go")),
		OSContext(&OSContextInfo{
			Name:          "google-appengine",
			Version:       appengine.ServerSoftware,
		}),
		DeviceContext(&DeviceContextInfo{
			Model:        "Unknown",
		}),
	)
}
