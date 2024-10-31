package source

import (
	"errors"
	"log"
	"math/rand/v2"

	"github.com/shrihariharanba/book-your-train/src/grpc_server/pb"
)

var trainObject *Train

type Train struct {
	TrainId  int32
	Name     string
	From     string
	To       string
	Cost     int32
	SectionA *SecionA
	SectionB *SectionB
}

type SecionA struct {
	SeatDetails map[int32]*TicketDetails
}

type SectionB struct {
	SeatDetails map[int32]*TicketDetails
}

type TicketDetails struct {
	TicketId int
	TrainId  int
	User     *User
	Cost     int
	SeatNo   int32
}

type User struct {
	FirstName string
	LastName  string
	EmailId   string
}

func init() {
	trainObject = new(Train)
	trainObject.Name = "London Express"
	trainObject.From = "London"
	trainObject.To = "France"
	trainObject.Cost = 20
	trainObject.TrainId = 12345
	trainObject.SectionA = new(SecionA)
	trainObject.SectionB = new(SectionB)
	trainObject.SectionA.SeatDetails = refreshSeatLoad()
	trainObject.SectionB.SeatDetails = refreshSeatLoad()
}

func refreshSeatLoad() map[int32]*TicketDetails {
	seatDetails := make(map[int32]*TicketDetails, 10)
	for i := 0; i < 10; i++ {
		seatDetails[int32(i)] = nil
	}
	return seatDetails
}

func GettrainDetails() *Train {
	return trainObject
}

func GetAvailableSeats() ([]int32, []int32) {
	sectionAAvailableSeats := make([]int32, 0, 10)
	sectionBAvailableSeats := make([]int32, 0, 10)
	for k, v := range trainObject.SectionA.SeatDetails {
		if v == nil {
			sectionAAvailableSeats = append(sectionAAvailableSeats, k)
		}
	}

	for k, v := range trainObject.SectionB.SeatDetails {
		if v == nil {
			sectionBAvailableSeats = append(sectionBAvailableSeats, k)
		}
	}

	return sectionAAvailableSeats, sectionBAvailableSeats
}

func BookTicket(ticketPurchaseDetails *pb.TicketPurchaseDetails) (*pb.TicketPurchaseDetails, error) {
	totalSeatNeeded := ticketPurchaseDetails.NoOfTicktes
	sectionAallocatedSeat := SeatAllocationSectionA(ticketPurchaseDetails)
	log.Printf("Total no seat allocated in Section A %d", sectionAallocatedSeat)
	remainingSeatNeeded := int(totalSeatNeeded) - sectionAallocatedSeat
	if remainingSeatNeeded == 0 {
		log.Print("Successfully Booking Completed")
		return ticketPurchaseDetails, nil
	} else if remainingSeatNeeded > 0 {
		sectionBAllocatedSeat := SeatAllocationSectionB(remainingSeatNeeded, ticketPurchaseDetails)
		log.Printf("Total no seat allocated in Section B %d", sectionBAllocatedSeat)
		totalSeatAllocated := sectionAallocatedSeat + sectionBAllocatedSeat
		if int(totalSeatNeeded) == totalSeatAllocated {
			log.Print("Successfully Booking Completed")
			return ticketPurchaseDetails, nil
		} else {
			return nil, errors.New("Seats no filled completely")
		}
	}
	return nil, errors.New("Seats no filled completely")
}

func ModifySeats(ticketPurchaseDetails *pb.TicketPurchaseDetails) (*pb.TicketPurchaseDetails, error) {
	_, _ = DeleteUserFromTrain(ticketPurchaseDetails.EmailId)
	return DirectSeatAllocation(ticketPurchaseDetails, ticketPurchaseDetails.SectionASeat, ticketPurchaseDetails.SectionBSeat)
}
func DirectSeatAllocation(ticketPurchaseDetails *pb.TicketPurchaseDetails, sectionAnewSeat []int32, sectionBnewSeat []int32) (*pb.TicketPurchaseDetails, error) {
	sectionASeatdetails := trainObject.SectionA.SeatDetails
	sectionBSeatdetails := trainObject.SectionB.SeatDetails
	for _, i := range sectionAnewSeat {
		user := &User{
			FirstName: ticketPurchaseDetails.FirstName,
			LastName:  ticketPurchaseDetails.LastName,
			EmailId:   ticketPurchaseDetails.EmailId,
		}
		ticketDetails := &TicketDetails{
			TicketId: rand.Int(),
			TrainId:  int(ticketPurchaseDetails.TrainId),
			SeatNo:   i,
			Cost:     20,
			User:     user,
		}
		sectionASeatdetails[i] = ticketDetails
	}

	for _, j := range sectionBnewSeat {
		user := &User{
			FirstName: ticketPurchaseDetails.FirstName,
			LastName:  ticketPurchaseDetails.LastName,
			EmailId:   ticketPurchaseDetails.EmailId,
		}
		ticketDetails := &TicketDetails{
			TicketId: rand.Int(),
			TrainId:  int(ticketPurchaseDetails.TrainId),
			SeatNo:   j,
			Cost:     20,
			User:     user,
		}
		sectionBSeatdetails[j] = ticketDetails
	}
	trainObject.SectionA.SeatDetails = sectionASeatdetails
	trainObject.SectionB.SeatDetails = sectionBSeatdetails
	return ticketPurchaseDetails, nil
}
func SeatAllocationSectionA(ticketPurchaseDetails *pb.TicketPurchaseDetails) (totalSeatAllocated int) {
	sectioASeatDetails := trainObject.SectionA.SeatDetails
	noOfSeatNeeded := ticketPurchaseDetails.NoOfTicktes
	var totalNoOfSeatAllocated int = 0
	for k, v := range sectioASeatDetails {
		if v == nil && int(noOfSeatNeeded) != totalNoOfSeatAllocated {
			user := &User{
				FirstName: ticketPurchaseDetails.FirstName,
				LastName:  ticketPurchaseDetails.LastName,
				EmailId:   ticketPurchaseDetails.EmailId,
			}
			ticketDetails := &TicketDetails{
				TicketId: rand.Int(),
				TrainId:  int(ticketPurchaseDetails.TrainId),
				SeatNo:   k,
				Cost:     20,
				User:     user,
			}
			sectioASeatDetails[k] = ticketDetails
			totalNoOfSeatAllocated++
		}
	}
	trainObject.SectionA.SeatDetails = sectioASeatDetails
	return totalNoOfSeatAllocated
}

func SeatAllocationSectionB(remainingSeatsNeeded int, ticketPurchaseDetails *pb.TicketPurchaseDetails) (totalSeatAllocated int) {
	sectionBSeatDetails := trainObject.SectionB.SeatDetails
	var totalNoOfSeatAllocated int = 0
	for k, v := range sectionBSeatDetails {
		if v == nil && int(remainingSeatsNeeded) != totalNoOfSeatAllocated {
			user := &User{
				FirstName: ticketPurchaseDetails.FirstName,
				LastName:  ticketPurchaseDetails.LastName,
				EmailId:   ticketPurchaseDetails.EmailId,
			}
			ticketDetails := &TicketDetails{
				TicketId: rand.Int(),
				TrainId:  int(ticketPurchaseDetails.TrainId),
				SeatNo:   k,
				Cost:     20,
				User:     user,
			}
			sectionBSeatDetails[k] = ticketDetails
			totalNoOfSeatAllocated++
		}
	}
	trainObject.SectionB.SeatDetails = sectionBSeatDetails
	return totalNoOfSeatAllocated
}

func GetTicketPurchaseDetails(emaiId string) (*pb.TicketPurchaseDetails, error) {
	sectionASeatDetails := trainObject.SectionA.SeatDetails
	sectionBSeatDetails := trainObject.SectionB.SeatDetails
	tpd := &pb.TicketPurchaseDetails{}
	tpd.TrainId = trainObject.TrainId
	tpd.TrainName = trainObject.Name
	tpd.From = trainObject.From
	tpd.To = trainObject.To
	tpd.EmailId = emaiId
	return GetSectionSeatNoByEmailId(sectionASeatDetails, sectionBSeatDetails, emaiId, tpd)
}

func GetSectionSeatNoByEmailId(sectionASeatDetails map[int32]*TicketDetails, sectionBSeatDetails map[int32]*TicketDetails,
	emailId string, tpd *pb.TicketPurchaseDetails) (*pb.TicketPurchaseDetails, error) {
	sectionASeatNo := make([]int32, 0, 10)
	sectionBSeatNo := make([]int32, 0, 10)
	log.Printf("Size of the map: %d", len(sectionASeatDetails))
	for _, v := range sectionASeatDetails {
		if v != nil && v.User.EmailId == emailId {
			sectionASeatNo = append(sectionASeatNo, v.SeatNo)
		}
	}

	for _, v := range sectionBSeatDetails {
		if v != nil && v.User.EmailId == emailId {
			sectionBSeatNo = append(sectionBSeatNo, v.SeatNo)
		}
	}

	tpd.NoOfTicktes = (int32(len(sectionASeatNo)) + int32(len(sectionBSeatNo)))
	tpd.AmountPaid = (int32(len(sectionASeatNo)) + int32(len(sectionBSeatNo))) * 20
	tpd.SectionASeat = sectionASeatNo
	tpd.SectionBSeat = sectionBSeatNo
	return tpd, nil
}

func DeleteUserFromTrain(emailId string) (*pb.TicketPurchaseDetails, error) {
	sectionASeatDetails := trainObject.SectionA.SeatDetails
	sectionBSeatDetails := trainObject.SectionB.SeatDetails
	for k, v := range sectionASeatDetails {
		if v != nil && v.User.EmailId == emailId {
			sectionASeatDetails[k] = nil
		}
	}

	for k, v := range sectionBSeatDetails {
		if v != nil && v.User.EmailId == emailId {
			sectionBSeatDetails[k] = nil
		}
	}

	trainObject.SectionA.SeatDetails = sectionASeatDetails
	trainObject.SectionB.SeatDetails = sectionBSeatDetails
	log.Printf("Deleted USer mail: %s", emailId)
	return &pb.TicketPurchaseDetails{EmailId: emailId}, nil
}
