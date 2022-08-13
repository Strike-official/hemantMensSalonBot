package core

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Strike-official/hemantMensSalonBot/internal/model"
	"github.com/Strike-official/hemantMensSalonBot/pkg/payment"
	"github.com/strike-official/go-sdk/strike"
)

func BookAppointment(req model.Request_Structure) *strike.Response_structure {
	strikeObject := strike.Create("slotSelection", model.Conf.APIEp+"slotSelection")

	log.Println("URL: ", model.Conf.APIEp+"slotSelection")
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
		// if curTime.Hour() < 12 {
		// 	slot = 1
		// } else if curTime.Hour() < 15 {
		// 	slot = 2
		// 	fmt.Println(slot)
		// }
		if curTime.Hour() <= 24 {
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

func PaymentPortal(request model.Request_Structure) *strike.Response_structure {

	paymentDetails := request.User_session_variables.PaymentDetails

	paymentStringArr := strings.Split(paymentDetails[0], ": Rs")
	paymentInt, _ := strconv.Atoi(paymentStringArr[1])
	// Get payment link from payment gateway
	payeeData := payment.PayeeData{
		PayeePrefix:      request.Bybrisk_session_variables.BusinessId,
		PayeeDisplayName: "Hemant Men's Salon",
		PaymentAmount:    paymentInt,
	}

	linkResponse := payment.RequestPaymentLink(request, payeeData)

	strikeObject := strike.Create("getting_started", model.Conf.APIEp+"payment/confirm_payment?linkid="+linkResponse.LinkID+"&linkurl="+linkResponse.LinkURL)

	question_object1 := strikeObject.Question("payment_confirmation").QuestionText().
		SetTextToQuestion("Hi "+request.Bybrisk_session_variables.Username+", click the link below to proceed with the payment.", "Text Description, getting used for testing purpose.")

	if linkResponse.LinkURL == "PAYMENT FAILED" {
		question_object1.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
			AnswerCard().SetHeaderToAnswer(10, strike.HALF_WIDTH).
			AddGraphicRowToAnswer(strike.PICTURE_ROW, []string{"https://m.media-amazon.com/images/I/71xyy-CkuUL._AC_SL1500_.jpg"}, []string{}).
			AddTextRowToAnswer(strike.H4, linkResponse.LinkURL, "#b56a00", false).
			AddTextRowToAnswer(strike.H5, "It's not you it's us, please try again.", "#343b40", false)
	} else {
		question_object1.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
			AnswerCard().SetHeaderToAnswer(10, strike.FULL_WIDTH).
			AddGraphicRowToAnswer(strike.PICTURE_ROW, []string{"https://ecommercenews.eu/wp-content/uploads/2013/06/most_common_payment_methods_in_europe.png"}, []string{}).
			AddTextRowToAnswer(strike.H4, linkResponse.LinkURL, "#0285d6", false).
			AddTextRowToAnswer(strike.H5, "After payment click on confirm below to complete the booking", "#343b40", false).
			AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).
			AddTextRowToAnswer(strike.H4, "Confirm ✅", "#438c46", true)
	}

	return strikeObject

}

func ConfirmPaymentStatus(request model.Request_Structure, linkID, linkURL string) *strike.Response_structure {
	status := payment.GetPaymentStatus(linkID)
	strikeObject := strike.Create("getting_started", model.Conf.APIEp+"payment/confirm_payment?linkid="+linkID+"&linkurl="+linkURL)

	if status == "PAID" {
		strikeObject.Question("").
			QuestionCard().SetHeaderToQuestion(10, strike.HALF_WIDTH).
			AddTextRowToQuestion(strike.H4, "Great", "#438c46", true).
			AddTextRowToQuestion(strike.H4, "Your booking is successful", "#3c3d3c", false)
	} else if status == "ACTIVE" {
		strikeObject.Question("payment_reconfirm").
			QuestionCard().SetHeaderToQuestion(2, strike.HALF_WIDTH).
			AddTextRowToQuestion(strike.H4, "Oops!", "#3c3d3c", true).
			AddTextRowToQuestion(strike.H4, "Your payment is still pending", "#3c3d3c", false).
			AddTextRowToQuestion(strike.H4, linkURL, "#0285d6", false).
			AddTextRowToQuestion(strike.H5, "Retry confirming after payment is done", "#3c3d3c", false).
			Answer(true).AnswerCardArray(strike.VERTICAL_ORIENTATION).AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).
			AddTextRowToAnswer(strike.H4, "Confirm ✅", "#438c46", true)
	} else if status == "EXPIRED" {
		strikeObject.Question("payment_reconfirm").
			QuestionCard().SetHeaderToQuestion(2, strike.HALF_WIDTH).
			AddTextRowToQuestion(strike.H4, "Oops!", "#3c3d3c", true).
			AddTextRowToQuestion(strike.H4, "Your payment link is expired", "#3c3d3c", false).
			AddTextRowToQuestion(strike.H5, "Retry confirming after payment is done", "#3c3d3c", false).
			Answer(true).AnswerCardArray(strike.VERTICAL_ORIENTATION).AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).
			AddTextRowToAnswer(strike.H4, "Confirm ✅", "#438c46", true)
	} else {
		strikeObject.Question("payment_reconfirm").
			QuestionCard().SetHeaderToQuestion(2, strike.HALF_WIDTH).
			AddTextRowToQuestion(strike.H4, "Oops!", "#3c3d3c", true).
			AddTextRowToQuestion(strike.H4, "Some error occured", "#3c3d3c", false).
			AddTextRowToQuestion(strike.H5, "Retry confirming after payment is done", "#3c3d3c", false)
	}

	return strikeObject
}
