package strike

import (
	"github.com/go-playground/validator/v10"
)

//post request for registering a user
type Strike_Meta_Request_Structure struct {

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
}

// Response structure
// Boiler plate response strucutre

type Response_structure struct {

	// Internal status of the API
	//
	Status int64 `json:"status"`

	// Body structure
	//
	Body *Body_structure `json:"body"`
}

type Body_structure struct {

	// Handler name of the API chain
	//
	ActionHandler string `json:"actionHandler,omitempty"`

	// URI of the next API in the chain
	//
	NextActionHandler string `json:"nextActionHandlerURL,omitempty"`

	// Question Array structure
	//
	QuestionArray []Transaction_structure `json:"questionArray,omitempty"`
}

type Transaction_structure struct {

	// Question object
	//
	Question Question_structure `json:"question,omitempty"`

	// Answer Object
	//
	Answer1 Answer_structure `json:"answer,omitempty"`
}

type Question_structure struct {

	// UI type of the question
	//
	QuestionType string `json:"questionType,omitempty"`

	// This would depend on the questionType
	//
	// This defines Text for the UI
	QText string `json:"qText,omitempty"`

	//This defines Card object for the UI
	//
	QCard []Card_Row_Object `json:"qCard,omitempty"`

	// Context of the question
	// The value will be binded in this key
	//
	QContext string `json:"qContext,omitempty"`

	// Discription of the data strucutre of the question object for strike UI engine
	//
	QuestionDS string `json:"questionDS,omitempty"`
}

type Answer_structure struct {

	// UI type of the answer
	//
	ResponseType string `json:"responseType,omitempty"`

	// This would depend on the responseType
	// So we use all the type for meta description
	//
	//This defines Card object for the UI
	//
	QCard [][]Card_Row_Object `json:"qCard,omitempty"`

	//This defines the text for the location input UI
	//
	QLocationInput []string `json:"qLocation-Input,omitempty"`

	//Set this value if card is being used
	// HORIZONTAL or VERTICAL
	CardOrientation string `json:"card-orientation,omitempty"`

	// Set this to true if we want multiple values could be selected by the user for this particular question
	//
	MultipleSelect bool `json:"multipleSelect,omitempty"`

	// Discription of the data strucutre of the question object for strike UI engine
	//
	ResponseDS string `json:"responseDS,omitempty"`
}

// UI specific structures

type Card_Row_Object struct {

	// Type of row object
	// valid values are header, pic_row, h1, h2, h3, h4, h5, h6, price_row, video_row
	//
	Type string `json:"type,omitempty"`

	// Discriptor of the row object
	//
	Descriptor Descriptor_Structure `json:"descriptor,omitempty"`
}

type Descriptor_Structure struct {

	// context-objext is used when the type is set to header
	// It defines value of which row object to select if the user selects that card
	// Must be equal to the index of the row
	//
	ContextObject int `json:"context-object,omitempty"`

	// card-type is used when the type is set to header
	// FULL is used when we want our card to take full width of the screen else HALF is used
	//
	CardType string `json:"card-type,omitempty"`

	// Values of the row for the user to see
	// It will always be an array of strings
	//
	Value []string `json:"value,omitempty"`

	// Thumbnail is the picture url for the video row
	// It will always be an array of strings
	//
	Thumbnail []string `json:"thumbnail,omitempty"`

	// Color is set when the type is h1, h2, etc
	// It defines the color of the text
	// We can enter any hexadecimal value too
	//
	Color string `json:"color,omitempty"`

	// Bold is set when the type is h1, h2, etc
	// It is set to make the text bold
	//
	Bold bool `json:"bold,omitempty"`

	// Original is set when the row type is price_row
	// It is the original price of the item
	//
	Original float64 `json:"original,omitempty"`

	// Currency is set when the type is price_row
	// It defines the currency we are dealing with
	//
	Currency string `json:"currency,omitempty"`
}

func (d *Strike_Meta_Request_Structure) ValidateStrike_Meta_Request_Structure() error {
	validate := validator.New()
	return validate.Struct(d)
}
