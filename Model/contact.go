package Model

import "fmt"

const (
	TABLE_CONTACT        = "contact"
	TABLE_CONTACT_NUMBER = "number"
	TABLE_CONTACT_NAME   = "name"
)

type Contact struct {
	Number string `json:"number"`
	Name   string `json:"name"`
}

func (r *Contact) Validate() error {
	if r.Name == "" || r.Number == "" {
		return fmt.Errorf(" can't be empty")
	}
	if len([]rune(r.Number)) != 10 {
		return fmt.Errorf("Invalid number")
	}
	return nil
}
