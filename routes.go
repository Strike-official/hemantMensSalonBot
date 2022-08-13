package main

import (
	"github.com/Strike-official/hemantMensSalonBot/internal/controller"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	cb := router.Group("/hemantMensSalon")
	{
		// Book, Cancel, Booking History
		cb.POST("/book", controller.BookAppointment)
		// cb.POST("/cancel", controller.CreateBot_2)
		// cb.POST("/bookingHistory", controller.CreateBot_3)

		// Book > Select Services
		// cb.POST("/services", controller.YourBots)
		// // Book > Select Services > Slots
		cb.POST("/slotSelection", controller.SlotSelection)
		// // Book > Select Services > Slots > PaymentGateway
		cb.POST("/paymentLink", controller.PaymentPortal)
		cb.POST("/payment/confirm_payment", controller.ConfirmPayment)

		// Cancel

		// Booking History
	}
}

// UserID
// PhoneNumber
// SlotStartTime
// Duration
// SlotEndTime
// SlotID
// PaymentLink
// PaymentStatus

// QueueingLogic

// 1. When a new booking comes in:
// 	- Fetch Current Date and Time
// 	- From DB, get all the bookings for CurDate.

// - Calculate Possible Slot based on the Time.
// - Find Next Best Slot for each Possible Slot.
// - Send it to Customer, and based on the selection.
// - Add it to the DB as a new entry.
