package main

import (
	rpcdemo "PRO02/lang/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(rpcdemo.DemoService{})
	if err != nil {
		log.Panic(err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}

/* 使用方法
telnet localhost 1234
{"method":"abc.def"}
>> {"id":null,"result":null,"error":"rpc: can't find service abc.def"}
{"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
>> {"id":1,"result":0.75,"error":null}
{"method":"DemoService.Div","params":[{"A":3,"B":0}],"id":12}
>> {"id":12,"result":null,"error":"division by zero"}
*/
