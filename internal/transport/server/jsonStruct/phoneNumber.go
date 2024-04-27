package jsonStruct

type PhoneNumberWithoutId struct {
	StudentId   string `json:"student_id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

type PhoneNumber struct {
	Id          string `json:"id"`
	StudentId   string `json:"student_id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
