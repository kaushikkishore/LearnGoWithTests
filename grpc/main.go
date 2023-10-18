package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/kaushikkishore/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	// log.Println(ctx)
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

// Create a LogMessage function.
// The function will be doing the streaming for the logs in a infinite loop.
// It will send a log message every 1 second.
// Learn how can i debug this and make it work.
func (s myInvoicerServer) Logging(req *invoicer.CreateLogRequest, stream invoicer.Invoicer_LoggingServer) error {
	log.Printf("Received: %v", req)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(5)
		go func(count int64) {
			defer wg.Done()

			// Sleep to simulate server working process.
			time.Sleep(time.Duration(count) * time.Second)

			resp := invoicer.CreateLogResponse{
				LogMessage: "test " + strconv.Itoa(int(count)),
				Timestamp:  time.Now().String(),
				Type:       "debug",
			}
			if err := stream.Send(&resp); err != nil {
				log.Fatalf("Failed to sream %v", err)

			}

			log.Printf("Sent: %d", count)
		}(int64(i))
	}

	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("Cannot create listner %s", err)
	}

	serverRegistrat := grpc.NewServer()

	srvc := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegistrat, srvc)

	err = serverRegistrat.Serve(lis)

	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
