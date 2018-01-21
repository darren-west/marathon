package main

import (
	"log"
	"net"

	"github.com/darren-west/marathon/tracker"
	"github.com/darren-west/marathon/tracker/server/handler"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	tracker.RegisterLocationServiceServer(grpcServer, &handler.LocationHandler{})

	log.Fatal(grpcServer.Serve(lis))
}

//func GetUnaryServerInterceptor() grpc.UnaryServerInterceptor {
//	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//		methodName := fmt.Sprintf("server:%s", info.FullMethod)
//		log.Print(methodName)
//
//		md, _ := metadata.FromIncomingContext(ctx)
//		log.Print(md)
//		return handler(ctx, req)
//	}
//}
