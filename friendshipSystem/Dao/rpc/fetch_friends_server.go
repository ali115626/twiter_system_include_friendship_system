package rpc

import (
	"context"
	"fmt"
	"strconv"
	Dao "twiter/friendshipSystem/Dao/mysql"
	"twiter/friendshipSystem/Dao/rpc/message"

	//"try_rpc_proto/message"

	//"try_rpc_proto/message"
)

//todo 这个就是grpc框架下的client和server之间的通信


type ImplementFriends struct{
}
//	GetOrderInfo(context.Context, *OrderRequest) (*OrderInfo, error)

func (this *ImplementFriends) GetFriendsList(ctx context.Context, request *message.FetchRequest) (*message.FriendsResponse, error) {
	//这里面应该是执行一些逻辑后返回
	userId :=int(request.UserId)
	friendList,err := Dao.SelectFriendsFromFriendTable(userId)
	if err !=nil{
		fmt.Println(err)
		return nil,err
	}

	fmt.Println("friendList====friendList===",friendList)
	//todo 这个FriendList里面要加上自己  自己也要fanout进去 因为还需要做成timeline
	friendList =append(friendList,userId)
	//for循环一下 将string弄成int
	var friendListString []string
	//这里总是个坑
	for _,friend := range friendList{
		friendListString=append(friendListString,strconv.Itoa(friend))
	}
	//FriendListMar,err:=json.Marshal(FriendList)
	response := message.FriendsResponse{
		FriendList:friendListString,
	}
	fmt.Println("response===server===",response)
	//fmt.Fprintf(w,string(FriendListMar))
	return &response,nil
}
//
//func main(){
//	server := grpc.NewServer()
//	//Orderinfo2:=new(Orderinfo1)
//	//orderServiceClient.GetOrderInfo
//	message.RegisterFriendShipServiceServer(server,new(ImplementFriends))
//	//message.RegisterOrderServiceServer(server,new(ImplementFriends))
//	lis,err :=net.Listen("tcp",":9000")
//	if err != nil{
//		panic(err.Error())
//	}
//	server.Serve(lis)
//}