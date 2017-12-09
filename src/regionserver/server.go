package main

import "./server"

func main() {
	// ws://192.168.1.51:8080/game
	// {"Code":1, "Info":"test"}

	gs := server.NewDerivedGameServer()
	gs.Start()
}
