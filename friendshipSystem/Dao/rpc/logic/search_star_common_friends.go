package rpclogic

import (
	"fmt"
	//"strconv"
	Dao "twiter/friendshipSystem/Dao/mysql"
)

func SearchStarCommonFriendsLogic(userIdInt int)([]string,[]string,error){
	//userId :=int(request.UserId)
	//friendList,err := Dao.SelectFriendsFromFriendTable(userId)
	//if err !=nil{
	//	fmt.Println(err)
	//	return nil,err
	////}
	//userIdInt,err:=strconv.Atoi(userId)
	//if err !=nil{
	//	fmt.Println(err)
	//	return nil,nil,err
	//}

	commonFriendList,err:=Dao.SelectCommonFriendsFromFriendTable(userIdInt)
	if err!=nil{
		fmt.Println(err)
		return nil,nil,err

	}
	starList,err:=Dao.SelectStarsFromFriendTable(userIdInt)
	if err!=nil{
		fmt.Println(err)
		return nil,nil,err

	}
	return commonFriendList,starList,nil





}



