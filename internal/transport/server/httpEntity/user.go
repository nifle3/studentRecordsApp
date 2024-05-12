package httpEntity

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	// Role must be "Сотрудник" or "Админ"
	Role string `json:"role"`
}
