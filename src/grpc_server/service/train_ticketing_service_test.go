package service

import (
	"context"
	"testing"

	"github.com/shrihariharanba/book-your-train/src/grpc_server/pb"
	"github.com/shrihariharanba/book-your-train/src/source"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the source package functions to isolate service tests.
type mockSource struct {
	mock.Mock
}

func (m *mockSource) GettrainDetails() *source.Train {
	return &source.Train{
		TrainId: 12345,
		Name:    "London Express",
		From:    "London",
		To:      "France",
		Cost:    20,
		SectionA: &source.SecionA{
			SeatDetails: map[int32]*source.TicketDetails{},
		},
		SectionB: &source.SectionB{
			SeatDetails: map[int32]*source.TicketDetails{},
		},
	}
}

func (m *mockSource) GetAvailableSeats() ([]int32, []int32) {
	return []int32{1, 2, 3}, []int32{4, 5, 6}
}

func (m *mockSource) BookTicket(details *pb.TicketPurchaseDetails) (*pb.TicketPurchaseDetails, error) {
	return details, nil
}

func (m *mockSource) ModifySeats(details *pb.TicketPurchaseDetails) (*pb.TicketPurchaseDetails, error) {
	return details, nil
}

func (m *mockSource) GetTicketPurchaseDetails(email string) (*pb.TicketPurchaseDetails, error) {
	return &pb.TicketPurchaseDetails{
		TrainId:    12345,
		EmailId:    email,
		From:       "London",
		To:         "France",
		AmountPaid: 20,
	}, nil
}

func (m *mockSource) DeleteUserFromTrain(email string) (*pb.TicketPurchaseDetails, error) {
	return &pb.TicketPurchaseDetails{EmailId: email}, nil
}

func TestGetTrainDetails(t *testing.T) {
	service := &TrainTicketingService{}
	mockSrc := &mockSource{}
	mockSrc.On("GetTrainDetails").Return(mockSrc.GettrainDetails())
	mockSrc.On("GetAvailableSeats").Return(mockSrc.GetAvailableSeats())

	resp, err := service.GetTrainDetails(context.Background(), &pb.TrainDetails{})
	assert.NoError(t, err)
	assert.Equal(t, int32(12345), resp.TrainId)
	assert.Equal(t, "London Express", resp.Name)
	assert.Equal(t, int32(3), resp.AvailableSeatSectionACount)
	assert.Equal(t, int32(3), resp.AvailableSeatSecitonBCount)
}

func TestBookTrainTicket(t *testing.T) {
	service := &TrainTicketingService{}
	mockSrc := &mockSource{}
	ticketDetails := &pb.TicketPurchaseDetails{
		TrainId:     12345,
		FirstName:   "John",
		LastName:    "Doe",
		EmailId:     "johndoe@example.com",
		NoOfTicktes: 2,
	}
	mockSrc.On("BookTicket").Return(mockSrc.BookTicket(ticketDetails))

	resp, err := service.BookTrainTicket(context.Background(), ticketDetails)
	assert.NoError(t, err)
	assert.Equal(t, ticketDetails.EmailId, resp.EmailId)
}

func TestModifyTrainSeat(t *testing.T) {
	service := &TrainTicketingService{}
	mockSrc := &mockSource{}

	ticketDetails := &pb.TicketPurchaseDetails{
		TrainId:      12345,
		FirstName:    "Jane",
		LastName:     "Smith",
		EmailId:      "janesmith@example.com",
		SectionASeat: []int32{1, 2},
		SectionBSeat: []int32{3},
	}
	mockSrc.On("ModifySeats").Return(mockSrc.ModifySeats)

	resp, err := service.ModifyTrainSeat(context.Background(), ticketDetails)
	assert.NoError(t, err)
	assert.Equal(t, ticketDetails.EmailId, resp.EmailId)
}

func TestGetTicketPurchaseDetails(t *testing.T) {
	service := &TrainTicketingService{}
	mockSrc := &mockSource{}

	ticketDetails := &pb.TicketPurchaseDetails{
		EmailId: "johndoe@example.com",
	}
	mockSrc.On("GetTicketPurchaseDetails").Return(mockSrc.GetTicketPurchaseDetails("johndoe@example.com"))

	resp, err := service.GetTicketPurchaseDetails(context.Background(), ticketDetails)
	assert.NoError(t, err)
	assert.Equal(t, ticketDetails.EmailId, resp.EmailId)
	assert.Equal(t, int32(12345), resp.TrainId)
	assert.Equal(t, int32(20), resp.AmountPaid)
}

func TestDeleteTicketPurchaseDetails(t *testing.T) {
	service := &TrainTicketingService{}
	mockSrc := &mockSource{}
	mockSrc.On("DeleteUserFromTrain").Return(mockSrc.DeleteUserFromTrain)

	ticketDetails := &pb.TicketPurchaseDetails{
		EmailId: "johndoe@example.com",
	}

	resp, err := service.DeleteTicketPurchaseDetails(context.Background(), ticketDetails)
	assert.NoError(t, err)
	assert.Equal(t, ticketDetails.EmailId, resp.EmailId)
}
