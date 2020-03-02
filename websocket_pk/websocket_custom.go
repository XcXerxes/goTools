package main

import (
	"crypto/sha1"
	"encoding/base64"
	"io"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn)  {
	content := make([]byte, 1024)
	_, err := conn.Read(content)
	log.Println(string(content))
	if err != nil {
		log.Println(err)
	}
	isHttp := false
	if string(content[0:3]) == "GET" {
		isHttp = true
	}
	log.Println("isHttp============", isHttp)
	if isHttp {
		headers := paraseHandshake(string(content))
		log.Println("headers===========", headers)
		secWebsocketKey := headers["Sec-WebSocket-Key"]
		// NoTE: 省略其他验证
		guid := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
		// 计算Sec-WebSocket-Key
		h := sha1.New()
		log.Println("accept raw:", secWebsocketKey + guid)
		io.WriteString(h, secWebsocketKey + guid)
		accept := make([]byte, 28)
		base64.StdEncoding.Encode(accept, h.Sum(nil))
		log.Println(string(accept))
		response := "HTTP/1.1 101 Switching Protocols\r\n"
		response = response + "Sec-WebSocket-Accept: " + string(accept) + "\r\n"
		response = response + "Connection: Upgrade\r\n"
		response = response + "Upgrade: websocket\r\n\r\n"
		log.Println("response:", response)
		if lenth, err := conn.Write([]byte(response)); err != nil {
			log.Println(err)
		} else {
			log.Println("send len:", lenth)
		}
		wssocket := NewWsSocket(conn)
		for{
			data, err := wssocket.Conn.ReadI
		}
	}
}

type WsSocket struct {
	MaskingKey []byte
	Conn net.Conn
}
func NewWsSocket(conn net.Conn) *WsSocket  {
	return &WsSocket{Conn:conn}
}

func (this *WsSocket) ReadIframe() (data []byte, err error) {
	err = nil
	// 第一个字节 FIN + RSV1-3 + OPCODE
	opcodeByte := make([]byte, 1)
	this.Conn.Read(opcodeByte)

}

func paraseHandshake(content string) map[string]string  {
	headers := make(map[string]string, 10)
	lines := strings.Split(content, "\r\n")
	for _, line := range lines {
		if len(line) >= 0 {
			words := strings.Split(line, ":")
			if len(words) == 2 {
				headers[strings.Trim(words[0], " ")] = strings.Trim(words[1], " ")
			}
		}
	}
	return headers
}
