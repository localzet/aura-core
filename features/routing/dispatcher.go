package routing

import (
	"context"

	"github.com/localzet/aura/common/net"
	"github.com/localzet/aura/features"
	"github.com/localzet/aura/transport"
)

// Dispatcher is a feature that dispatches inbound requests to outbound handlers based on rules.
// Dispatcher is required to be registered in a Aura instance to make Aura function properly.
//
// aura:api:stable
type Dispatcher interface {
	features.Feature

	// Dispatch returns a Ray for transporting data for the given request.
	Dispatch(ctx context.Context, dest net.Destination) (*transport.Link, error)
	DispatchLink(ctx context.Context, dest net.Destination, link *transport.Link) error
}

// DispatcherType returns the type of Dispatcher interface. Can be used to implement common.HasType.
//
// aura:api:stable
func DispatcherType() interface{} {
	return (*Dispatcher)(nil)
}
