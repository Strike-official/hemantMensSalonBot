package controller

import (
	"fmt"

	"github.com/Strike-official/hemantMensSalonBot/internal/core"
	"github.com/Strike-official/hemantMensSalonBot/internal/model"
	"github.com/gin-gonic/gin"
)

func BookAppointment(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.BookAppointment(request)
	ctx.JSON(200, strikeObj)
}

func SlotSelection(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.SlotSelection(request)
	ctx.JSON(200, strikeObj)
}

func PaymentPortal(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	strikeObj := core.PaymentPortal(request)
	ctx.JSON(200, strikeObj)
}

func ConfirmPayment(ctx *gin.Context) {
	var request model.Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Error:", err)
	}
	linkID := ctx.Query("linkid")
	linkurl := ctx.Query("linkurl")
	strikeObj := core.ConfirmPaymentStatus(request, linkID, linkurl)
	ctx.JSON(200, strikeObj)
}
