package main

import (
	"fmt"
	pb "grpc/proto"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func generateID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewWishListServiceClient(conn)

	res, err := serviceClient.Create(context.Background(), &pb.CreateWishListReq{
		WishList: &pb.WishList{
			Id:   generateID(),
			Name: "my wishlist",
		},
	})

	if err != nil {
		panic("wishlist is not created  " + err.Error())
	}

	fmt.Println(res.WishListId)

}
