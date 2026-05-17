package request

type RegisterDto struct {
	Username    string `json: "username"`
	Password    string `json: "password"`
	PhoneNumber string `json: "phone_number"`
	Email       string `json: "email"`
}
