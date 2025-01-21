package main

import (
	"github.com/iscodefather/uri-to-json/pkgs/cmd"
	_ "github.com/iscodefather/uri-to-json/pkgs/outbound/xray"
	_ "github.com/iscodefather/uri-to-json/pkgs/parser"
)

func main() {
	// parser.VlessTest()
	// parser.TrojanTest()
	// parser.SSRTest()
	// parser.WireguardTest()
	// parser.SocksTest()

	// s := xray.GetPattern()
	// fmt.Println(s)
	// xray.TestVmess()
	// xray.TestTrojan()
	// xray.TestSS()

	// sing.TestVmess()
	// sing.TestVless()
	// sing.TestTrojan()
	// sing.TestSS()

	cmd.StartApp()

	// rawUri := "vmess://{\"add\":\"ms.shabijichang.com\",\"port\":\"80\",\"id\":\"f1865e50-2510-46d1-bcb2-e00b4b656305\",\"aid\":\"0\",\"scy\":\"auto\",\"net\":\"ws\",\"v\":\"2\",\"ps\":\"未知_0915019\",\"host\":\"\",\"path\":\"\",\"tls\":\"\",\"sni\":\"\",\"type\":\"none\",\"serverPort\":0,\"nation\":\"🏁ZZ\"}"
	// fmt.Println(p.GetOutbound())
}
