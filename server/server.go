package main

// Sample grpc server with a streaming response.
import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "gRPC-Streaming/proto"
)

const (
	port             = ":50050"
	responseInterval = time.Second
)

// server is used to implement TimeServiceServer.
type timeService struct {
	// Embed the unimplemented server
	pb.UnimplementedTimeServiceServer
}

// StreamTime implements TimeServiceServer
func (timeService) StreamTime(req *pb.Request, resp pb.TimeService_StreamTimeServer) error {
	log.Println("Streaming request received")
	durationSeconds := req.GetDurationSecs()
	finish := time.Now().Add(time.Second * time.Duration(durationSeconds))

	for time.Now().Before(finish) {
		if err := resp.Send(&pb.TimeResponse{
			CurrentTime: timestamppb.Now()}); err != nil {
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

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}
	log.Printf("timeserver: starting at %v", lis.Addr())
	srv := grpc.NewServer()
	pb.RegisterTimeServiceServer(srv, &timeService{})
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
