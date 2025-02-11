package outbound

import (
	"fmt"

	"github.com/iscodefather/uri-to-json/pkgs/outbound/xray"
	"github.com/iscodefather/uri-to-json/pkgs/parser"
	"github.com/iscodefather/uri-to-json/pkgs/utils"
)

type ClientType string

const (
	XrayCore ClientType = "xray"
)

func GetOutbound(clientType ClientType, rawUri string) (result IOutbound) {
	scheme := utils.ParseScheme(rawUri)
	switch clientType {
	case XrayCore:
		switch scheme {
		case parser.SchemeVmess:
			result = &xray.VmessOut{RawUri: rawUri}
		case parser.SchemeVless:
			result = &xray.VlessOut{RawUri: rawUri}
		case parser.SchemeTrojan:
			result = &xray.TrojanOut{RawUri: rawUri}
		case parser.SchemeSS:
			result = &xray.ShadowSocksOut{RawUri: rawUri}
		case parser.SchemeSocks:
			result = &xray.SocksOut{RawUri: rawUri}
		case parser.SchemeHttp:
			result = &xray.HTTPOut{RawUri: rawUri}
		default:
			fmt.Println("unsupported protocol: ", scheme)
		}
	default:
		fmt.Println("unsupported client type")
	}
	return
}
