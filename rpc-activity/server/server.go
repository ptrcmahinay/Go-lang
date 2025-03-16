package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// Define a struct for the Calculator
type Calculator struct{}

// Define a struct to hold the arguments for RPC calls
type Args struct {
	A, B float64
}

// Server Methods  - it stores the result in reply var
func (c *Calculator) Add(args *Args, result *float64) error {
	*result = args.A + args.B
	return nil
}

func (c *Calculator) Subtract(args *Args, result *float64) error {
	*result = args.A - args.B
	return nil
}

func (c *Calculator) Multiply(args *Args, result *float64) error {
	*result = args.A * args.B
	return nil
}

func (c *Calculator) Divide(args *Args, result *float64) error {
	if args.B == 0 {
		return fmt.Errorf("cannot divide by zero")
	}
	*result = args.A / args.B
	return nil
}

func main() {
	// Create a new RPC server instance
	calculator := new(Calculator)
	err := rpc.Register(calculator)
	if err != nil {
		log.Fatal("Error registering RPC server:", err)
	}

	// Start listening on a TCP port
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("RPC Server is listening on port 3000...")

	// Accept incoming client connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn) // Handle the client request concurrently
	}
}
