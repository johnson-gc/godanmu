package godanmu

import (
	"time"
	"net"
	"log"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)
type Client struct {
	//房间号
	rid string
	//callback
	fn  func(p *Packet)
}
const (
	//弹幕服务器地址
	ADDRESS    = "openbarrage.douyutv.com:8601"
	//发送心跳包间隔
	KEEP_ALIVE = 45 * time.Second
)
var msgChan = make(chan *Packet)
//弹幕客户端
//
func (c *Client) run() {
	conn, err := net.Dial("tcp", ADDRESS)
	if err != nil {
		log.Fatal("连接弹幕服务器失败")
	}
	//登陆
	send(conn, NewPacket(int16(689), &map[string]string{"type": "loginreq", "roomid": c.rid}))
	//接收弹幕goroutine
	go receive(conn,c)
	//每45秒发送一个keep_alive包
	ticker := time.NewTicker(KEEP_ALIVE)

	defer func() {
		ticker.Stop()
		conn.Close()
	}()

	for {
		select {
		case p := <-msgChan:
			c.fn(p)
		case <-ticker.C:
			send(conn, NewPacket(int16(689), &map[string]string{"type": "keeplive", "tick": strconv.FormatInt(time.Now().Unix(), 10)}))
		}

	}
}

func send(conn net.Conn, p *Packet) {
	conn.Write(p.toRaw())
}
func receive(conn net.Conn,c *Client) {
	bufReader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(bufReader)
	scanner.Split(splitByTypeFunc)

	for {
		if ok := scanner.Scan(); !ok {
			fmt.Println("not ok")
			continue
		}

		b := scanner.Bytes()
		length := len(b)

		if length < 20 { //只会执行一次
			send(conn, NewPacket(int16(689), &map[string]string{"type": "joingroup", "rid":  c.rid, "gid": "-9999"}))
			continue
		}

		b = b[:length-13]

		p := &Packet{
			ptype: new(int16),
		}

		p.fromRaw(b)

		msgChan <- p
	}
}
//按\x00type@=分割,不需要parse协议
func splitByTypeFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := strings.Index(string(data), "\x00type@="); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return
}
