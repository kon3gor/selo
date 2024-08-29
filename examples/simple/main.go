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

func NewMyClientOther() *MyClient {
	return &MyClient{
		msg: "bye world",
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
		SetTag("client1").
		SetFactory(NewMyClient).
		Record()

	selo.
		Unique[*MyClient]().
		SetLazy(true).
		SetTag("client2").
		SetFactory(NewMyClientOther).
		Record()

	v1 := selo.GetTagged[*MyClient]("client1")
	v2 := selo.GetTagged[*MyClient]("client2")

	fmt.Println(v1.msg)
	fmt.Println(v2.msg)
}
