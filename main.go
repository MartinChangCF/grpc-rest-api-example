package main

import "./server"

func main() {
	g := server.New()
	g.Start()
	g.WaitStop()
}
