package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"grpc_client/pkg/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
)

func call(ctx context.Context, c helloworld.ApiClient) error {
	res, err := c.Call(ctx, &helloworld.ScalarValueTypes{})
	if err != nil {
		_, _ = fmt.Printf("Failed to call: %v\n", err)
		return err
	}

	_, _ = fmt.Printf("%#v\n", res)
	return nil
}

func download(ctx context.Context, c helloworld.ApiClient) error {
	downStream, err := c.Download(ctx, &emptypb.Empty{})
	if err != nil {
		_, _ = fmt.Printf("Failed to download: %v\n", err)
		return err
	}

	const count int = 3
	for range count {
		msg, err := downStream.Recv()
		if err != nil {
			_, _ = fmt.Printf("Failed to download message: %v\n", err)
			return err
		}

		_, _ = fmt.Printf("Message: %v\n", msg.Message)
	}

	return nil
}

func upload(ctx context.Context, c helloworld.ApiClient) error {
	upStream, err := c.Upload(ctx)
	if err != nil {
		_, _ = fmt.Printf("Failed to upload: %v\n", err)
		return err
	}

	const count int = 3
	for i := range count {
		err = upStream.Send(&helloworld.Message{
			Message: fmt.Sprintf("count %d", i),
		})
		if err != nil {
			_, _ = fmt.Printf("Failed to upload message: %v\n", err)
			return err
		}
	}
	_, err = upStream.CloseAndRecv()
	if err != nil {
		_, _ = fmt.Printf("Failed to upload message: %v\n", err)
		return err
	}

	return nil
}

func async(ctx context.Context, c helloworld.ApiClient) error {
	biStream, err := c.Async(ctx)
	if err != nil {
		_, _ = fmt.Printf("Failed to async: %v\n", err)
		return err
	}

	const count int = 3
	for i := range count {
		err = biStream.Send(&helloworld.Message{
			Message: fmt.Sprintf("count %d", i),
		})
		if err != nil {
			_, _ = fmt.Printf("Failed to async message: %v\n", err)
			return err
		}

		msg, err := biStream.Recv()
		if err != nil {
			_, _ = fmt.Printf("Failed to async message: %v\n", err)
			return err
		}

		_, _ = fmt.Printf("Message: %v\n", msg.Message)
	}

	return nil
}

func main() {
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	conn, err := grpc.NewClient(
		"localhost:5001",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		_, _ = fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	c := helloworld.NewApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Call
	err = call(ctx, c)
	if err != nil {
		return
	}

	// Download
	err = download(ctx, c)
	if err != nil {
		return
	}

	// Upload
	err = upload(ctx, c)
	if err != nil {
		return
	}

	// Async
	err = async(ctx, c)
	if err != nil {
		return
	}
}
