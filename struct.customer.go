package main

//CUSTOMER document structure
type CUSTOMER struct {
	CustomerID     string       `json:"CustomerID" bson:"CustomerID" validate:"min=1,max=100"`
	Active         bool         `json:"Active" bson:"Active"`
	FirstName      string       `json:"FirstName" bson:"FirstName" validate:"min=1,max=100"`
	LastName       string       `json:"LastName" bson:"LastName" validate:"min=1,max=100"`
	Email          string       `json:"Email" bson:"Email" validate:"min=3,max=40, validateEmail"`
	PhoneNumbers   []string     `json:"PhoneNumbers" bson:"PhoneNumbers" validate:"validateTypeArrayLength=10"`
	Password       string       `json:"Password" bson:"Password" validate:"min=5,max=1024"`
	AddressBook    []ADDRESS    `json:"AddressBook" bson:"AddressBook" validate:"validateTypeArrayLength=10"`
	PaymentOptions []CREDITCARD `json:"PaymentOptions" bson:"PaymentOptions" validate:"validateTypeArrayLength=50"`
	WishList       []string     `json:"WishList" bson:"WishList" validate:"validateTypeArrayLength=1000"`
	SaveForLater   []string     `json:"SaveForLater" bson:"SaveForLater" validate:"validateTypeArrayLength=1000"`
	Updated        int64        `json:"Updated" bson:"Updated"`
}
