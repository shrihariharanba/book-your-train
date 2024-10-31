package source

import (
	"testing"

	"github.com/shrihariharanba/book-your-train/src/grpc_server/pb"
	"github.com/stretchr/testify/assert"
)

func TestGettrainDetails(t *testing.T) {
	train := GettrainDetails()
	assert.Equal(t, int32(12345), train.TrainId)
	assert.Equal(t, "London Express", train.Name)
	assert.Equal(t, "London", train.From)
	assert.Equal(t, "France", train.To)
	assert.Equal(t, int32(20), train.Cost)
}

func TestGetAvailableSeats(t *testing.T) {
	sectionASeats, sectionBSeats := GetAvailableSeats()
	assert.Equal(t, 10, len(sectionASeats), "Expected 10 seats available in Section A initially")
	assert.Equal(t, 10, len(sectionBSeats), "Expected 10 seats available in Section B initially")
}

func TestBookTicket_SuccessfulBooking(t *testing.T) {
	ticketDetails := &pb.TicketPurchaseDetails{
		TrainId:     12345,
		NoOfTicktes: 5,
		FirstName:   "John",
		LastName:    "Doe",
		EmailId:     "johndoe@example.com",
	}

	response, err := BookTicket(ticketDetails)
	assert.NoError(t, err)
	assert.Equal(t, ticketDetails.EmailId, response.EmailId, "The email in the response should match the booking request")
}

func TestBookTicket_InsufficientSeats(t *testing.T) {
	ticketDetails := &pb.TicketPurchaseDetails{
		TrainId:     12345,
		NoOfTicktes: 25, // Trying to book more than available seats
		FirstName:   "John",
		LastName:    "Doe",
		EmailId:     "johndoe@example.com",
	}

	response, err := BookTicket(ticketDetails)
	assert.Error(t, err, "Expected an error due to insufficient seats")
	assert.Nil(t, response, "Response should be nil when booking fails")
}

func TestModifySeats(t *testing.T) {
	ticketDetails := &pb.TicketPurchaseDetails{
		TrainId:      12345,
		FirstName:    "John",
		LastName:     "Doe",
		EmailId:      "johndoe@example.com",
		SectionASeat: []int32{1, 2},
		SectionBSeat: []int32{3},
	}

	response, err := ModifySeats(ticketDetails)
	assert.NoError(t, err)
	assert.Equal(t, ticketDetails.EmailId, response.EmailId, "The email in the response should match the modification request")
}

func TestDeleteUserFromTrain(t *testing.T) {
	email := "johndoe@example.com"

	// Book a ticket to ensure the user has data to delete
	ticketDetails := &pb.TicketPurchaseDetails{
		TrainId:     12345,
		NoOfTicktes: 2,
		FirstName:   "John",
		LastName:    "Doe",
		EmailId:     email,
	}
	_, _ = BookTicket(ticketDetails)

	response, err := DeleteUserFromTrain(email)
	assert.NoError(t, err, "Expected no error when deleting a user")
	assert.Equal(t, email, response.EmailId, "The email in the response should match the deleted user email")
}

func TestGetTicketPurchaseDetails(t *testing.T) {
	email := "janedoe@example.com"

	// Book a ticket to create data for retrieval
	ticketDetails := &pb.TicketPurchaseDetails{
		TrainId:     12345,
		NoOfTicktes: 2,
		FirstName:   "Jane",
		LastName:    "Doe",
		EmailId:     email,
	}
	_, _ = BookTicket(ticketDetails)

	response, err := GetTicketPurchaseDetails(email)
	assert.NoError(t, err, "Expected no error when retrieving ticket purchase details")
	assert.Equal(t, email, response.EmailId, "The email in the response should match the requested email")
	assert.Equal(t, int32(12345), response.TrainId, "The train ID in the response should match the booking train ID")
}
