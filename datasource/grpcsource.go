package datasource

import (
	"github.com/birjemin/iris-structure/conf"
	"google.golang.org/grpc"
)

var grpcConn *grpc.ClientConn

func GetGRPC() *grpc.ClientConn {
	return grpcConn
}

// 关闭db
func CloseGRPC() error {
	if grpcConn != nil {
		return grpcConn.Close()
	} else {
		return nil
	}
}

func init() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(conf.Sysconfig.GRPCServer, grpc.WithInsecure())
	if err != nil {
		panic("grpc connect failed.")
	}
	grpcConn = conn
}
