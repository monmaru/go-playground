package main

import "log"
import "net/rpc"
import "net/http"
import "net"
import "net/rpc/jsonrpc"

type Calculator int

func (c *Calculator) Mutiply(args Args, result *int) error {
	log.Printf("Multiply called: %d %d\n", args.A, args.B)
	*result = args.A * args.B
	return nil
}

type Args struct {
	A, B int
}

func main() {
	calculator := new(Calculator)
	server := rpc.NewServer()
	server.Register(calculator)
	http.Handle(rpc.DefaultRPCPath, server)
	log.Println("start http listenning :18888")
	listener, err := net.Listen("tcp", ":18888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
