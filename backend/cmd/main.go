package main

import "backend/internal/server"

func main() {
	serv := server.New(":8090")
	serv.Start()
}
