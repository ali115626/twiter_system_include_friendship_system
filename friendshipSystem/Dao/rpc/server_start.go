package rpc

import (
	"google.golang.org/grpc"
	"net"
	"twiter/friendshipSystem/Dao/rpc/message"
	"twiter/friendshipSystem/Dao/rpc/message2"
)

func RpcFriendServer(){
	server := grpc.NewServer()
	message.RegisterFriendShipServiceServer(server, new(ImplementFriends))
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}


func RpcFriendCommonStarServer(){

	server := grpc.NewServer()

	message2.RegisterFriendShipCommonStarServiceServer(server, new(ImplementFriendsPlusStars))
	lis, err := net.Listen("tcp", ":9010")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
