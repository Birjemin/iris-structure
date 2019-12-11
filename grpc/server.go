package main

import (
	"flag"
	pb "github.com/birjemin/iris-structure/grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// server is used to implement hello.GreeterServer.
type server struct{}

// SayHello implements Hello.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	print("iris-structure")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	var port = flag.String("port", "50051", "port flag")
	flag.Parse()

	lis, err := net.Listen("tcp", ":" + *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
