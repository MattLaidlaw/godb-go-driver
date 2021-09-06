package driver

import (
	"fmt"
	"github.com/MattLaidlaw/go-jsonrpc2"
)

type Client struct {
	rpc *jsonrpc2.Client
}

func NewClient(address string) (*Client, error) {
	rpc, err := jsonrpc2.Dial(address)
	if err != nil {
		return nil, err
	}
	return &Client {
		rpc: rpc,
	}, nil
}

func (c *Client) Set(key string, val string) (float64, error) {
	res, err := c.rpc.Call("Handler.Set", key, val)
	if err != nil {
		return 0, err
	}
	if res.Error != (jsonrpc2.Error{}) {
		return 0, fmt.Errorf("%d: %s", res.Error.Code, res.Error.Message)
	}
	return res.Result.(float64), nil
}

func (c *Client) Get(key string) (string, error) {
	res, err := c.rpc.Call("Handler.Get", key)
	if err != nil {
		return "", err
	}
	if res.Error != (jsonrpc2.Error{}) {
		return "", fmt.Errorf("%d: %s", res.Error.Code, res.Error.Message)
	}
	return res.Result.(string), nil
}

func (c *Client) Del(key string) (float64, error) {
	res, err := c.rpc.Call("Handler.Del", key)
	if err != nil {
		return 0, err
	}
	if res.Error != (jsonrpc2.Error{}) {
		return 0, fmt.Errorf("%d: %s", res.Error.Code, res.Error.Message)
	}
	return res.Result.(float64), nil
}

//
//func (c *Client) Set(key KeyType, val ValType) (int, error) {
//	req := SetRequest{Key: key, Val: val}
//	var res SetResult
//	err := c.RPCClient.Call("Handler.Set", &req, &res)
//	return res.InsertedCount, err
//}
//
//func (c *Client) Get(key KeyType) (ValType, error) {
//	req := GetRequest{Key: key}
//	var res GetResult
//	err := c.RPCClient.Call("Handler.Get", &req, &res)
//	return res.Val, err
//}
//
//func (c *Client) Del(key KeyType) (int, error) {
//	req := DelRequest{Key: key}
//	var res DelResult
//	err := c.RPCClient.Call("Handler.Del", &req, &res)
//	return res.DeletedCount, err
//}
