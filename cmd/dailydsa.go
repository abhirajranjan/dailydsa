package main

import "github.com/abhirajranjan/dailydsa/internal/server"

func main() {
	a := server.Serve()
	a.Run(":8080")
}
