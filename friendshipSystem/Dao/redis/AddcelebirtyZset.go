package cache

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func AddcelebirtyZset(userId int)error{
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)  //这里面要打印日志
		return errors.New(fmt.Sprintf("redis.Dial err=", err))
	}
	defer conn.Close()
	celebrityKey := "celebrity"
	_, err = conn.Do("sadd", celebrityKey,userId)
	if err != nil {
		fmt.Println("sadd key error,err=", err) //这里面要打印日志
		return errors.New(fmt.Sprintf("sadd key error,err=", err))
	}
	return nil
}
