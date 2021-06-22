package main

import (
	"red/conf"
	"red/server"
)

func main() {
	conf.Init()

	r := server.NewRouter()
	_ = r.Run(":8081")
}
