package parser

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

type Peer struct {
	Endpoint     string   `json:"endpoint"`
	PublicKey    string   `json:"publicKey"`
	PreSharedKey string   `json:"preSharedKey,omitempty"`
	KeepAlive    int      `json:"keepAlive,omitempty"`
	AllowedIPs   []string `json:"allowedIPs,omitempty"`
}

type ParserWireguard struct {
	SecretKey      string   `json:"secretKey"`
	Address        []string `json:"address,omitempty"`
	Peers          []Peer   `json:"peers"`
	NoKernelTun    bool     `json:"noKernelTun,omitempty"`
	MTU            int      `json:"mtu,omitempty"`
	Reserved       []int    `json:"reserved,omitempty"`
	Workers        int      `json:"workers,omitempty"`
	DomainStrategy string   `json:"domainStrategy,omitempty"`
	*StreamField
}

func (that *ParserWireguard) Parse(rawUri string) {
	u, err := url.Parse(rawUri)
	if err != nil {
		return
	}

	that.StreamField = &StreamField{}

	that.SecretKey = u.User.Username()
	query := u.Query()

	that.Address = strings.Split(query.Get("address"), ",")
	that.NoKernelTun, _ = strconv.ParseBool(query.Get("noKernelTun"))
	that.MTU, _ = strconv.Atoi(query.Get("mtu"))
	that.Workers, _ = strconv.Atoi(query.Get("workers"))
	that.DomainStrategy = query.Get("domainStrategy")

	reserved := query.Get("reserved")
	if reserved != "" {
		for _, r := range strings.Split(reserved, ",") {
			i, _ := strconv.Atoi(r)
			that.Reserved = append(that.Reserved, i)
		}
	}

	peer := Peer{
		Endpoint:     u.Host,
		PublicKey:    query.Get("publicKey"),
		PreSharedKey: query.Get("preSharedKey"),
	}
	peer.KeepAlive, _ = strconv.Atoi(query.Get("keepAlive"))
	peer.AllowedIPs = strings.Split(query.Get("allowedIPs"), ",")
	that.Peers = append(that.Peers, peer)
}

func (that *ParserWireguard) GetAddr() string {
	if len(that.Peers) > 0 {
		return that.Peers[0].Endpoint
	}
	return ""
}

func (that *ParserWireguard) GetPort() int {
	if len(that.Peers) > 0 {
		_, portStr, _ := net.SplitHostPort(that.Peers[0].Endpoint)
		port, _ := strconv.Atoi(portStr)
		return port
	}
	return 0
}

func (that *ParserWireguard) Show() {
	jsonData, _ := json.MarshalIndent(that, "", "  ")
	fmt.Println(string(jsonData))
}
