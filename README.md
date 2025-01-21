# vpnparser
uri-to-json parses VPN URI to xray-core Outbound.

## install

```bash
go install github.com/iscodefather/uri-to-json@latest
```

## commands

```bash
moqsien> vpnparser help

NAME:
   vpnparser.exe - vpnparser <Command> <SubCommand>...

USAGE:
   vpnparser.exe [global options] command [command options] [arguments...]

DESCRIPTION:
   vpnparser, download files from github for gvc.

COMMANDS:
   xray, x  Generate xray-core outbound from vpn url.
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## examples

```bash
moqsien> vpnparser s "ss://chacha20-ietf-poly1305:t0srmdxrm3xyjnvqz9ewlxb2myq7rjuv@4e168c3.h4.gladns.com:2377/?plugin=obfs-local\u0026obfs=tls\u0026obfs-host=(TG@WangCai_1)c68b799:50307#8DKJ|www.zyw.asia ZYW免费节点"

ss://chacha20-ietf-poly1305:t0srmdxrm3xyjnvqz9ewlxb2myq7rjuv@4e168c3.h4.gladns.com:2377/?plugin=obfs-local\u0026obfs=tls\u0026obfs-host=(TG@WangCai_1)c68b799:50307#8DKJ|www.zyw.asia ZYW免费节点
{
        "method": "chacha20-ietf-poly1305",
        "password": "t0srmdxrm3xyjnvqz9ewlxb2myq7rjuv",
        "plugin": "obfs-local\\u0026obfs=tls\\u0026obfs-host=(TG@WangCai_1)c68b799:50307",
        "server": "4e168c3.h4.gladns.com",
        "server_port": 2377,
        "tag": "proxy",
        "tls": {
                "enabled": false
        },
        "transport": {},
        "type": "shadowsocks"
}
```

```bash
moqsien> .\vpnparser.exe x '"vmess://{\"add\":\"us47.encrypted.my.id\",\"port\":\"80\",\"id\":\"4bf9b7e0-85d1-4a59-9a29-e6619dcd7c50\",\"aid\":\"0\",\"scy\":\"auto\",\"net\":\"ws\",\"v\":\"2\",\"ps\":\"美国_0828698\",\"host\":\"\",\"path\":\"/pSAXxD8Ib7FZloqUMG\",\"tls\":\"\",\"sni\":\"\",\"type\":\"none\",\"serverPort\":0,\"nation\":\"🇺🇸US\"}"'

vmess://{"add":"us47.encrypted.my.id","port":"80","id":"4bf9b7e0-85d1-4a59-9a29-e6619dcd7c50","aid":"0","scy":"auto","net":"ws","v":"2","ps":"美国_0828698","host":"","path":"/pSAXxD8Ib7FZloqUMG","tls":"","sni":"","type":"none","serverPort":0,"nation":"🇺🇸US"}
{
        "protocol": "vmess",
        "sendThrough": "0.0.0.0",
        "settings": {
                "vnext": [
                        {
                                "address": "us47.encrypted.my.id",
                                "port": 80,
                                "users": [
                                        {
                                                "alterId": 0,
                                                "id": "4bf9b7e0-85d1-4a59-9a29-e6619dcd7c50",
                                                "security": "auto"
                                        }
                                ]
                        }
                ]
        },
        "streamSettings": {
                "network": "ws",
                "security": "",
                "wsSettings": {
                        "headers": {
                                "Host": ""
                        },
                        "path": "/pSAXxD8Ib7FZloqUMG"
                }
        },
        "tag": "proxy"
}
```
