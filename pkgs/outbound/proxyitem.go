package outbound

import (
	"encoding/json"
	"fmt"

	"github.com/iscodefather/uri-to-json/pkgs/parser"
	"github.com/iscodefather/uri-to-json/pkgs/utils"
)

type ProxyItem struct {
	Scheme       string    `json:"scheme"`
	Address      string    `json:"address"`
	Port         int       `json:"port"`
	RTT          int64     `json:"rtt"`
	RawUri       string    `json:"raw_uri"`
	Location     string    `json:"location"`
	Outbound     string    `json:"outbound"`
	OutboundType ClientType `json:"outbound_type"`
}

func NewItem(rawUri string) *ProxyItem {
	return &ProxyItem{RawUri: rawUri}
}

func NewItemByEncryptedRawUri(enRawUri string) (item *ProxyItem) {
	rawUri := parser.ParseRawUri(enRawUri)
	if rawUri == "" {
		return
	}
	return &ProxyItem{RawUri: rawUri}
}

func (that *ProxyItem) parse() bool {
	that.Scheme = utils.ParseScheme(that.RawUri)
	that.OutboundType = XrayCore
	ob := GetOutbound(XrayCore, that.RawUri)

	if ob == nil {
		return false
	}
	ob.Parse(that.RawUri)
	that.Outbound = ob.GetOutboundStr()
	that.Address = ob.Addr()
	that.Port = ob.Port()
	return true
}

// Item string for conf.txt
func (that *ProxyItem) String() string {
	if that.Outbound == "" {
		if ok := that.parse(); !ok {
			return ""
		}
	}
	if r, err := json.Marshal(that); err == nil {
		return string(r)
	}
	return ""
}

func (that *ProxyItem) GetHost() string {
	if that.Address == "" && that.Port == 0 {
		return ""
	}
	return fmt.Sprintf("%s:%d", that.Address, that.Port)
}

func (that *ProxyItem) GetOutbound() string {
	if that.Outbound == "" {
		that.parse()
	}
	return that.Outbound
}

func (that *ProxyItem) GetOutboundType() ClientType {
	return that.OutboundType
}

// Automatically parse rawUri to ProxyItem for certain Client[xray-core]
func ParseRawUriToProxyItem(rawUri string, clientType ...ClientType) (p *ProxyItem) {
	p = NewItem(rawUri)
	p.Scheme = utils.ParseScheme(p.RawUri)
	p.OutboundType = XrayCore
	ob := GetOutbound(XrayCore, p.RawUri)
	if ob == nil {
		return
	}
	ob.Parse(p.RawUri)
	p.Outbound = ob.GetOutboundStr()
	p.Address = ob.Addr()
	p.Port = ob.Port()
	return
}

func ParseEncryptedRawUriToProxyItem(rawUri string) (p *ProxyItem) {
	rawUri = parser.ParseRawUri(rawUri)
	return ParseRawUriToProxyItem(rawUri)
}

// Transfer ProxyItem to specified ClientType: xray-core only
func TransferProxyItem(oldProxyItem *ProxyItem) (newProxyItem *ProxyItem) {
	if oldProxyItem == nil {
		return
	}
	if oldProxyItem.OutboundType == XrayCore {
		return oldProxyItem
	}
	newProxyItem = ParseRawUriToProxyItem(oldProxyItem.RawUri)
	newProxyItem.Location = oldProxyItem.Location
	newProxyItem.RTT = oldProxyItem.RTT
	return
}
