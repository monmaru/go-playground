package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}

	var result int
	args := &Args{4, 5}
	err = client.Call("Calculator.Mutiply", args, &result)
	if err != nil {
		panic(err)
	}
	log.Printf("4 × 5 = %d\n", result)
}
