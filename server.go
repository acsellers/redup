package main

import "flag"
import "fmt"
import "github.com/simonz05/godis/redis"
import "sync"

var RedisPass = flag.String("password", "", "Password for the redis server")
var (
	Conn     *redis.Client
	Starting sync.WaitGroup
)

func ParseVariables() {
	flag.Parse()
}
func ConnectToRedis() {
	Conn = redis.New("tcp:127.0.0.1:6379", 0, *RedisPass)
}
func main() {

	ParseVariables()
	ConnectToRedis()
	fmt.Println("Starting Http Server on port", *PortNum)
	StartHttpServer()
}
