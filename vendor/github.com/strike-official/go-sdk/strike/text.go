package strike

// Text is only used as Question. See TextInput for more info.
func (t *Transaction_structure) QuestionText() *Transaction_structure {
	t.Question.QuestionType = "Text"

	StrikeGlobal.Body.QuestionArray = append(StrikeGlobal.Body.QuestionArray, *t)

	return t
}

func (t *Transaction_structure) SetTextToQuestion(text, desc string) *Transaction_structure {
	t.Question.QText = text
	t.Question.QuestionDS = desc

	Update_Question_Array(t)

	return t
}
