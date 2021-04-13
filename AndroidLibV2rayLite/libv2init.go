package libv2ray

import (
	// The following are necessary as they register handlers in their init functions.

	// Required features. Can't remove unless there is replacements.
	_ "v2ray.com/core/app/dispatcher"
	_ "v2ray.com/core/app/proxyman/inbound"
	_ "v2ray.com/core/app/proxyman/outbound"

	// Other optional features.
	_ "v2ray.com/core/app/dns"
	_ "v2ray.com/core/app/log"
	_ "v2ray.com/core/app/policy"
	_ "v2ray.com/core/app/router"
	_ "v2ray.com/core/app/stats"

	// Inbound and outbound proxies.
	_ "v2ray.com/core/proxy/blackhole"
	_ "v2ray.com/core/proxy/dns"
	_ "v2ray.com/core/proxy/dokodemo"
	_ "v2ray.com/core/proxy/freedom"
	_ "v2ray.com/core/proxy/http"
	_ "v2ray.com/core/proxy/mtproto"
	_ "v2ray.com/core/proxy/shadowsocks"
	_ "v2ray.com/core/proxy/socks"
	_ "v2ray.com/core/proxy/trojan"
	_ "v2ray.com/core/proxy/vless/inbound"
	_ "v2ray.com/core/proxy/vless/outbound"
	_ "v2ray.com/core/proxy/vmess/inbound"
	_ "v2ray.com/core/proxy/vmess/outbound"

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
	_ "v2ray.com/core/main/jsonem"
	// Load config from file or http(s)
	// _ "github.com/XTLS/Xray-core/main/confloader/external"
)
