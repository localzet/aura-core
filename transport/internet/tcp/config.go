package tcp

import (
	"github.com/localzet/aura/common"
	"github.com/localzet/aura/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
