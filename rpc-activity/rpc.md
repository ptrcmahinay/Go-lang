# Lab Activity 1 - RPC 
## Overview
In this lab, we learned about Remote Procedure Calls (RPC) using the Go programming language. The goal was to build a basic RPC client (server was provided) that can communicate over a network and perform simple arithmetic operations like addition, subtraction, multiplication, and division.

### Files in Lab Activity 1
1. [server.go](./server/server.go):  contains the logic for setting up an RPC server and performing arithmetic operations.
2. [client.go](./client/client.go): contains the client logic to connect to the server and perform arithmetic operations.

### Steps to Run
1. Start the Server:
   - Open a terminal and navigate to the server folder.
   - Run the server:
```cd server```
```go run server.go```
    - it should says: ```RPC server is listening on port 3000```

2. Start the Client: 
   - Open another terminal and navigate to the client folder.
   - Run the client: ```cd client```
```go run client.go```
