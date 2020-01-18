package main

//CUSTOMER document structure
type CUSTOMER struct {
	CustomerID     string       `json:"CustomerID" bson:"CustomerID"`
	Active         bool         `json:"Active" bson:"Active"`
	FirstName      string       `json:"FirstName" bson:"FirstName"`
	LastName       string       `json:"LastName" bson:"LastName"`
	Email          string       `json:"Email" bson:"Email"`
	PhoneNumbers   []string     `json:"PhoneNumbers" bson:"PhoneNumbers"`
	Password       string       `json:"Password" bson:"Password"`
	AddressBook    []ADDRESS    `json:"AddressBook" bson:"AddressBook"`
	PaymentOptions []CREDITCARD `json:"PaymentOptions" bson:"PaymentOptions"`
	WishList       []string     `json:"WishList" bson:"WishList"`
	SaveForLater   []string     `json:"SaveForLater" bson:"SaveForLater"`
	Updated        int64        `json:"Updated" bson:"Updated"`
}
