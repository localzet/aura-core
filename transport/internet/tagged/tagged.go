package tagged

import (
	"context"

	"github.com/localzet/aura/common/net"
	"github.com/localzet/aura/features/routing"
)

type DialFunc func(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
