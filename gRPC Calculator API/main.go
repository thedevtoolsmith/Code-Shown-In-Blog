package main

import (
	server "gRPC-example/server"
	client "gRPC-example/client"
)


func main(){
	go server.Run()
	client.Run()
}