package udp

import (
	"github.com/localzet/aura/common/buf"
	"github.com/localzet/aura/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
