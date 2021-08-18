package driver

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type Client struct {
	Connection net.Conn
	Buffer *bufio.ReadWriter
}

func NewClient(connString string) *Client {
	conn, _ := net.Dial("tcp", connString)
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)
	buf := bufio.NewReadWriter(rd, wr)
	return &Client{conn, buf}
}

func (client *Client) Send(s string) string {
	_, _ = client.Buffer.WriteString(s)
	_ = client.Buffer.Flush()
	ret, _ := client.Buffer.ReadString('\n')
	return ret[:len(ret)-1]
}

func (client *Client) Set(k string, v string) chan SetResult {
	ret := make(chan SetResult)

	go func(k string, v string) {
		s := fmt.Sprintf("SET$%s$%s\n", k, v)
		serial := client.Send(s)
		obj := SetResult{}
		_ = json.Unmarshal([]byte(serial), &obj)
		ret <- obj
	}(k, v)

	return ret
}

func (client *Client) Get(k string) chan GetResult {
	ret := make(chan GetResult)

	go func(k string) {
		s := fmt.Sprintf("GET$%s\n", k)
		serial := client.Send(s)
		obj := GetResult{}
		_ = json.Unmarshal([]byte(serial), &obj)
		ret <- obj
	}(k)

	return ret
}

func (client *Client) Del(k string) chan DelResult {
	ret := make(chan DelResult)

	go func(k string) {
		s := fmt.Sprintf("DEL$%s\n", k)
		serial := client.Send(s)
		obj := DelResult{}
		_ = json.Unmarshal([]byte(serial), &obj)
		ret <- obj
	}(k)

	return ret
}

func (client *Client) Exit() {
	_ = client.Connection.Close()
}
