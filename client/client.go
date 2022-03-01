package main

// Package client is a small tool to query the streaming gRPC endpoint.
import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "gRPC-Streaming/proto"
)

func streamTime(client pb.TimeServiceClient, duration uint) error {
	ctx := context.Background()

	resp, err := client.StreamTime(ctx, &pb.Request{
		DurationSecs: uint32(duration)})
	if err != nil {
		return fmt.Errorf("StreamTime gRPC failed: %w", err)
	}
	log.Print("gRPC established to timeserver, starting to stream")

	for {
		msg, err := resp.Recv()
		if err == io.EOF {
			log.Printf("end of stream")
			return nil
		} else if err != nil {
			return fmt.Errorf("error receiving message: %w", err)
		}
		ts := msg.GetCurrentTime()
		if err := ts.CheckValid(); err != nil {
			log.Println(err.Error())
		}
		log.Printf("Current TimeStamp: %v", ts.AsTime().Format(time.RFC3339))
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := Connection()
	if err != nil {
		log.Printf("failed to dial server %s: %v", *serverAddr, err)
	}
	defer conn.Close()

	client := pb.NewTimeServiceClient(conn)

	if err := streamTime(client, *duration); err != nil {
		log.Fatal(err)
	}
}
