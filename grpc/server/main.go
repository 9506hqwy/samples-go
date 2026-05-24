package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"grpc_server/pkg/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	httpPort  = ":5000"
	httpsPort = ":5001"
)

type server struct {
	helloworld.UnimplementedApiServer
}

//revive:disable:unused-receiver

func (s *server) Call(
	_ context.Context,
	req *helloworld.ScalarValueTypes,
) (*helloworld.ScalarValueTypes, error) {
	res, ok := proto.Clone(req).(*helloworld.ScalarValueTypes)
	if ok {
		return res, nil
	}

	return nil, status.Errorf(codes.Internal, "failed to copy.")
}

func (s *server) Download(
	_ *emptypb.Empty,
	stream grpc.ServerStreamingServer[helloworld.Message],
) error {
	var count uint
	for {
		time.Sleep(1 * time.Second)

		err := stream.Send(&helloworld.Message{
			Message: fmt.Sprintf("count %d", count),
		})
		if err != nil {
			break
		}

		count++
	}

	return nil
}

func (s *server) Upload(
	stream grpc.ClientStreamingServer[helloworld.Message, emptypb.Empty],
) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}

		_, _ = fmt.Printf("Message: %v\n", msg)
	}

	stream.SendAndClose(&emptypb.Empty{})

	return nil
}

func (s *server) Async(
	stream grpc.BidiStreamingServer[helloworld.Message, helloworld.Message],
) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}

		err = stream.Send(msg)
		if err != nil {
			break
		}
	}

	return nil
}

//revive:enable:unused-receiver

func main() {
	httpListener, err := net.Listen("tcp", httpPort)
	if err != nil {
		_, _ = fmt.Printf("Could not listen(http): %v\n", err)
		return
	}
	defer httpListener.Close()

	httpServer := grpc.NewServer()
	helloworld.RegisterApiServer(httpServer, &server{})

	httpsListener, err := net.Listen("tcp", httpsPort)
	if err != nil {
		_, _ = fmt.Printf("Could not listen(https): %v\n", err)
		return
	}
	defer httpsListener.Close()

	creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	if err != nil {
		_, _ = fmt.Printf("Coud not load certificate: %v\n", err)
		return
	}

	serverOpts := []grpc.ServerOption{grpc.Creds(creds)}
	httpsServer := grpc.NewServer(serverOpts...)
	helloworld.RegisterApiServer(httpsServer, &server{})

	go httpServer.Serve(httpListener)
	httpsErr := httpsServer.Serve(httpsListener)
	if httpsErr != nil {
		_, _ = fmt.Printf("Could not serve(https): %v\n", httpsErr)
		return
	}
}
