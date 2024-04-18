package main

import (
	"context"
	"io"
	"log"
	"sync"

	pb "github.com/kaushikkishore/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

func main() {
	// Dial the server
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// Create the gRPC client
	client := pb.NewInvoicerClient(conn)
	in := &pb.CreateLogRequest{
		Name:  "cloudwatch",
		Group: "ec2-describe",
	}
	stream, err := client.Logging(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	// Create a channel to signal when stream processing is complete
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break // Stream is finished
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			log.Printf("Response received: %v", resp)
		}

		// Signal that the stream processing is complete
		close(done)
	}()

	// Wait for the stream processing to complete
	<-done

	// Close the gRPC connection
	if err := conn.Close(); err != nil {
		log.Fatalf("Failed to close the connection: %v", err)
	}

	// Now log "Finished" after ensuring proper synchronization
	log.Printf("Finished")
}
