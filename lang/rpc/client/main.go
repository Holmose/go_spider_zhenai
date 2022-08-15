package main

import (
	rpcdemo "PRO02/lang/rpc"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div",
		rpcdemo.Args{
			A: 10,
			B: 3,
		}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result, err)
	}
	err = client.Call("DemoService.Div",
		rpcdemo.Args{
			A: 10,
			B: 0,
		}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result, err)
	}

}
