package model

import (
	"fmt"
	"time"
)

const (
	layout     = "2006-01-02 18:17:22"
	TimeLatout = "2006-01-02 15:04:05"
)

var (
	SalonOpeningTime time.Time
	SalonClosingTime time.Time
)

const (
	HAIRCUT            = "Hair Cut"
	BEARDCUT           = "Beard Shave/Trim"
	FACEMESSAGE        = "Face Message"
	CLEANUP            = "Cleanup"
	DETAN              = "Detan"
	HAIRSPA            = "Hair Spa"
	FACIAL             = "Facial"
	ADVANCEFACIAL      = "Advance Facial"
	ADVANCEFACEMESSAGE = "Advance Face Message"
)

// Ex: "hairCut":60 (in mins)
var Services = map[string]int{
	HAIRCUT:            40,
	BEARDCUT:           15,
	FACEMESSAGE:        20,
	CLEANUP:            30,
	DETAN:              25,
	HAIRSPA:            30,
	FACIAL:             60,
	ADVANCEFACIAL:      120,
	ADVANCEFACEMESSAGE: 30,
}

var SlotDetail map[string]map[string]map[int]SlotTime

type SlotTime struct {
	SlotStartTime time.Time
	SlotEndTime   time.Time
}

type SlotDetails struct {
	Slot     int       `json:"Slot"`
	Chair    int       `json:"Chair"`
	ComeAt   time.Time `json:"ComeAt"`
	WaitTime int       `json:"WaitTime"`
	Cost     int       `json:"Cost"`
}

func iniSlots() {
	d := "2022-08-10 18:17:22"
	t, _ := time.Parse(layout, d)
	st := SlotTime{
		SlotStartTime: t,
		SlotEndTime:   t.Local().Add(time.Minute * time.Duration(55)),
	}

	SlotDetail["chair1"]["slot1"][1] = st

	fmt.Println(SlotDetail)
}
