package models

import (
	"github.com/google/uuid"
)

// Order represents a collection of products that should be shipped to the specified shipping address
type Order struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Products []Product `json:"products"`
	Customer Customer  `json:"customer"`
}

// Product represents a single product in an order
type Product struct {
	ProductCode string `json:"productCode"`
	Quantity    int    `json:"quantity"`
}

// Address represents an customers address
type Address struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2,omitempty"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
}

// Customer represents information about the customer placing the order
type Customer struct {
	FirstName       string  `json:"firstName"`
	LastName        string  `json:"lastName"`
	EmailAddress    string  `json:"emailAddress"`
	ShippingAddress Address `json:"shippingAddress"`
}
