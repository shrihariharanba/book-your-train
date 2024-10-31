package grpc

import (
	"fmt"
	"log"
	"net"

	pb "github.com/shrihariharanba/book-your-train/src/grpc_server/pb"
	srvc "github.com/shrihariharanba/book-your-train/src/grpc_server/service"
	grpc_server "google.golang.org/grpc"
)

func StartServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", 8081))

	if err != nil {
		log.Fatalf("Unable to listen to the port:%d", 8081)
	}

	gs := grpc_server.NewServer()

	pb.RegisterTrainTicketingServiceServer(gs, &srvc.TrainTicketingService{})

	log.Printf("gRPC server listening at %v", lis.Addr())

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Print("Server started successfully")
}
