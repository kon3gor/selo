package main

import (
	"fmt"
	"unsafe"

	"github.com/kon3gor/selo"
	sk "github.com/kon3gor/selo/examples/simple/selokeys"
)

//go:generate selogen -type=MyClient
type MyClient struct {
	msg string
}

//go:generate selogen -type=MyClient2
type MyClient2 struct {
	client *MyClient
}

func NewMyClient() *MyClient {
	return &MyClient{
		msg: "hello world",
	}
}

func NewMyClient2() *MyClient2 {
	return &MyClient2{
		client: selo.Get[*MyClient](sk.MyClient),
	}
}

func main() {
	selo.Init()

	selo.Unique(sk.MyClient, NewMyClient)
	selo.Unique(sk.MyClient2, NewMyClient2, selo.WithLazy(true))

	v := selo.Get[*MyClient2](sk.MyClient2)

	fmt.Println(v.client.msg)
}

