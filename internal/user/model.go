package user

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" validate:"required,min=2,max=50"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=18,lte=100"`
}
