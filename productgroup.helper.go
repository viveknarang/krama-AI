package main

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func syncProductGroup(w http.ResponseWriter, r *http.Request, p PRODUCT) bool {

	var response bool

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductGroupExtension

	results := find(ExternalDB, dbcol, bson.M{"groupID": p.GroupID})

	if r.Method == http.MethodPost {

		if len(results) == 0 {

			var npg PRODUCTGROUP
			npg.GroupID = p.GroupID
			npg.Name = p.Name
			npg.RegularPriceMin = p.RegularPrice
			npg.RegularPriceMax = p.RegularPrice
			npg.PromotionPriceMin = p.PromotionPrice
			npg.PromotionPriceMax = p.PromotionPrice
			npg.Description = p.Description
			npg.Active = p.Active
			npg.Currency = p.Currency

			setInit()
			addAllInSet(p.SearchKeywords)
			npg.SearchKeywords = toArrayFromSet()

			setInit()
			addInSet(p.Size)
			npg.Sizes = toArrayFromSet()

			setInit()
			addInSet(p.Color)
			npg.Colors = toArrayFromSet()

			setInit()
			addInSet(p.Brand)
			npg.Brands = toArrayFromSet()

			setInit()
			addInSet(p.Sku)
			npg.Skus = toArrayFromSet()

			setInit()
			addAllInSet(p.Category)
			npg.Category = toArrayFromSet()

			setInit()
			addAllInSet(p.Images)
			npg.Images = toArrayFromSet()

			npg.Updated = time.Now().UnixNano()

			insert(ExternalDB, dbcol, npg)

			response = true

		} else {

		}

	} else if r.Method == http.MethodPut {

	} else if r.Method == http.MethodDelete {

	}

	return response
}
