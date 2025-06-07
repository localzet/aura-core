package singbridge

import (
	"context"
	"os"

	"github.com/localzet/aura/common/net"
	"github.com/localzet/aura/common/net/cnc"
	"github.com/localzet/aura/common/session"
	"github.com/localzet/aura/proxy"
	"github.com/localzet/aura/transport"
	"github.com/localzet/aura/transport/internet"
	"github.com/localzet/aura/transport/pipe"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
)

var _ N.Dialer = (*AuraDialer)(nil)

type AuraDialer struct {
	internet.Dialer
}

func NewDialer(dialer internet.Dialer) *AuraDialer {
	return &AuraDialer{dialer}
}

func (d *AuraDialer) DialContext(ctx context.Context, network string, destination M.Socksaddr) (net.Conn, error) {
	return d.Dialer.Dial(ctx, ToDestination(destination, ToNetwork(network)))
}

func (d *AuraDialer) ListenPacket(ctx context.Context, destination M.Socksaddr) (net.PacketConn, error) {
	return nil, os.ErrInvalid
}

type AuraOutboundDialer struct {
	outbound proxy.Outbound
	dialer   internet.Dialer
}

func NewOutboundDialer(outbound proxy.Outbound, dialer internet.Dialer) *AuraOutboundDialer {
	return &AuraOutboundDialer{outbound, dialer}
}

func (d *AuraOutboundDialer) DialContext(ctx context.Context, network string, destination M.Socksaddr) (net.Conn, error) {
	outbounds := session.OutboundsFromContext(ctx)
	if len(outbounds) == 0 {
		outbounds = []*session.Outbound{{}}
		ctx = session.ContextWithOutbounds(ctx, outbounds)
	}
	ob := outbounds[len(outbounds)-1]
	ob.Target = ToDestination(destination, ToNetwork(network))

	opts := []pipe.Option{pipe.WithSizeLimit(64 * 1024)}
	uplinkReader, uplinkWriter := pipe.New(opts...)
	downlinkReader, downlinkWriter := pipe.New(opts...)
	conn := cnc.NewConnection(cnc.ConnectionInputMulti(downlinkWriter), cnc.ConnectionOutputMulti(uplinkReader))
	go d.outbound.Process(ctx, &transport.Link{Reader: downlinkReader, Writer: uplinkWriter}, d.dialer)
	return conn, nil
}

func (d *AuraOutboundDialer) ListenPacket(ctx context.Context, destination M.Socksaddr) (net.PacketConn, error) {
	return nil, os.ErrInvalid
}
