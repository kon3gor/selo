package main

import (
	"fmt"

	"github.com/kon3gor/selo"
)

type MyClient struct {
	msg string
}

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
		client: selo.Get[*MyClient](),
	}
}

func main() {
	selo.Init(selo.WithDebug(true))

	selo.
		Unique[*MyClient]().
		SetLazy(true).
		SetTag("client").
		SetFactory(NewMyClient).
		Record()

	selo.
		Unique[*MyClient2]().
		SetLazy(true).
		SetFactory(NewMyClient2).
		Record()

	v := selo.Get[*MyClient2]()

	fmt.Println(v.client.msg)
}
