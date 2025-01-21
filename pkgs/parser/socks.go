package parser

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/gvcgo/goutils/pkgs/gtui"
)

type ParserSocks struct {
	Address string
	Port    int
	User    string
	Pass    string
}

func (that *ParserSocks) Parse(rawUri string) {
	if u, err := url.Parse(rawUri); err == nil {
		that.Address = u.Hostname()
		that.Port, _ = strconv.Atoi(u.Port())
		that.User = u.User.Username()
		that.Pass, _ = u.User.Password()
	} else {
		gtui.PrintError(err)
		fmt.Println(rawUri)
		return
	}
}

func (that *ParserSocks) GetAddr() string {
	return that.Address
}

func (that *ParserSocks) GetPort() int {
	return that.Port
}

func (that *ParserSocks) Show() {
	fmt.Printf("addr: %s, port: %v, user: %s, pass: %s\n",
		that.Address,
		that.Port,
		that.User,
		that.Pass)
}

func SocksTest() {
	type Socks struct {
		Socks []string `json:"Socks"`
	}

	t := &Socks{}
	content, _ := os.ReadFile(`C:\Users\moqsien\data\projects\go\src\vpnparser\misc\socks.json`)
	json.Unmarshal(content, t)
	i := 0
	for _, rawUri := range t.Socks {
		p := &ParserSocks{}
		p.Parse(rawUri)
		if p.Address != "" {
			i++
		}
		p.Show()
	}
	fmt.Println("total: ", i, len(t.Socks))
}
