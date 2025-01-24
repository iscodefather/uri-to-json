package xray

import (
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iscodefather/uri-to-json/pkgs/parser"
	"github.com/iscodefather/uri-to-json/pkgs/utils"
)

var XrayWireguard string = `{
    "secretKey": "",
    "address": [],
    "peers": [
        {
            "endpoint": "",
            "publicKey": ""
        }
    ],
    "noKernelTun": false,
    "mtu": 1420,
    "reserved": [],
    "workers": 2,
    "domainStrategy": "ForceIP"
}`

type WireguardOut struct {
	RawUri   string
	Parser   *parser.ParserWireguard
	outbound string
}

func (that *WireguardOut) Parse(rawUri string) {
	that.Parser = &parser.ParserWireguard{}
	that.Parser.Parse(rawUri)
}

func (that *WireguardOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *WireguardOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *WireguardOut) Scheme() string {
	return "wireguard"
}

func (that *WireguardOut) GetRawUri() string {
	return that.RawUri
}

func (that *WireguardOut) getSettings() string {
	j := gjson.New(XrayWireguard)
	j.Set("secretKey", that.Parser.SecretKey)
	j.Set("address", that.Parser.Address)
	j.Set("noKernelTun", that.Parser.NoKernelTun)
	j.Set("mtu", that.Parser.MTU)
	j.Set("reserved", that.Parser.Reserved)
	j.Set("workers", that.Parser.Workers)
	j.Set("domainStrategy", that.Parser.DomainStrategy)

	if len(that.Parser.Peers) > 0 {
		peer := that.Parser.Peers[0]
		j.Set("peers.0.endpoint", peer.Endpoint)
		j.Set("peers.0.publicKey", peer.PublicKey)
		j.Set("peers.0.preSharedKey", peer.PreSharedKey)
		j.Set("peers.0.keepAlive", peer.KeepAlive)
		j.Set("peers.0.allowedIPs", peer.AllowedIPs)
	}

	return j.MustToJsonString()
}

func (that *WireguardOut) setProtocolAndTag(outStr string) string {
	j := gjson.New(outStr)
	j.Set("protocol", "wireguard")
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonString()
}

func (that *WireguardOut) GetOutboundStr() string {
	if that.Parser.SecretKey == "" || len(that.Parser.Peers) == 0 {
		return ""
	}
	if that.outbound == "" {
		settings := that.getSettings()
		stream := PrepareStreamString(that.Parser.StreamField)
		outStr := fmt.Sprintf(XrayOut, settings, stream)
		that.outbound = that.setProtocolAndTag(outStr)
	}
	return that.outbound
}

func TestWireguard() {
	rawUri := "wireguard://?secretKey=PRIVATE_KEY&address=10.0.0.1,fd59:7153:2388:b5fd:0000:0000:0000:0001&peers=ENDPOINT_ADDR,2408,PUBLIC_KEY,PRE_SHARED_KEY,0,0.0.0.0/0&noKernelTun=false&mtu=1420&reserved=1,2,3&workers=2&domainStrategy=ForceIP"
	wo := &WireguardOut{}
	wo.Parse(rawUri)
	o := wo.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
	fmt.Println(o)
}
