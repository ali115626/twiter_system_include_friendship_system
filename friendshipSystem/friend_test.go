package main

import (
	"fmt"
	"testing"
	rpclogic "twiter/friendshipSystem/Dao/rpc/logic"
)
//TestAdd(t *testing.T) {
//
//userId:=1
//resultStr,err:=rpc.GetFriendList(userId)
//if err!=nil{
//
//}
//fmt.Println(resultStr)
//}
func TestAdd(t *testing.T){
	userId :="10"


	commonFriendList,starList,err:=rpclogic.SearchStarCommonFriendsLogic(userId)
	if err!=nil{

	}

	fmt.Println("commonFriendList==",commonFriendList)
	fmt.Println("starList===",starList)



	//userId :=1
	//err:= Dao.ModifyIsCelebrityStatus(userId)
	//if err !=nil{
	//	fmt.Println(err)
	//	//return nil,err
	//}

	//fmt.Println("friendList====friendList===",friendList)



}


