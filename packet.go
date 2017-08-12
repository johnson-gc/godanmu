package godanmu

import (
	"strings"
	"bytes"
	"encoding/binary"
)
//弹幕协议格式: http://dev-bbs.douyutv.com
type Packet struct {
	ptype *int16
	body map[string]string
}

func (p *Packet) fromRaw(buf []byte){
	p.body = *(deserialize(string(buf)))
}

func (p *Packet) toRaw()([]byte){
	sb := serialize(&p.body)

	rawLen := len(sb)+8

	buf := new(bytes.Buffer)

	binary.Write(buf, binary.LittleEndian, int32(rawLen))
	binary.Write(buf, binary.LittleEndian, int32(rawLen))
	binary.Write(buf, binary.LittleEndian, p.ptype)
	binary.Write(buf, binary.LittleEndian, byte(0))
	binary.Write(buf, binary.LittleEndian, byte(0))
	binary.Write(buf, binary.LittleEndian, []byte(sb))

	return buf.Bytes()
}

func NewPacket(t int16,b *map[string]string)*Packet{
	return &Packet{
		ptype:&t,
		body:*b,
	}
}

func serialize(body *map[string]string)string{
	var buffer bytes.Buffer

	for k, v := range *body {
		buffer.WriteString(escape(k))
		buffer.WriteString("@=")
		buffer.WriteString(escape(v))
		buffer.WriteString("/")
	}
	buffer.WriteString("\x00")
	return buffer.String()
}

func deserialize(s string)*map[string]string{
	result := make(map[string]string)
	pairs := strings.Split(s,"/")
	for _,p := range pairs{
		if len(p) ==0{
			continue
		}
		kv := strings.Split(p,"@=")

		if len(kv) !=2 {
			continue
		}

		k := unescape(kv[0])
		v := unescape(kv[1])
		result[k]=v
	}
	return &result
}

func escape(s string) string{
	v := strings.Replace(s,"@","@A",-1)
	v = strings.Replace(v,"/","@S",-1)
	return v
}

func unescape(s string) string{
	v := strings.Replace(s,"@S","/",-1)
	v = strings.Replace(v,"@A","@",-1)
	return v
}