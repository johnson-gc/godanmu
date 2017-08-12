package main

import(
	"flag"
	"fmt"
	"time"
	"github.com/zypperin/godanmu"
)

//./example -room=67373
var roomid = flag.String("room", "67373", "房间号")

var enterUserMap = make(map[string]int)
var dmUserMap = make(map[string]int)


func onMsg(p *Packet){
	switch p.body["type"] {
	case "chatmsg":
		dmUserMap[p.body["uid"]] = 1
	case "uenter":
		enterUserMap[p.body["uid"]] = 1
	}
}

func main() {
	flag.Parse()
    client := &Client{
		rid:*roomid,
		fn:onMsg,
	}
	tc := time.NewTicker(time.Second * 1)
	go func() {
		for t := range tc.C {
			fmt.Printf("【%v】累计进入直播间人数: %d, 参与发送弹幕人数: %d \r",t, len(enterUserMap),len(dmUserMap))
		}
	}()

	client.run()
}



