package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gvcgo/goutils/pkgs/gtui"
)

type ParserWireguard struct {
	SecretKey   string
	Address     []string
	Endpoint    string
	PublicKey   string
	MTU         int
	Reserved    []int
	Workers     int
	StreamField *StreamField
}

func (that *ParserWireguard) Parse(rawUri string) {
	type WireguardConfig struct {
		PrivateKey string `json:"PrivateKey"`
		AddrV4     string `json:"AddrV4"`
		AddrV6     string `json:"AddrV6"`
		Endpoint   string `json:"Endpoint"`
		PublicKey  string `json:"PublicKey"`
		MTU        int    `json:"MTU"`
		Reserved   []int  `json:"Reserved"`
		Workers    int    `json:"Workers"`
	}

	var config WireguardConfig
	if err := json.Unmarshal([]byte(rawUri), &config); err != nil {
		gtui.PrintError(err)
		fmt.Println("Invalid JSON:", rawUri)
		return
	}

	that.SecretKey = config.PrivateKey
	that.Address = []string{config.AddrV4, config.AddrV6}
	that.Endpoint = config.Endpoint
	that.PublicKey = config.PublicKey
	that.MTU = config.MTU
	that.Reserved = config.Reserved
	that.Workers = config.Workers
}

func (that *ParserWireguard) Show() {
	fmt.Printf("SecretKey: %s\n", that.SecretKey)
	fmt.Printf("Addresses: %v\n", that.Address)
	fmt.Printf("Endpoint: %s\n", that.Endpoint)
	fmt.Printf("PublicKey: %s\n", that.PublicKey)
	fmt.Printf("MTU: %d\n", that.MTU)
	fmt.Printf("Reserved: %v\n", that.Reserved)
	fmt.Printf("Workers: %d\n", that.Workers)
}

func WireguardTest() {
	type Wireguard struct {
		Configs []string `json:"Wireguard"`
	}

	w := &Wireguard{}
	content, err := os.ReadFile(`C:\Users\moqsien\data\projects\go\src\vpnparser\misc\wireguard.json`)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	json.Unmarshal(content, w)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	i := 0
	for _, rawUri := range w.Configs {
		p := &ParserWireguard{}
		p.Parse(rawUri)

		if p.Endpoint != "" {
			i++
			fmt.Printf("SecretKey: %s, Endpoint: %s, PublicKey: %s\n", p.SecretKey, p.Endpoint, p.PublicKey)
			fmt.Printf("Addresses: %v\n", p.Address)
			fmt.Printf("MTU: %d, Reserved: %v, Workers: %d\n", p.MTU, p.Reserved, p.Workers)
		}
	}
	fmt.Println("total:", i, len(w.Configs))
}
