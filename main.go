package main

import (
	"github.com/wteja/go-covid-api/server"
)

func main() {
	r := server.CreateServer()
	r.Run()
}
