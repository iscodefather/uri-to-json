package xray

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iscodefather/uri-to-json/pkgs/parser"
	"github.com/iscodefather/uri-to-json/pkgs/utils"
)

/*
https://xtls.github.io/config/outbounds/wireguard.html#outboundconfigurationobject

{
  "secretKey": "PRIVATE_KEY",
  "address": [
    // optional, default ["10.0.0.1", "fd59:7153:2388:b5fd:0000:0000:0000:0001"]
    "IPv4_CIDR",
    "IPv6_CIDR",
    "and more..."
  ],
  "peers": [
    {
      "endpoint": "ENDPOINT_ADDR",
      "publicKey": "PUBLIC_KEY"
    }
  ],
  "noKernelTun": false,
  "mtu": 1420, // optional, default 1420
  "reserved": [1, 2, 3],
  "workers": 2, // optional, default runtime.NumCPU()
  "domainStrategy": "ForceIP"
}
*/

var XrayWireguard string = `{
	"secretKey": "PRIVATE_KEY",
	"address": [
		"IPv4_CIDR",
		"IPv6_CIDR"
	],
	"peers": [
		{
			"endpoint": "ENDPOINT_ADDR",
			"publicKey": "PUBLIC_KEY"
		}
	],
	"noKernelTun": false,
	"mtu": 1420,
	"reserved": [1, 2, 3],
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

func (that *WireguardOut) GetRawUri() string {
	return that.RawUri
}

func (that *WireguardOut) getSettings() string {
	j := gjson.New(XrayWireguard)
	j.Set("secretKey", that.Parser.SecretKey)
	j.Set("address.0", that.Parser.Address)
	j.Set("peers.0.endpoint", that.Parser.Endpoint)
	j.Set("peers.0.publicKey", that.Parser.PublicKey)
	return j.MustToJsonString()
}

func (that *WireguardOut) setProtocolAndTag(outStr string) string {
	j := gjson.New(outStr)
	j.Set("protocol", "wireguard")
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonString()
}

func (that *WireguardOut) GetOutboundStr() string {
	if that.Parser.Address == "" && that.Parser.Port == 0 {
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

func TestWireguardOut() {
	rawUri := `wireguard://{"PrivateKey":"2B8LLjlXkJ608ct0LD0UnuuR9A2GuZUFBMBQJ9GFn1I=","AddrV4":"172.16.0.2","AddrV6":"2606:4700:110:8dad:87b4:b141:584d:e9dc","DNS":"1.1.1.1","MTU":1280,"PublicKey":"bmXOC+F1FxEMF9dyiK2H5/1SUtzH0JuVo51h2wPfgyo=","AllowedIPs":["0.0.0.0/0","::/0"],"Endpoint":"198.41.222.233:2087","ClientID":"GpxH","DeviceName":"D9D669","Reserved":null,"Address":"198.41.222.233","Port":2087}`
	to := &WireguardOut{}
	to.Parse(rawUri)
	o := to.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
	fmt.Println(o)
}
