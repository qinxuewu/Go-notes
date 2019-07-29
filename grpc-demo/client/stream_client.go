package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "grpc-demo/proto"
)

func main()  {
	conn,err:=grpc.Dial(":8899",grpc.WithInsecure())
	if err !=nil{
		log.Fatal("grpc.Dial err : %v",err)
	}

	defer  conn.Close()

	client:=pb.NewStreamServiceClient(conn)

	// 发送普通请求 接收服务端返回的流式RPC 数据
	err = printLists(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "客户端发送普通请求", Value: 2018}})
	if err != nil {
		log.Fatalf("printLists.err: %v", err)
	}

	// 客户端发送流式请求 接收服务端返回的普通数据
	err = printRecord(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Record", Value: 2018}})
	if err != nil {
		log.Fatalf("printRecord.err: %v", err)
	}

	// 双向流式数据通信
	err = printRoute(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Route", Value: 2018}})
	if err != nil {
		log.Fatalf("printRoute.err: %v", err)
	}
}


func printLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream,err:=client.List(context.Background(),r)
	if err !=nil {
		return  err
	}
	for  {
		resp,err:=stream.Recv()
		if err == io.EOF {
			break
		}
		if err !=nil{
			return  err
		}
		log.Printf("接收服务端返回的流式数据: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}
	return nil
}
// 客户端发送流式rpc 接收服务端普通返回
func printRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}
	for n := 0; n < 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("接收服务端返回的普通rpc数据: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	return  nil
}

func printRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	return nil
}