package main

//ORDER document structure
type ORDER struct {
	OrderID           string             `json:"OrderID" bson:"OrderID"`
	OrderCreationDate int64              `json:"OrderCreationDate" bson:"OrderCreationDate"`
	OrderUpdateDate   int64              `json:"OrderUpdateDate" bson:"OrderUpdateDate"`
	CustomerID        string             `json:"CustomerID" bson:"CustomerID"`
	Products          map[string]PRODUCT `json:"Products" bson:"Products"`
	ProductQuantity   map[string]int64   `json:"ProductQuantity" bson:"ProductQuantity"`
	PaymentStatus     string             `json:"PaymentStatus" bson:"PaymentStatus"`
	PaymentAmount     float64            `json:"PaymentAmount" bson:"PaymentAmount"`
	Currency          string             `json:"Currency" bson:"Currency"`
	OrderStatus       string             `json:"OrderStatus" bson:"OrderStatus"`
	ShippingAddress   ADDRESS            `json:"ShippingAddress" bson:"ShippingAddress"`
	Attributes        interface{}        `json:"Attributes" bson:"Attributes"`
}
