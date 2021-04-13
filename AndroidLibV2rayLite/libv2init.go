package libv2ray

import (
	// The following are necessary as they register handlers in their init functions.

	// Required features. Can't remove unless there is replacements.
	_ "github.com/XTLS/Xray-core/app/dispatcher"
	_ "github.com/XTLS/Xray-core/app/proxyman/inbound"
	_ "github.com/XTLS/Xray-core/app/proxyman/outbound"

	// Other optional features.
	_ "github.com/XTLS/Xray-core/app/dns"
	_ "github.com/XTLS/Xray-core/app/log"
	_ "github.com/XTLS/Xray-core/app/policy"
	_ "github.com/XTLS/Xray-core/app/router"
	_ "github.com/XTLS/Xray-core/app/stats"

	// Inbound and outbound proxies.
	_ "github.com/XTLS/Xray-core/proxy/blackhole"
	_ "github.com/XTLS/Xray-core/proxy/dns"
	_ "github.com/XTLS/Xray-core/proxy/dokodemo"
	_ "github.com/XTLS/Xray-core/proxy/freedom"
	_ "github.com/XTLS/Xray-core/proxy/http"
	_ "github.com/XTLS/Xray-core/proxy/mtproto"
	_ "github.com/XTLS/Xray-core/proxy/shadowsocks"
	_ "github.com/XTLS/Xray-core/proxy/socks"
	_ "github.com/XTLS/Xray-core/proxy/trojan"
	_ "github.com/XTLS/Xray-core/proxy/vless/inbound"
	_ "github.com/XTLS/Xray-core/proxy/vless/outbound"
	_ "github.com/XTLS/Xray-core/proxy/vmess/inbound"
	_ "github.com/XTLS/Xray-core/proxy/vmess/outbound"

	// Transport
	_ "github.com/XTLS/Xray-core/transport/internet/http"
	_ "github.com/XTLS/Xray-core/internet/kcp"
	_ "github.com/XTLS/Xray-core/internet/quic"
	_ "github.com/XTLS/Xray-core/internet/tcp"
	_ "github.com/XTLS/Xray-core/internet/tls"
	_ "github.com/XTLS/Xray-core/internet/udp"
	_ "github.com/XTLS/Xray-core/internet/websocket"
	_ "github.com/XTLS/Xray-core/internet/grpc"

	// Transport headers
	_ "github.com/XTLS/Xray-core/transport/internet/headers/http"
	_ "github.com/XTLS/Xray-core/transport/internet/headers/noop"
	_ "github.com/XTLS/Xray-core/transport/internet/headers/srtp"
	_ "github.com/XTLS/Xray-core/transport/internet/headers/tls"
	_ "github.com/XTLS/Xray-core/transport/internet/headers/utp"
	_ "github.com/XTLS/Xray-core/transport/internet/headers/wechat"
	_ "github.com/XTLS/Xray-core/transport/internet/headers/wireguard"

	// JSON config support. Choose only one from the two below.
	// The following line loads JSON from v2ctl
	// _ "github.com/XTLS/Xray-core/main/json"
	// The following line loads JSON internally
	_ "github.com/XTLS/Xray-core/main/yaml"
	// Load config from file or http(s)
	// _ "github.com/XTLS/Xray-core/main/confloader/external"
)
