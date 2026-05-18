package request

type GetListProfileDto struct {
	Status *string `json:"status"`
	Limit  uint    `json:"limit"`
	Offset uint    `json:"offset"`
}
