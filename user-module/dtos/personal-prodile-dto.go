package dtos

type PersonalProfieDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}
