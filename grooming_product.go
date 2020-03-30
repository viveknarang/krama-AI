package main

import (
	"strings"
)

func groomProductData(product *PRODUCT) {

	product.Sku = strings.TrimSpace(product.Sku)
	product.Name = strings.TrimSpace(product.Name)
	product.GroupID = strings.TrimSpace(product.GroupID)
	product.Description = strings.TrimSpace(product.Description)
	product.Size = strings.TrimSpace(product.Size)
	product.Color = strings.TrimSpace(product.Color)
	product.Brand = strings.TrimSpace(product.Brand)
	product.Currency = strings.TrimSpace(product.Currency)

	for i, val := range product.Images {

		product.Images[i] = strings.TrimSpace(val)

	}

	for i, val := range product.SearchKeywords {

		product.SearchKeywords[i] = strings.TrimSpace(val)

	}

	for i, val := range product.Category {

		product.Category[i] = cleanCategoryPath(val)

	}

}
