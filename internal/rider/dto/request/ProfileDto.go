package request

type ProfileDto struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	AvatarUrl string 	`json:"avatar_url"`
}