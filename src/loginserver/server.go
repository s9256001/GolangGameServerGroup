package main

import "./server"

func main() {
	gs := server.NewDerivedGameServer()
	gs.Start()
}
