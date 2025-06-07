// Package core provides an entry point to use Aura core functionalities.
//
// Aura makes it possible to accept incoming network connections with certain
// protocol, process the data, and send them through another connection with
// the same or a difference protocol on demand.
//
// It may be configured to work with multiple protocols at the same time, and
// uses the internal router to tunnel through different inbound and outbound
// connections.
package core

import (
	"fmt"
	"runtime"

	"github.com/localzet/aura/common/serial"
)

var (
	Version_x byte = 25
	Version_y byte = 5
	Version_z byte = 16
)

var (
	build    = "Custom"
	codename = "Aura, Penetrates Everything."
	intro    = "A unified platform for anti-censorship."
)

// Version returns Aura's version as a string, in the form of "x.y.z" where x, y and z are numbers.
// ".z" part may be omitted in regular releases.
func Version() string {
	return fmt.Sprintf("%v.%v.%v", Version_x, Version_y, Version_z)
}

// VersionStatement returns a list of strings representing the full version info.
func VersionStatement() []string {
	return []string{
		serial.Concat("Aura ", Version(), " (", codename, ") ", build, " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")"),
		intro,
	}
}
