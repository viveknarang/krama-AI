package main

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func syncProductGroup(w http.ResponseWriter, r *http.Request, p PRODUCT) bool {

	var response bool

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductGroupExtension

	results := find(ExternalDB, dbcol, bson.M{"groupid": p.GroupID})

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

			pm := make(map[string]PRODUCT)
			pm[p.Sku] = p
			npg.Products = pm

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

			var productGroup PRODUCTGROUP

			j, err0 := bson.MarshalExtJSON(results[0], false, false)

			if err0 != nil {
				respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError)
				return false
			}

			err1 := json.Unmarshal([]byte(j), &productGroup)

			if err1 != nil {
				respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError)
				return false
			}

			setInit()
			addAllInSet(p.SearchKeywords)
			addAllInSet(productGroup.SearchKeywords)
			productGroup.SearchKeywords = toArrayFromSet()

			setInit()
			addInSet(p.Size)
			addAllInSet(productGroup.Sizes)
			productGroup.Sizes = toArrayFromSet()

			setInit()
			addInSet(p.Color)
			addAllInSet(productGroup.Colors)
			productGroup.Colors = toArrayFromSet()

			setInit()
			addInSet(p.Brand)
			addAllInSet(productGroup.Brands)
			productGroup.Brands = toArrayFromSet()

			setInit()
			addAllInSet(p.Category)
			addAllInSet(productGroup.Category)
			productGroup.Category = toArrayFromSet()

			productGroup.Products[p.Sku] = p

			nrpmin := math.MaxFloat64
			var nrpmax float64
			nppmin := math.MaxFloat64
			var nppmax float64
			active := false

			setInit()

			for key, value := range productGroup.Products {

				if value.RegularPrice < nrpmin {
					nrpmin = value.RegularPrice
				}
				if value.RegularPrice > nrpmax {
					nrpmax = value.RegularPrice
				}
				if value.PromotionPrice < nppmin {
					nppmin = value.PromotionPrice
				}
				if value.PromotionPrice > nppmax {
					nppmax = value.PromotionPrice
				}

				if value.IsMain {
					productGroup.Name = value.Name
					productGroup.Description = value.Description
					productGroup.Images = value.Images
				}

				active = active || value.Active
				addInSet(key)

			}

			productGroup.Skus = append(productGroup.Skus, p.Sku)

			productGroup.RegularPriceMin = nrpmin
			productGroup.RegularPriceMax = nrpmax
			productGroup.PromotionPriceMin = nppmin
			productGroup.PromotionPriceMax = nppmax
			productGroup.Active = active
			productGroup.Skus = toArrayFromSet()

			result := update(ExternalDB, dbcol, bson.M{"groupid": p.GroupID}, bson.M{"$set": productGroup})

			if result[0] == 1 && result[1] == 1 {
				response = true
			} else {
				response = false
			}

		}

	} else if r.Method == http.MethodPut {

		var productGroup PRODUCTGROUP

		j, err0 := bson.MarshalExtJSON(results[0], false, false)

		if err0 != nil {
			respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError)
			return false
		}

		err1 := json.Unmarshal([]byte(j), &productGroup)

		if err1 != nil {
			respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError)
			return false
		}

		setInit()
		addAllInSet(p.SearchKeywords)
		addAllInSet(productGroup.SearchKeywords)
		productGroup.SearchKeywords = toArrayFromSet()

		setInit()
		addInSet(p.Size)
		addAllInSet(productGroup.Sizes)
		productGroup.Sizes = toArrayFromSet()

		setInit()
		addInSet(p.Color)
		addAllInSet(productGroup.Colors)
		productGroup.Colors = toArrayFromSet()

		setInit()
		addInSet(p.Brand)
		addAllInSet(productGroup.Brands)
		productGroup.Brands = toArrayFromSet()

		setInit()
		addAllInSet(p.Category)
		addAllInSet(productGroup.Category)
		productGroup.Category = toArrayFromSet()

		productGroup.Products[p.Sku] = p

		nrpmin := math.MaxFloat64
		var nrpmax float64
		nppmin := math.MaxFloat64
		var nppmax float64
		active := false

		setInit()

		for key, value := range productGroup.Products {

			if value.RegularPrice < nrpmin {
				nrpmin = value.RegularPrice
			}
			if value.RegularPrice > nrpmax {
				nrpmax = value.RegularPrice
			}
			if value.PromotionPrice < nppmin {
				nppmin = value.PromotionPrice
			}
			if value.PromotionPrice > nppmax {
				nppmax = value.PromotionPrice
			}

			if value.IsMain {
				productGroup.Name = value.Name
				productGroup.Description = value.Description
				productGroup.Images = value.Images
			}

			active = active || value.Active
			addInSet(key)

		}

		productGroup.Skus = append(productGroup.Skus, p.Sku)

		productGroup.RegularPriceMin = nrpmin
		productGroup.RegularPriceMax = nrpmax
		productGroup.PromotionPriceMin = nppmin
		productGroup.PromotionPriceMax = nppmax
		productGroup.Active = active
		productGroup.Skus = toArrayFromSet()

		result := update(ExternalDB, dbcol, bson.M{"groupid": p.GroupID}, bson.M{"$set": productGroup})

		if result[0] == 1 && result[1] == 1 {
			response = true
		} else {
			response = false
		}

	} else if r.Method == http.MethodDelete {

	}

	return response
}
