package strike

// Text is only used as Question. See TextInput for more info.
func (t *Transaction_structure) TextInput(desc string) *Transaction_structure {
	t.Answer1 = Answer_structure{
		MultipleSelect: false,
		ResponseType:   "Text-Input",
	}
	Update_Question_Array(t)
	return t
}
