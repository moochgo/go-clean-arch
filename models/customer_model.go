package models

// Customer : Model name
type Customer struct {
	BaseModel
	Fullname string `json:"fullname"`
	PhoneNo  string `json:"phone_no"`
	Email    string `json:"email"`
}

// TableName : Set table name
func (Customer) TableName() string {
	return "customers"
}
