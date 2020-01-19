package main

import "strings"

func groomCustomerData(customer *CUSTOMER) {

	customer.FirstName = strings.TrimSpace(customer.FirstName)
	customer.LastName = strings.TrimSpace(customer.LastName)
	customer.CustomerID = strings.TrimSpace(customer.CustomerID)
	customer.Email = strings.TrimSpace(customer.Email)

	for i, phoneNumber := range customer.PhoneNumbers {
		customer.PhoneNumbers[i] = strings.TrimSpace(phoneNumber)
	}

	customer.Password = strings.TrimSpace(customer.Password)

	for i, address := range customer.AddressBook {

		customer.AddressBook[i].FirstName = strings.TrimSpace(address.FirstName)
		customer.AddressBook[i].LastName = strings.TrimSpace(address.LastName)
		customer.AddressBook[i].AddressLineOne = strings.TrimSpace(address.AddressLineOne)
		customer.AddressBook[i].AddressLineTwo = strings.TrimSpace(address.AddressLineTwo)
		customer.AddressBook[i].City = strings.TrimSpace(address.City)
		customer.AddressBook[i].State = strings.TrimSpace(address.State)
		customer.AddressBook[i].Country = strings.TrimSpace(address.Country)
		customer.AddressBook[i].Pincode = strings.TrimSpace(address.Pincode)

	}

	for i, paymentOption := range customer.PaymentOptions {

		customer.PaymentOptions[i].Name = strings.TrimSpace(paymentOption.Name)
		customer.PaymentOptions[i].CardNumber = strings.TrimSpace(paymentOption.CardNumber)
		customer.PaymentOptions[i].CardExpiryMM = strings.TrimSpace(paymentOption.CardExpiryMM)
		customer.PaymentOptions[i].CardExpiryYY = strings.TrimSpace(paymentOption.CardExpiryYY)
		customer.PaymentOptions[i].SecurityCode = strings.TrimSpace(paymentOption.SecurityCode)
		customer.PaymentOptions[i].ZipCode = strings.TrimSpace(paymentOption.ZipCode)

	}

	return

}
