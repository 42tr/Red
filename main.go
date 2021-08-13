package main

import (
	"red/conf"
	"red/server"
)

func init() {
	conf.Init()
}

func main() {
	r := server.NewRouter()
	_ = r.Run(":8081")
}
