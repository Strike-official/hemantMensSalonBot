package strike

// Date input is only used in answer. Can be used to let user enter date
func (t *Transaction_structure) DateInput(desc string) *Transaction_structure {
	t.Answer1 = Answer_structure{
		MultipleSelect: false,
		ResponseType:   "Date-Input",
	}
	Update_Question_Array(t)
	return t
}