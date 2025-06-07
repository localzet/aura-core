package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/localzet/aura/app/dispatcher"
	_ "github.com/localzet/aura/app/proxyman/inbound"
	_ "github.com/localzet/aura/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/localzet/aura/app/commander"
	_ "github.com/localzet/aura/app/log/command"
	_ "github.com/localzet/aura/app/proxyman/command"
	_ "github.com/localzet/aura/app/stats/command"

	// Developer preview services
	_ "github.com/localzet/aura/app/observatory/command"

	// Other optional features.
	_ "github.com/localzet/aura/app/dns"
	_ "github.com/localzet/aura/app/dns/fakedns"
	_ "github.com/localzet/aura/app/log"
	_ "github.com/localzet/aura/app/metrics"
	_ "github.com/localzet/aura/app/policy"
	_ "github.com/localzet/aura/app/reverse"
	_ "github.com/localzet/aura/app/router"
	_ "github.com/localzet/aura/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/localzet/aura/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/localzet/aura/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/localzet/aura/proxy/blackhole"
	_ "github.com/localzet/aura/proxy/dns"
	_ "github.com/localzet/aura/proxy/dokodemo"
	_ "github.com/localzet/aura/proxy/freedom"
	_ "github.com/localzet/aura/proxy/http"
	_ "github.com/localzet/aura/proxy/loopback"
	_ "github.com/localzet/aura/proxy/shadowsocks"
	_ "github.com/localzet/aura/proxy/socks"
	_ "github.com/localzet/aura/proxy/trojan"
	_ "github.com/localzet/aura/proxy/vless/inbound"
	_ "github.com/localzet/aura/proxy/vless/outbound"
	_ "github.com/localzet/aura/proxy/vmess/inbound"
	_ "github.com/localzet/aura/proxy/vmess/outbound"
	_ "github.com/localzet/aura/proxy/wireguard"

	// Transports
	_ "github.com/localzet/aura/transport/internet/grpc"
	_ "github.com/localzet/aura/transport/internet/httpupgrade"
	_ "github.com/localzet/aura/transport/internet/kcp"
	_ "github.com/localzet/aura/transport/internet/reality"
	_ "github.com/localzet/aura/transport/internet/splithttp"
	_ "github.com/localzet/aura/transport/internet/tcp"
	_ "github.com/localzet/aura/transport/internet/tls"
	_ "github.com/localzet/aura/transport/internet/udp"
	_ "github.com/localzet/aura/transport/internet/websocket"

	// Transport headers
	_ "github.com/localzet/aura/transport/internet/headers/http"
	_ "github.com/localzet/aura/transport/internet/headers/noop"
	_ "github.com/localzet/aura/transport/internet/headers/srtp"
	_ "github.com/localzet/aura/transport/internet/headers/tls"
	_ "github.com/localzet/aura/transport/internet/headers/utp"
	_ "github.com/localzet/aura/transport/internet/headers/wechat"
	_ "github.com/localzet/aura/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/localzet/aura/main/json"
	_ "github.com/localzet/aura/main/toml"
	_ "github.com/localzet/aura/main/yaml"

	// Load config from file or http(s)
	_ "github.com/localzet/aura/main/confloader/external"

	// Commands
	_ "github.com/localzet/aura/main/commands/all"
)
