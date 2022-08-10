package strike

const (
	HALF_WIDTH string = "HALF"
	FULL_WIDTH        = "FULL"
	PICTURE_ROW        = "pic_row"
	VIDEO_ROW        = "video_row"
	HORIZONTAL_ORIENTATION       = "HORIZONTAL"
	VERTICAL_ORIENTATION       = "VERTICAL"
	H1        = "h1"
	H2        = "h2"
	H3        = "h3"
	H4        = "h4"
	H5        = "h5"
	H6        = "h6"
)

func (t *Transaction_structure) QuestionCard() *Transaction_structure {
	context := t.Question.QContext

	t = &Transaction_structure{
		Question: Question_structure{
			QuestionType: "Card",
			QContext:     context,
			QCard:        []Card_Row_Object{},
		},
	}
	StrikeGlobal.Body.QuestionArray = append(StrikeGlobal.Body.QuestionArray, *t)

	return t
}

func (t *Transaction_structure) SetHeaderToQuestion(card_context int, width string) *Transaction_structure {
	context := t.Question.QContext
	qcard := t.Question.QCard

	card_row := Card_Row_Object{
		Type: "header",
		Descriptor: Descriptor_Structure{
			ContextObject: card_context,
			CardType:      width,
		},
	}

	qcard = append(qcard, card_row)

	t = &Transaction_structure{
		Question: Question_structure{
			QuestionType: "Card",
			QContext:     context,
			QCard:        qcard,
		},
	}

	Update_Question_Array(t)

	return t
}

func (t *Transaction_structure) AddTextRowToQuestion(row_type string, value string, color string, boldr bool) *Transaction_structure {
	context := t.Question.QContext
	qcard := t.Question.QCard

	card_row := Card_Row_Object{
		Type: row_type,
		Descriptor: Descriptor_Structure{
			Value: []string{value},
			Bold:  boldr,
			Color: color,
		},
	}

	qcard = append(qcard, card_row)

	t = &Transaction_structure{
		Question: Question_structure{
			QuestionType: "Card",
			QContext:     context,
			QCard:        qcard,
		},
	}

	Update_Question_Array(t)

	return t
}

//The thumbnail_url is an array of url for the thumbnail of the video , it is used only when graphic_type is video_row
//If the graphic_type is pic_row then thumbnail_url is an empty string array
func (t *Transaction_structure) AddGraphicRowToQuestion(graphic_type string, url []string, thumbnail_url []string) *Transaction_structure {
	context := t.Question.QContext
	qcard := t.Question.QCard

	card_row := Card_Row_Object{
		Type: graphic_type,
		Descriptor: Descriptor_Structure{
			Value: url,
			Thumbnail: thumbnail_url,
		},
	}

	qcard = append(qcard, card_row)

	t = &Transaction_structure{
		Question: Question_structure{
			QuestionType: "Card",
			QContext:     context,
			QCard:        qcard,
		},
	}

	Update_Question_Array(t)

	return t
}

func (t *Transaction_structure) AnswerCardArray(card_orientation string) *Transaction_structure {
	q := t.Question
	ms := t.Answer1.MultipleSelect

	t = &Transaction_structure{
		Question: q,
		Answer1: Answer_structure{
			MultipleSelect:  ms,
			ResponseType:    "Card",
			CardOrientation: card_orientation,
			QCard:           [][]Card_Row_Object{},
		},
	}

	Update_Question_Array(t)

	return t
}

func (t *Transaction_structure) AnswerCard() *Transaction_structure {
	q := t.Question
	ms := t.Answer1.MultipleSelect
	co := t.Answer1.CardOrientation
	qcard := t.Answer1.QCard

	card := []Card_Row_Object{}

	qcard = append(qcard, card)

	t = &Transaction_structure{
		Question: q,
		Answer1: Answer_structure{
			MultipleSelect:  ms,
			ResponseType:    "Card",
			CardOrientation: co,
			QCard:           qcard,
		},
	}

	Update_Question_Array(t)

	return t
}

func (t *Transaction_structure) SetHeaderToAnswer(card_context int, width string) *Transaction_structure {
	q := t.Question
	ms := t.Answer1.MultipleSelect
	co := t.Answer1.CardOrientation
	qcard := t.Answer1.QCard

	//last array
	card := qcard[len(qcard)-1]

	card_row := Card_Row_Object{
		Type: "header",
		Descriptor: Descriptor_Structure{
			ContextObject: card_context,
			CardType:      width,
		},
	}

	card = append(card, card_row)

	//remove the last element form qcard
	//Add the modified card to qcard
	qcard = Update_QCard_Array(qcard, card)

	t = &Transaction_structure{
		Question: q,
		Answer1: Answer_structure{
			MultipleSelect:  ms,
			ResponseType:    "Card",
			CardOrientation: co,
			QCard:           qcard,
		},
	}

	Update_Question_Array(t)

	return t
}

//The thumbnail_url is an array of url for the thumbnail of the video , it is used only when graphic_type is video_row
//If the graphic_type is pic_row then thumbnail_url is an empty string array

func (t *Transaction_structure) AddGraphicRowToAnswer(graphic_type string, url []string, thumbnail_url []string) *Transaction_structure {
	q := t.Question
	ms := t.Answer1.MultipleSelect
	co := t.Answer1.CardOrientation
	qcard := t.Answer1.QCard

	//last array
	card := qcard[len(qcard)-1]

	card_row := Card_Row_Object{
		Type: graphic_type,
		Descriptor: Descriptor_Structure{
			Value: url,
			Thumbnail: thumbnail_url,
		},
	}

	card = append(card, card_row)

	//remove the last element form qcard
	//Add the modified card to qcard
	qcard = Update_QCard_Array(qcard, card)

	t = &Transaction_structure{
		Question: q,
		Answer1: Answer_structure{
			MultipleSelect:  ms,
			ResponseType:    "Card",
			CardOrientation: co,
			QCard:           qcard,
		},
	}

	Update_Question_Array(t)

	return t
}

func (t *Transaction_structure) AddTextRowToAnswer(row_type string, value string, color string, boldr bool) *Transaction_structure {
	q := t.Question
	ms := t.Answer1.MultipleSelect
	co := t.Answer1.CardOrientation
	qcard := t.Answer1.QCard

	//last array
	card := qcard[len(qcard)-1]

	card_row := Card_Row_Object{
		Type: row_type,
		Descriptor: Descriptor_Structure{
			Value: []string{value},
			Color: color,
			Bold:  boldr,
		},
	}

	card = append(card, card_row)

	//remove the last element form qcard
	//Add the modified card to qcard
	qcard = Update_QCard_Array(qcard, card)

	t = &Transaction_structure{
		Question: q,
		Answer1: Answer_structure{
			MultipleSelect:  ms,
			ResponseType:    "Card",
			CardOrientation: co,
			QCard:           qcard,
		},
	}

	Update_Question_Array(t)

	return t
}
