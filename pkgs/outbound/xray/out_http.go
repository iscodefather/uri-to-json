package xray

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iscodefather/uri-to-json/pkgs/parser"
	"github.com/iscodefather/uri-to-json/pkgs/utils"
)

/*
https://xtls.github.io/config/outbounds/http.html#outboundconfigurationobject

{
  "servers": [
    {
      "address": "192.168.108.1",
      "port": 3128,
      "users": [
        {
          "user": "my-username",
          "pass": "my-password"
        }
      ]
    }
  ],
  "headers": {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
    "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2"
  }
}
*/

var XrayHTTP string = `{
	"servers": [
		{
			"address": "127.0.0.1",
			"port": 8080,
			"users": [
				{
					"user": "test user",
					"pass": "test pass"
				}
			]
		}
	]
}`

type HTTPOut struct {
	RawUri   string
	Parser   *parser.ParserHTTP
	outbound string
}

func (that *HTTPOut) Parse(rawUri string) {
	that.Parser = &parser.ParserHTTP{}
	that.Parser.Parse(rawUri)
}

func (that *HTTPOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *HTTPOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *HTTPOut) Scheme() string {
	return "http"
}

func (that *HTTPOut) GetRawUri() string {
	return that.RawUri
}

func (that *HTTPOut) getSettings() string {
	j := gjson.New(XrayHTTP)
	j.Set("servers.0.address", that.Parser.Address)
	j.Set("servers.0.port", that.Parser.Port)
	j.Set("servers.0.users.0.user", that.Parser.User)
	j.Set("servers.0.users.0.pass", that.Parser.Pass)
	return j.MustToJsonString()
}

func (that *HTTPOut) setProtocolAndTag(outStr string) string {
	j := gjson.New(outStr)
	j.Set("protocol", "http")
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonString()
}

func (that *HTTPOut) GetOutboundStr() string {
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

func TestHTTP() {
	rawUri := "http://user:pass@127.0.0.1:8080#TestHTTP"
	to := &HTTPOut{}
	to.Parse(rawUri)
	o := to.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
	fmt.Println(o)
}
