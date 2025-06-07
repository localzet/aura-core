package conf_test

import (
	"testing"

	"github.com/localzet/aura/common/net"
	"github.com/localzet/aura/common/protocol"
	"github.com/localzet/aura/common/serial"
	. "github.com/localzet/aura/infra/conf"
	"github.com/localzet/aura/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() Buildable {
		return new(ShadowsocksServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"method": "aes-256-GCM",
				"password": "aura-password"
			}`,
			Parser: loadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				Users: []*protocol.User{{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_256_GCM,
						Password:   "aura-password",
					}),
				}},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}
