package service

import (
	"context"
	"log"

	"github.com/shrihariharanba/book-your-train/src/grpc_server/pb"
	"github.com/shrihariharanba/book-your-train/src/source"
)

type TrainTicketingService struct {
	pb.UnimplementedTrainTicketingServiceServer
}

func (*TrainTicketingService) GetTrainDetails(context.Context, *pb.TrainDetails) (*pb.TrainDetails, error) {
	trainInfo := source.GettrainDetails()
	availSeatSectionA, availSeatsSectionB := source.GetAvailableSeats()
	return &pb.TrainDetails{TrainId: int32(trainInfo.TrainId),
		Name:                       trainInfo.Name,
		From:                       trainInfo.From,
		To:                         trainInfo.To,
		Cost:                       trainInfo.Cost,
		AvailableSeatSectionA:      availSeatSectionA,
		AvailableSeatSectionB:      availSeatsSectionB,
		AvailableSeatSectionACount: int32(len(availSeatSectionA)),
		AvailableSeatSecitonBCount: int32(len(availSeatsSectionB))}, nil
}

func (*TrainTicketingService) BookTrainTicket(ctx context.Context, in *pb.TicketPurchaseDetails) (*pb.TicketPurchaseDetails, error) {
	return source.BookTicket(in)
}

func (*TrainTicketingService) ModifyTrainSeat(ctx context.Context, in *pb.TicketPurchaseDetails) (*pb.TicketPurchaseDetails, error) {
	log.Printf("Modifying the seats for the user, %s", in.EmailId)
	return source.ModifySeats(in)
}

func (*TrainTicketingService) GetTicketPurchaseDetails(ctx context.Context, in *pb.TicketPurchaseDetails) (*pb.TicketPurchaseDetails, error) {
	return source.GetTicketPurchaseDetails(in.EmailId)
}

func (*TrainTicketingService) DeleteTicketPurchaseDetails(ctx context.Context, in *pb.TicketPurchaseDetails) (*pb.TicketPurchaseDetails, error) {
	log.Printf("Deleting the user %s", in.EmailId)
	return source.DeleteUserFromTrain(in.EmailId)
}
