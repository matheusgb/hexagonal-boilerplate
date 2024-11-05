package dtos

type UserInputDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type UserOutputDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
