# godanmu
斗鱼弹幕接口for go

# 安装
  go get github.com/zypperin/godanmu
# 使用
```go
package main

import(
	"fmt"
	"github.com/zypperin/godanmu"
)

func main() {
  client := &godanmu.Client{
	Rid:"67373",
	Fn:onMsg,
  }
  client.Run()
}

func onMsg(p *godanmu.Packet){
	fmt.Println(p.Body)
}
```
