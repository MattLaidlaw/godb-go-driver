package driver

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Client struct {
	RPCClient *rpc.Client
}

func NewClient(connectionString string) (*Client, error) {
	rpcClient, err := jsonrpc.Dial("tcp", connectionString)
	return &Client{RPCClient: rpcClient}, err
}

func (c *Client) Set(key KeyType, val ValType) (int, error) {
	req := SetRequest{Key: key, Val: val}
	var res SetResult
	err := c.RPCClient.Call("Handler.Set", &req, &res)
	return res.InsertedCount, err
}

func (c *Client) Get(key KeyType) (ValType, error) {
	req := GetRequest{Key: key}
	var res GetResult
	err := c.RPCClient.Call("Handler.Get", &req, &res)
	return res.Val, err
}

func (c *Client) Del(key KeyType) (int, error) {
	req := DelRequest{Key: key}
	var res DelResult
	err := c.RPCClient.Call("Handler.Del", &req, &res)
	return res.DeletedCount, err
}
