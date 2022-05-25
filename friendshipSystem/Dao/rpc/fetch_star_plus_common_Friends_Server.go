package rpc

import (
	"context"
	"fmt"
	rpclogic "twiter/friendshipSystem/Dao/rpc/logic"
	"twiter/friendshipSystem/Dao/rpc/message2"

	//"try_rpc_proto/message"

	//"try_rpc_proto/message"
)

//todo 这个就是grpc框架下的client和server之间的通信

type ImplementFriendsPlusStars struct {
}

func (this *ImplementFriendsPlusStars) GetFriendsListCommonPersonPlusCelebrity(ctx context.Context,request *message2.FetchRequestCommonStar) (*message2.FriendsResponseCommonStars, error) {

	//strconv.Atoi(request.UserId)
	commonFriendList,starList,err:=rpclogic.SearchStarCommonFriendsLogic(int(request.UserId))
	if err!=nil{
		return nil,err
	}

	result:=message2.FriendsResponseCommonStars{}
	result.CommonPersonList=commonFriendList
	result.StarPersonList=starList
	fmt.Println("the friendList are",result)
	return &result,nil
}
//type ImplementFriendsPlusStars struct{
//}
//	GetOrderInfo(context.Context, *OrderRequest) (*OrderInfo, error)

//GetFriendsListCommonPersonPlusCelebrity(context.Context, *FetchRequest) (*FriendsResponse, error)

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
