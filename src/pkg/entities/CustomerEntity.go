package entities

// CustomerEntity related with response model
type CustomerEntity struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
