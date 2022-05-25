package friendService

import (
	"fmt"
	"net/http"
	Dao "twiter/friendshipSystem/Dao/mysql"

	//"twiter/friendshipSystem/Dao"

	//Dao "twiter/tweetSystem/Dao/mysql"
)

func AddFriends(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	userId := requestMap["userId"][0]
	friendId := requestMap["friendId"][0]
	err=Dao.InsertFriend(userId,friendId)
	if err !=nil{
		return
	}
	fmt.Fprintf(w,"添加好友成功")
}
//这个的话就是把朋友信息加到了redis中的集合里面

func AddFriendsNew(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	requestMap := r.Form
	userId := requestMap["userId"][0]
	friendId := requestMap["friendId"][0]
	err=Dao.InsertFriendNew(userId,friendId)
	if err !=nil{
		return
	}
	fmt.Fprintf(w,"添加好友成功")
}


