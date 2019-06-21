package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIp string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"shanghai_VPn","serverIP":"127.0.0.1"},{"serverName":"beijing_VPn","serverIP":"47.93.15.195"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}
