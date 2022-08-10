package model

import "github.com/Strike-official/hemantMensSalonBot/configmanager"

var Conf *configmanager.AppConfig

type Request_Structure struct {

	// Bybrisk variable from strike bot
	//
	Bybrisk_session_variables Bybrisk_session_variables_struct `json: "bybrisk_session_variables"`

	// Our own variable from previous API
	//
	User_session_variables User_session_variables_struct `json: "user_session_variables"`
}

type Bybrisk_session_variables_struct struct {

	// User ID on Bybrisk
	//
	UserId string `json:"userId"`

	// Our own business Id in Bybrisk
	//
	BusinessId string `json:"businessId"`

	// Handler Name for the API chain
	//
	Handler string `json:"handler"`

	// Current location of the user
	//
	Location GeoLocation_struct `json:"location"`

	// Username of the user
	//
	Username string `json:"username"`

	// Address of the user
	//
	Address string `json:"address"`

	// Phone number of the user
	//
	Phone string `json:"phone"`
}

type GeoLocation_struct struct {
	// Latitude
	//
	Latitude float64 `json:"latitude"`

	// Longitude
	//
	Longitude float64 `json:"longitude"`
}

type User_session_variables_struct struct {
	Preference         []string           `json:"preference,omitempty"`
	Lat_long           GeoLocation_struct `json:"lat_long,omitempty"`
	UserName           string             `json:"userName,omitempty"`
	BusinessName       string             `json:"businessName"`
	BusinessCategory   string             `json:"businessCategory"`
	PicURL             string             `json:"picURL"`
	Title              string             `json:"title"`
	Subtitle           string             `json:"subtitle"`
	Story              string             `json:"story"`
	BotUserName        []string           `json:"botUserName"`
	ActionOnBot        []string           `json:"actionOnBot"`
	ActionHandlerName  []string           `json:"actionHandlerName"`
	ActionOnHandler    []string           `json:"actionOnHandler"`
	ModifyDetailField  []string           `json:"modifyDetailField"`
	ModifyDetailValue  string             `json:"modifyDetailValue"`
	NA                 []string           `json:"na"`
	NewHandlerName     string             `json:"newHandlerName"`
	NewHandlerEndpoint string             `json:"newHandlerEndpoint"`
	EmailId            string             `json:"emailId"`
	SelectedServices   []string           `json:"selectedServices"`
	PaymentDetails     []string           `json:"paymentAmount"`
}
