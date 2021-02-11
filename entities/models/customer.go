package models

// Customer :nodoc:
type Customer struct {
	Name        string `json:"customer_name" bson:"customer_name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Address     string `json:"address" bson:"address"`
}
