package dtos

type LoginDTO struct {
	Password string `json:"password"`
	Username string `json:"username" binding:"required"`
}
