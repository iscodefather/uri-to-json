package xray

import (
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iscodefather/uri-to-json/pkgs/parser"
	"github.com/iscodefather/uri-to-json/pkgs/utils"
)

/*
https://xtls.github.io/config/outbounds/socks.html#outboundconfigurationobject

{
 "servers": [
  {
   "address": "127.0.0.1",
   "port": 1234,
   "users": [
    {
     "user": "test user",
     "pass": "test pass",
     "level": 0
    }
   ]
  }
 ]
}
*/

var XraySocks string = `{
	"servers": [
		{
			"address": "127.0.0.1",
			"port": 1234,
			"users": [
				{
					"user": "test user",
					"pass": "test pass"
				}
			]
		}
	]
}`

type SocksOut struct {
	RawUri   string
	Parser   *parser.ParserSocks
	outbound string
}

func (that *SocksOut) Parse(rawUri string) {
	that.Parser = &parser.ParserSocks{}
	that.Parser.Parse(rawUri)
}

func (that *SocksOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *SocksOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *SocksOut) Scheme() string {
	return parser.SchemeSocks
}

func (that *SocksOut) GetRawUri() string {
	return that.RawUri
}

func (that *SocksOut) getSettings() string {
	j := gjson.New(XraySocks)
	j.Set("servers.0.address", that.Parser.Address)
	j.Set("servers.0.port", that.Parser.Port)
	j.Set("servers.0.users.0.user", that.Parser.User)
	j.Set("servers.0.users.0.pass", that.Parser.Pass)
	return j.MustToJsonString()
}

func (that *SocksOut) setProtocolAndTag(outStr string) string {
	j := gjson.New(outStr)
	j.Set("protocol", "socks")
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonString()
}

func (that *SocksOut) GetOutboundStr() string {
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

func TestSocks() {
	rawUri := "socks://user:pass@127.0.0.1:1080#TestSocks"
	to := &SocksOut{}
	to.Parse(rawUri)
	o := to.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
	fmt.Println(o)
}
