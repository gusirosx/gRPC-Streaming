// Sample grpc server with a streaming response.
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	pb "gRPC-Streaming/proto"
)

const responseInterval = time.Second

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("timeserver: starting on port %s", port)
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}

	svc := new(timeService)
	server := grpc.NewServer()
	pb.RegisterTimeServiceServer(server, svc)
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

type timeService struct {
	pb.UnimplementedTimeServiceServer
}

func (timeService) StreamTime(req *pb.Request, resp pb.TimeService_StreamTimeServer) error {
	durationSeconds := req.GetDurationSecs()
	finish := time.Now().Add(time.Second * time.Duration(durationSeconds))

	for time.Now().Before(finish) {
		if err := resp.Send(&pb.TimeResponse{
			CurrentTime: ptypes.TimestampNow()}); err != nil {
			return fmt.Errorf("failed to send message: %w", err)
		}

		select {
		case <-time.After(responseInterval):
		case <-resp.Context().Done():
			log.Printf("response context closed, exiting response")
			return resp.Context().Err()

		}
	}
	return nil
}
