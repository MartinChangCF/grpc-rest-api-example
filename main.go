package main

import "github.com.tw/grpc-rest-api-example/server"

func main() {
	g := server.New()
	g.Start()
	g.WaitStop()
}
