package goblox

import (
	"runtime"
	"runtime/debug"
)

const modulePath = "github.com/luckfire/goblox"

var Version = moduleVersion()

var UserAgent = "goblox/" + Version + " (+https://" + modulePath + "; " + runtime.Version() + ")"

func AppUserAgent(app string) string {
	if app == "" {
		return UserAgent
	}

	return app + " " + UserAgent
}

func moduleVersion() string {
	info, ok := debug.ReadBuildInfo()

	if !ok {
		return "dev"
	}

	for _, dep := range info.Deps {
		if dep.Path == modulePath {
			if dep.Replace != nil {
				return "dev"
			}
			return dep.Version
		}
	}

	return "dev"
}
