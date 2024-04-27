package jsonStruct

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
}

type UserWithId struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
}
