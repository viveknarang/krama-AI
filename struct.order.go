package main

//ORDER document structure
type ORDER struct {
	OrderID           string             `json:"OrderID"`
	OrderCreationDate int64              `json:"OrderCreationDate"`
	OrderUpdateDate   int64              `json:"OrderUpdateDate"`
	CustomerID        string             `json:"CustomerID"`
	Products          map[string]PRODUCT `json:"Products"`
	ProductQuantity   map[string]int64   `json:"ProductQuantity"`
	PaymentStatus     string             `json:"PaymentStatus"`
	PaymentAmount     float64            `json:"PaymentAmount"`
	Currency          string             `json:"Currency"`
	OrderStatus       string             `json:"OrderStatus"`
	ShippingAddress   ADDRESS            `json:"ShippingAddress"`
	Attributes        interface{}        `json:"Attributes"`
}
