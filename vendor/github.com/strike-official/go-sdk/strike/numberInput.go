package strike

// Number input is only used in answer. Can be used to let user enter number
func (t *Transaction_structure) NumberInput(desc string) *Transaction_structure {
	t.Answer1 = Answer_structure{
		MultipleSelect: false,
		ResponseType:   "Number-Input",
	}
	Update_Question_Array(t)
	return t
}