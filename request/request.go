package request

type Contact struct {
	Name *string `json:"name"`
}
type UpdateContact struct {
	Previous_number *string `json:"previousnumber"`
	Number          *string `json:"number"`
}
