package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "gRPC-Streaming/proto"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	logger     = log.New(os.Stdout, "", 0)
	serverAddr = flag.String("server", "", "Server address (host:port)")
	serverHost = flag.String("server-host", "", "Host name to which server IP should resolve")
	insecure   = flag.Bool("insecure", false, "Skip SSL validation? [false]")
	skipVerify = flag.Bool("skip-verify", false, "Skip server hostname verification in SSL validation [false]")
	duration   = flag.Uint("duration", 10, "duration (in seconds) to stream the time from the server for")
)

func init() {
	flag.Parse()
	log.SetFlags(log.Flags() ^ log.Ltime ^ log.Ldate)
}

func main() {
	var opts []grpc.DialOption
	if *serverAddr == "" {
		log.Fatal("-server is empty")
	}
	if *serverHost != "" {
		opts = append(opts, grpc.WithAuthority(*serverHost))
	}
	if *insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		cred := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: *skipVerify,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		logger.Printf("failed to dial server %s: %v", *serverAddr, err)
	}
	defer conn.Close()
	client := pb.NewTimeServiceClient(conn)

	if err := streamTime(client, *duration); err != nil {
		log.Fatal(err)
	}
}

func streamTime(client pb.TimeServiceClient, duration uint) error {
	ctx := context.Background()

	resp, err := client.StreamTime(ctx, &pb.Request{
		DurationSecs: uint32(duration)})
	if err != nil {
		return fmt.Errorf("StreamTime rpc failed: %w", err)
	}
	log.Print("rpc established to timeserver, starting to stream")

	for {
		msg, err := resp.Recv()
		if err == io.EOF {
			log.Printf("end of stream")
			return nil
		} else if err != nil {
			return fmt.Errorf("error receiving message: %w", err)
		}

		ts, err := ptypes.Timestamp(msg.GetCurrentTime())
		if err != nil {
			return fmt.Errorf("failed to parse timestamp %v: %w", msg.GetCurrentTime(), err)
		}
		log.Printf("received message: current_timestamp: %v", ts.Format(time.RFC3339))
	}
}
