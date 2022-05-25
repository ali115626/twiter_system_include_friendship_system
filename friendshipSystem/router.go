package main

import (
	"fmt"
	"net/http"
	Dao "twiter/friendshipSystem/Dao/mysql"
	"twiter/friendshipSystem/Dao/rpc"
	"twiter/friendshipSystem/friendService"
	//"twiter/tweetSystem/friendService"
)
//
//
//func init() {
//	//var err error
//
//	Dao.MysqlConn,_=Dao.MysqlInit()
//	//fmt.Println("Dao.MysqlConn=",Dao.MysqlConn)
//	//if err!=nil{
//	//	fmt.Println(err)
//	//}
//	//
//	//
//	//
//	//}
//
//}
func main(){
		//哈哈这个就变成了在两个端口上监听了
		_,_=Dao.MysqlInit()
		go rpc.RpcFriendServer()
		go rpc.RpcFriendCommonStarServer()
		//todo  一个程序  连个端口监听  根本就无法到达下面的接口
		var err error
		//Dao.MysqlConn,err=Dao.MysqlInit()
		//fmt.Println("Dao.MysqlConn=", Dao.MysqlConn)
		if err != nil {
			fmt.Println(err)
		}

		////////-----------------------------------------------------------------------------------------------
		//////-------------------------------------------------------------------------------------------------
		http.HandleFunc("/add_friends", friendService.AddFriends)

		//InsertFriendNew
		http.HandleFunc("/add_friends_new", friendService.AddFriendsNew)

		//ShowSomeOneFriends
		//http.HandleFunc("/show_friends", friendService.ShowSomeOneFriends)
		//http.HandleFunc("/add_friends_new_add_redis", friendService.AddFriendsNewAddRedis)

		err = http.ListenAndServe(":9092", nil) // 设置监听的端口
		if err != nil {
			//    logger.Fatal("ListenAndServe: ", err)
		}

		fmt.Println("111222222")

	}
