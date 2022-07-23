package main

import (
	"fmt"
	"os"
	"slgserver/config"
	"slgserver/net"
	"slgserver/server/loginserver"
)

func getLoginServerAddr() string {
	host := config.File.MustValue("loginserver", "host", "127.0.0.1")
	port := config.File.MustValue("loginserver", "port", "8003")
	return host + ":" + port
}

func main() {
	fmt.Println(os.Getwd())
	loginserver.Init()
	needSecret := config.File.MustBool("loginserver", "need_secret", false)
	s := net.NewServer(getLoginServerAddr(), needSecret)
	s.Router(loginserver.MyRouter)
	s.Start()
}
