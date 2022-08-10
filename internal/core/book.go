package core

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/Strike-official/hemantMensSalonBot/internal/model"
	"github.com/strike-official/go-sdk/strike"
)

func BookAppointment(req model.Request_Structure) *strike.Response_structure {
	strikeObject := strike.Create("slotSelection", model.Conf.APIEp+"slotSelection")

	question_object := strikeObject.Question("selectedServices").
		QuestionCard().SetHeaderToQuestion(1, strike.FULL_WIDTH).AddTextRowToQuestion(strike.H3, "Hi "+req.Bybrisk_session_variables.Username+", What services would you like to opt? (long press to select multiple services)", "#2B4865", false)

	// Add answer
	answer_object := question_object.Answer(true).AnswerCardArray(strike.VERTICAL_ORIENTATION)

	keys := make([]string, 0)
	for k, _ := range model.Services {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		answer_object = answer_object.AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, k, "#256D85", true)
	}

	return strikeObject
}

func SlotSelection(req model.Request_Structure) *strike.Response_structure {

	selectedServices := req.User_session_variables.SelectedServices
	var totalTime int
	for _, s := range selectedServices {
		totalTime += model.Services[s]
		fmt.Println(s)
	}
	fmt.Println("Total Time -> ", totalTime)

	// QueueingLogic

	// 1. When a new booking comes in:
	// 	- Fetch Current Date and Time
	// 	- From DB, get all the bookings for CurDate.

	// - Calculate Possible Slot based on the Time.
	// - Find Next Best Slot for each Possible Slot.
	// - Send it to Customer, and based on the selection.
	// - Add it to the DB as a new entry.

	curTime := time.Now()

	currentTimeStr := fmt.Sprintf("%d%02d%02d", curTime.Hour(), curTime.Minute(), curTime.Second())
	closingTimeStr := fmt.Sprintf("%d%02d%02d", model.SalonClosingTime.Hour(), model.SalonClosingTime.Minute(), model.SalonClosingTime.Second())
	fmt.Println(currentTimeStr, "--", closingTimeStr)

	// model.SlotDetail
	var slot int
	var slots *[]model.SlotDetails
	if currentTimeStr < closingTimeStr {
		fmt.Println("Salon is Open!")
		if curTime.Hour() < 12 {
			slot = 1
		} else if curTime.Hour() < 15 {
			slot = 2
			fmt.Println(slot)
		} else if curTime.Hour() < 20 {
			slot = 3
			slots = getSlots(totalTime, slot)
		}
		fmt.Println(slot, slots)
	}

	timein := curTime.Local().Add(time.Minute * time.Duration(totalTime))
	fmt.Println("Start -> ", curTime)
	fmt.Println("EndTime ->", timein)

	// ---------------------------------
	// Create StrikeOBJ
	// ---------------------------------

	strikeObject := strike.Create("paymentAmount", model.Conf.APIEp+"paymentLink")

	question_object := strikeObject.Question("paymentAmount").
		QuestionCard().SetHeaderToQuestion(1, strike.FULL_WIDTH).AddTextRowToQuestion(strike.H3, "Hi "+req.Bybrisk_session_variables.Username+", Please choose any of the slot, to book your appointment.", "#2B4865", false)

	// Add answer
	answer_object := question_object.Answer(true).AnswerCardArray(strike.VERTICAL_ORIENTATION)

	for _, s := range *slots {
		answer_object = answer_object.AnswerCard().SetHeaderToAnswer(4, strike.HALF_WIDTH).
			AddTextRowToAnswer(strike.H2, "Slot: "+strconv.Itoa(s.Slot), "#86B049", true).
			AddTextRowToAnswer(strike.H3, "Wait Time: "+strconv.Itoa(s.WaitTime)+" Minutes", "black", false).
			AddTextRowToAnswer(strike.H3, "Come At: "+s.ComeAt.Format("2006-01-02 03:04:05 PM"), "black", false).
			AddTextRowToAnswer(strike.H3, "Payment Amount: Rs"+strconv.Itoa(s.Cost), "#212121", false)
	}
	return strikeObject
}

func PaymentPortal(req model.Request_Structure) *strike.Response_structure {

	paymentDetails := req.User_session_variables.PaymentDetails

	fmt.Println("Payment Amount: ", paymentDetails[0])
	return nil

}

// return
// [
// 	{
// 		"Slot":2,
// 		"Chair":1,
// 		"WaitTime":40,
// 		"ComeAt":"2022-08-10 19:15:00"
// 	},
// 	{
// 		"Slot":3,
// 		"Chair":2,
// 		"WaitTime":160,
// 		"ComeAt":"2022-08-10 21:15:00"
// 	},

// ]

func getSlots(totalTime, slot int) *[]model.SlotDetails {

	cost := totalTime * 4
	curTime := time.Now()
	sample1 := &model.SlotDetails{
		Slot:     2,
		Chair:    1,
		ComeAt:   curTime.Local().Add(time.Minute * time.Duration(totalTime)),
		WaitTime: 40,
		Cost:     cost,
	}
	sample2 := &model.SlotDetails{
		Slot:     3,
		Chair:    2,
		ComeAt:   curTime.Local().Add(time.Minute * time.Duration(totalTime+75)),
		WaitTime: 120,
		Cost:     cost,
	}

	var sm []model.SlotDetails
	sm = append(sm, *sample1)
	sm = append(sm, *sample2)

	return &sm
}
