package outbound

import (
	"os"
	"sync"

	"encoding/json"

	"github.com/iscodefather/uri-to-json/pkgs/parser"
	"github.com/iscodefather/uri-to-json/pkgs/utils"
	"github.com/gvcgo/goutils/pkgs/gutils"
)

type Result struct {
	Vmess        []*ProxyItem `json:"Vmess"`
	Vless        []*ProxyItem `json:"Vless"`
	Trojan       []*ProxyItem `json:"Trojan"`
	ShadowSocks  []*ProxyItem `json:"Shadowsocks"`
	Socks        []*ProxyItem `json:"Socks"`
	Http         []*ProxyItem `json:"Http"`
	UpdateAt       string       `json:"UpdateAt"`
	VmessTotal     int          `json:"VmessTotal"`
	VlessTotal     int          `json:"VlessTotal"`
	TrojanTotal    int          `json:"TrojanTotal"`
	SSTotal        int          `json:"SSTotal"`
	SocksTotal     int          `json:"SocksTotal"`
	HttpTotal      int          `json:"HttpTotal"`
	totalList    []*ProxyItem
	lock         *sync.Mutex
}

func NewResult() *Result {
	return &Result{
		lock: &sync.Mutex{},
	}
}

func (that *Result) Load(fPath string) {
	if ok, _ := gutils.PathIsExist(fPath); ok {
		if content, err := os.ReadFile(fPath); err == nil {
			that.lock.Lock()
			json.Unmarshal(content, that)
			that.lock.Unlock()
		}
	}
}

func (that *Result) Save(fPath string) {
	if content, err := json.Marshal(that); err == nil {
		that.lock.Lock()
		os.WriteFile(fPath, content, os.ModePerm)
		that.lock.Unlock()
	}
}

func (that *Result) AddItem(proxyItem *ProxyItem) {
	that.lock.Lock()
	if proxyItem == nil {
		return
	}
	switch utils.ParseScheme(proxyItem.RawUri) {
	case parser.SchemeVmess:
		that.Vmess = append(that.Vmess, proxyItem)
		that.VmessTotal++
	case parser.SchemeVless:
		that.Vless = append(that.Vless, proxyItem)
		that.VlessTotal++
	case parser.SchemeTrojan:
		that.Trojan = append(that.Trojan, proxyItem)
		that.TrojanTotal++
	case parser.SchemeSS:
		that.ShadowSocks = append(that.ShadowSocks, proxyItem)
		that.SSTotal++
	case parser.SchemeSocks:
		that.Socks = append(that.Socks, proxyItem)
		that.SocksTotal++
	case parser.SchemeHttp:
		that.Http = append(that.Http, proxyItem)
		that.HttpTotal++
	default:
	}
	that.totalList = append(that.totalList, proxyItem)
	that.lock.Unlock()
}

func (that *Result) Len() int {
	return that.VmessTotal + that.VlessTotal + that.TrojanTotal + that.SSTotal + that.SocksTotal + that.HttpTotal
}

func (that *Result) GetTotalList() []*ProxyItem {
	if len(that.totalList) != that.Len() {
		that.totalList = append(that.totalList, that.Vmess...)
		that.totalList = append(that.totalList, that.Vless...)
		that.totalList = append(that.totalList, that.Trojan...)
		that.totalList = append(that.totalList, that.ShadowSocks...)
		that.totalList = append(that.totalList, that.Socks...)
		that.totalList = append(that.totalList, that.Http...)
	}
	return that.totalList
}

func (that *Result) Clear() {
	that.lock.Lock()
	that.Vmess = []*ProxyItem{}
	that.VmessTotal = 0
	that.Vless = []*ProxyItem{}
	that.VlessTotal = 0
	that.Trojan = []*ProxyItem{}
	that.TrojanTotal = 0
	that.ShadowSocks = []*ProxyItem{}
	that.SSTotal = 0
	that.Socks = []*ProxyItem{}
	that.SocksTotal = 0
	that.Http = []*ProxyItem{}
	that.HttpTotal = 0
	that.totalList = []*ProxyItem{}
	that.lock.Unlock()
}
