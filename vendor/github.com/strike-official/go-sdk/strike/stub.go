package strike

import (
	"encoding/json"
	"fmt"
)

// var StrikeGlobals *Response_wrapper_structure
var StrikeGlobal *Response_structure

func Create(handler string, next string) *Response_structure {
	StrikeGlobal = &Response_structure{
		Status: 200,
		Body: &Body_structure{
			ActionHandler:     handler,
			NextActionHandler: next,
		},
	}

	return StrikeGlobal
}

func (create *Response_structure) Question(context string) *Transaction_structure {

	t := Transaction_structure{
		Question: Question_structure{
			QContext: context,
		},
	}

	return &t
}

func (t *Transaction_structure) Answer(multiple_select bool) *Transaction_structure {

	q := t.Question
	t = &Transaction_structure{
		Question: q,
		Answer1: Answer_structure{
			MultipleSelect: multiple_select,
		},
	}

	Update_Question_Array(t)

	return t
}

//Helper function

func Update_Question_Array(t *Transaction_structure) {
	newArray := []Transaction_structure{}

	if len(StrikeGlobal.Body.QuestionArray)-1 < 0 {
		newArray = StrikeGlobal.Body.QuestionArray[:0]
	} else {
		newArray = StrikeGlobal.Body.QuestionArray[:len(StrikeGlobal.Body.QuestionArray)-1]
	}

	newArray = append(newArray, *t)
}

func Update_QCard_Array(qcard [][]Card_Row_Object, t []Card_Row_Object) [][]Card_Row_Object {
	newArray := [][]Card_Row_Object{}

	if len(qcard)-1 < 0 {
		newArray = qcard[:0]
	} else {
		newArray = qcard[:len(qcard)-1]
	}

	newArray = append(newArray, t)
	return newArray
}

func (create Response_structure) ToJson() []byte {
	b, err := json.Marshal(create)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
