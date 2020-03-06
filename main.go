package main

import "github.com.tw/intri-type/server"

func main() {
	g := server.New()
	g.Start()
	g.WaitStop()
}
