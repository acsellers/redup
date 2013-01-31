package main

import "flag"
import "fmt"

var RedisPass = flag.String("password", "", "Password for the redis server")

func ParseVariables() {
	flag.Parse()
}

func main() {
	ParseVariables()
	ConnectToRedis()
	fmt.Println("Starting Http Server on port", *PortNum)
	StartHttpServer()
}
