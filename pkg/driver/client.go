package driver

import (
	"encoding/gob"
	"net"
)

type Client struct {
	Conn net.Conn
	Encoder *gob.Encoder
	Decoder *gob.Decoder
}

func NewClient(connection string) *Client {
	gob.Register(SetCommand{})
	gob.Register(GetCommand{})
	gob.Register(DelCommand{})
	gob.Register(ExitCommand{})
	gob.Register(SetResult{})
	gob.Register(DelResult{})

	conn, _ := net.Dial("tcp", connection)
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)
	return &Client{conn, encoder, decoder}
}

func (client *Client) Send(msg Message) (interface{}, error) {
	// send gob-encoded message over the TCP connection
	err := client.Encoder.Encode(&msg)
	if err != nil {
		return Message{}, err
	}

	// receive gob-encoded response from the TCP connection
	err = client.Decoder.Decode(&msg)
	if err != nil {
		return Message{}, err
	}

	return msg.Payload, nil
}

func (client *Client) Set(key string, val interface{}) chan SetResult {
	c := make(chan SetResult)

	// query the database with a SetCommand object, update the channel with a SetResult object
	go func(c chan SetResult) {
		msg := Message{
			Payload: SetCommand{
				Key: key,
				Value: val,
			},
		}
		res, _ := client.Send(msg)
		c <- res.(SetResult)
	}(c)

	return c
}

func (client *Client) Get(key string) chan interface{} {
	c := make(chan interface{})

	// query the database with a GetCommand object, update the channel with matched value
	go func(c chan interface{}) {
		msg := Message{
			Payload: GetCommand{
				Key: key,
			},
		}
		res, _ := client.Send(msg)
		c <- res
	}(c)

	return c
}

func (client *Client) Del(key string) chan DelResult {
	c := make(chan DelResult)

	// query the database with a DelCommand object, update the channel with a DelResult object
	go func(c chan DelResult) {
		msg := Message{
			Payload: DelCommand{
				Key: key,
			},
		}
		res, _ := client.Send(msg)
		c <- res.(DelResult)
	}(c)

	return c
}
