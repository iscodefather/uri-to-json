package parser

import (
	"fmt"
	"net/url"
	"os"
	"encoding/json"
	"strconv"

	"github.com/gvcgo/goutils/pkgs/gtui"
)

type ParserHTTP struct {
	Address string
	Port    int
	User    string
	Pass    string
}

func (that *ParserHTTP) Parse(rawUri string) {
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

func (that *ParserHTTP) GetAddr() string {
	return that.Address
}

func (that *ParserHTTP) GetPort() int {
	return that.Port
}

func (that *ParserHTTP) Show() {
	fmt.Printf("addr: %s, port: %v, user: %s, pass: %s\n",
		that.Address,
		that.Port,
		that.User,
		that.Pass)
}

func HTTPTest() {
	type HTTP struct {
		HTTP []string `json:"HTTP"`
	}

	t := &HTTP{}
	content, _ := os.ReadFile(`C:\Users\moqsien\data\projects\go\src\vpnparser\misc\http.json`)
	json.Unmarshal(content, t)
	i := 0
	for _, rawUri := range t.HTTP {
		p := &ParserHTTP{}
		p.Parse(rawUri)
		if p.Address != "" {
			i++
		}
		p.Show()
	}
	fmt.Println("total: ", i, len(t.HTTP))
}