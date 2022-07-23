package main

import (
	"fmt"
	"os"
	"slgserver/config"
	"slgserver/net"
	"slgserver/server/slgserver/run"
)

func getServerAddr() string {
	host := config.File.MustValue("slgserver", "host", "127.0.0.1")
	port := config.File.MustValue("slgserver", "port", "8001")
	return host + ":" + port
}

func main() {
	fmt.Println(os.Getwd())
	run.Init()
	needSecret := config.File.MustBool("slgserver", "need_secret", false)
	s := net.NewServer(getServerAddr(), needSecret)
	s.Router(run.MyRouter)
	s.Start()
}
