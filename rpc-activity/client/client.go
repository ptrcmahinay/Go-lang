package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B float64
}

func main() {
	serverAddress := "192.168.1.94:3000"

	client, err := rpc.Dial("tcp", serverAddress)

	if err != nil {
		fmt.Println("Error connecting to RPC server:", err)
		return
	}
	defer client.Close()

	args := Args{
		A: 22,
		B: 10,
	}

	var result float64

	err = client.Call("Calculator.Add", args, &result)
	if err != nil {
		fmt.Println("RPC call error: ", err)
		return
	}

	fmt.Println("Result: ", result)
	fmt.Println("Name: Patricia Ann C. Mahinay")
}
