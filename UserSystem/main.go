package main

import (
	"net/http"
	"twiter/UserSystem/Logic"
	"twiter/UserSystem/utitl"
)




func main(){

	utitl.MysqlInit()

	utitl.RedisInit()

	//mysqlInit()

	//register
	//Login1


	http.HandleFunc("/register", Logic.Register)

	http.HandleFunc("/Login1", Logic.Login1)




	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		//    logger.Fatal("ListenAndServe: ", err)
	}


}
