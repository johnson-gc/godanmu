# godanmu
斗鱼弹幕接口for go

# 安装
  go get github.com/zypperin/godanmu
# 使用
```go
package main

import(
	"flag"
	"fmt"
	"github.com/zypperin/godanmu"
)

var roomid = flag.String("room", "67373", "房间号")

func main() {
	flag.Parse()
    client := &godanmu.Client{
		Rid:*roomid,
		Fn:onMsg,
	}
	client.Run()
}

func onMsg(p *godanmu.Packet){
	fmt.Println(p.Body)
}
```
![Screen](/_examples/example.gif)
