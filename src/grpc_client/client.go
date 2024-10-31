package main

import (
	"context"

	"log"
	"time"

	pb "github.com/shrihariharanba/book-your-train/src/grpc_server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	c2 := pb.NewTrainTicketingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	if err != nil {
		log.Fatalf("error calling function SayHello: %v", err)
	}
	GetAvailableSeats(c2, ctx)
	BookTicket(c2, ctx, 2, "Shri Hari Haran", "B A", "shrihariharanba@live.com")
	GetAvailableSeats(c2, ctx)
	GetTicketPurchaseDetailsByEmail(c2, ctx, "shrihariharanba@live.com")
	GetAvailableSeats(c2, ctx)
	sectionANewSeats := make([]int32, 0, 10)
	sectionANewSeats = append(sectionANewSeats, 5, 2, 1)
	sectionBNewSeats := make([]int32, 0, 10)
	sectionBNewSeats = append(sectionBNewSeats, 8, 9)
	ModifyTicket(c2, ctx, 5, sectionANewSeats, sectionBNewSeats, "Shri Hari Haran", "B A", "shrihariharanba@live.com")
	GetAvailableSeats(c2, ctx)
	DeleteTicketPurchaseDetails(c2, ctx, "shrihariharanba@live.com")
	GetAvailableSeats(c2, ctx)
}

func GetAvailableSeats(client pb.TrainTicketingServiceClient, ctx context.Context) {
	r2, err := client.GetTrainDetails(ctx, &pb.TrainDetails{})
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	log.Println("Welcome to Booking Section")
	log.Println("Train Details:")
	log.Printf("Train Id: %d", r2.TrainId)
	log.Printf("Train Name: %s", r2.Name)
	log.Printf("Journey From: %s", r2.From)
	log.Printf("Journey To: %s", r2.To)
	log.Printf("Cost of each Ticket: %d", r2.Cost)
	log.Println("Available Seats:")
	log.Printf("Section A Available Seat : %v", r2.AvailableSeatSectionA)
	log.Printf("Section B Available Seat : %v", r2.AvailableSeatSectionB)
	log.Printf("Section A Available Seat Count: %v", r2.AvailableSeatSectionACount)
	log.Printf("Section B Available Seat Count: %v", r2.AvailableSeatSecitonBCount)
}

func BookTicket(client pb.TrainTicketingServiceClient, ctx context.Context, noOfTickets int32, firstName string, lastName string, emailId string) {
	r1, err := client.GetTrainDetails(ctx, &pb.TrainDetails{})
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	if noOfTickets > r1.AvailableSeatSectionACount+r1.AvailableSeatSecitonBCount {
		log.Fatal("No of Tickets exceeds the available Tickets")
	}
	totalCostOfTicket := noOfTickets * r1.Cost
	r2, err := client.BookTrainTicket(ctx, &pb.TicketPurchaseDetails{
		TrainId:     r1.TrainId,
		TrainName:   r1.Name,
		From:        r1.From,
		To:          r1.To,
		NoOfTicktes: noOfTickets,
		AmountPaid:  totalCostOfTicket,
		FirstName:   firstName,
		LastName:    lastName,
		EmailId:     emailId,
	})
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	} else {
		log.Printf("Ticket Booked Successfully %v", r2)
	}

}

func ModifyTicket(client pb.TrainTicketingServiceClient, ctx context.Context, noOfTickets int32,
	sectionANewSeat []int32, sectionBNewSeat []int32, firstName string, lastName string, emailId string) {
	r1, err := client.GetTrainDetails(ctx, &pb.TrainDetails{})
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	totalCostOfTicket := noOfTickets * r1.Cost
	r2, err := client.ModifyTrainSeat(ctx, &pb.TicketPurchaseDetails{
		TrainId:      r1.TrainId,
		TrainName:    r1.Name,
		From:         r1.From,
		To:           r1.To,
		NoOfTicktes:  noOfTickets,
		AmountPaid:   totalCostOfTicket,
		FirstName:    firstName,
		LastName:     lastName,
		EmailId:      emailId,
		SectionASeat: sectionANewSeat,
		SectionBSeat: sectionBNewSeat,
	})
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	} else {
		log.Printf("Ticket Modifed Successfully %v", r2)
	}

}

func GetTicketPurchaseDetailsByEmail(client pb.TrainTicketingServiceClient, ctx context.Context, emailId string) error {
	ticketPurchaseDetails, _ := client.GetTicketPurchaseDetails(ctx, &pb.TicketPurchaseDetails{EmailId: emailId})
	log.Printf("Ticket Purchase Details: %v", ticketPurchaseDetails)
	return nil
}

func DeleteTicketPurchaseDetails(client pb.TrainTicketingServiceClient, ctx context.Context, emailId string) error {
	ticketPurchaseDetails, _ := client.DeleteTicketPurchaseDetails(ctx, &pb.TicketPurchaseDetails{EmailId: emailId})
	log.Printf("User Deleted from the train: %v", ticketPurchaseDetails)
	return nil
}
