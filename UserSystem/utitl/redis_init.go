package utitl

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var Conn redis.Conn


func RedisInit(){


	Conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}


	fmt.Println("conn suc...", Conn)
	defer Conn.Close()




}

