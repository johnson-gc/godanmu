package main

import(
	"flag"
	"fmt"
	"github.com/zypperin/godanmu"
)

//./example -room=67373
var roomid = flag.String("room", "67373", "房间号")

func main() {
	flag.Parse()
    client := &Client{
		rid:*roomid,
		fn:onMsg,
	}
	client.run()
}

func onMsg(p *Packet){
	fmt.Println(p.body)
}

