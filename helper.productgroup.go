package main

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
)

func syncProductGroup(w http.ResponseWriter, r *http.Request, p PRODUCT) bool {

	rlog.Debug("syncProductGroup() handle function invoked ...")

	var response bool

	cidb := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()

	pgcol := cidb + ProductGroupExtension

	pgindex := cidb + ProductGroupExtension + SearchIndexExtension

	results := findMongoDocument(ExternalDB, pgcol, bson.M{"groupid": p.GroupID})

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

			npg.Attributes = make(map[string][]interface{})
			for key, value := range p.Attributes {

				var r []interface{}
				r = append(r, value)
				npg.Attributes[key] = r

			}

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

			insertMongoDocument(ExternalDB, pgcol, npg)
			npg.Products = nil
			response = indexES(pgindex, PGMapping, npg, npg.GroupID)

		} else {

			var productGroup PRODUCTGROUP

			j, err0 := bson.MarshalExtJSON(results[0], false, false)

			if err0 != nil {
				respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
				return false
			}

			err1 := json.Unmarshal([]byte(j), &productGroup)

			if err1 != nil {
				respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
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

			for key, value := range p.Attributes {

				var x []interface{}

				if productGroup.Attributes[key] != nil {
					x = append(x, productGroup.Attributes[key]...)
				}

				productGroup.Attributes[key] = append(x, value)

				gsetInit()
				addAllInGSet(productGroup.Attributes[key])
				productGroup.Attributes[key] = nil
				productGroup.Attributes[key] = toArrayFromGSet()

			}

			productGroup.Skus = append(productGroup.Skus, p.Sku)

			productGroup.RegularPriceMin = nrpmin
			productGroup.RegularPriceMax = nrpmax
			productGroup.PromotionPriceMin = nppmin
			productGroup.PromotionPriceMax = nppmax
			productGroup.Active = active
			productGroup.Skus = toArrayFromSet()

			result := updateMongoDocument(ExternalDB, pgcol, bson.M{"groupid": p.GroupID}, bson.M{"$set": productGroup})

			if result[0] == 1 && result[1] == 1 {
				productGroup.Products = nil
				response = indexES(pgindex, PGMapping, productGroup, productGroup.GroupID)
			} else {
				response = false
			}

		}

	} else if r.Method == http.MethodPut {

		var productGroup PRODUCTGROUP

		j, err0 := bson.MarshalExtJSON(results[0], false, false)

		if err0 != nil {
			respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return false
		}

		err1 := json.Unmarshal([]byte(j), &productGroup)

		if err1 != nil {
			respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
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

		updpg := make(map[string][]interface{})
		for _, valueP := range productGroup.Products {

			for key, value := range valueP.Attributes {

				updpg[key] = append(updpg[key], value)

			}

		}

		for key := range updpg {

			gsetInit()
			addAllInGSet(updpg[key])
			updpg[key] = nil
			updpg[key] = toArrayFromGSet()

		}

		productGroup.Attributes = updpg

		productGroup.Skus = append(productGroup.Skus, p.Sku)

		productGroup.RegularPriceMin = nrpmin
		productGroup.RegularPriceMax = nrpmax
		productGroup.PromotionPriceMin = nppmin
		productGroup.PromotionPriceMax = nppmax
		productGroup.Active = active
		productGroup.Skus = toArrayFromSet()

		result := updateMongoDocument(ExternalDB, pgcol, bson.M{"groupid": p.GroupID}, bson.M{"$set": productGroup})

		if result[0] == 1 && result[1] == 1 {
			productGroup.Products = nil
			response = indexES(pgindex, PGMapping, productGroup, productGroup.GroupID)
		} else {
			response = false
		}

	} else if r.Method == http.MethodDelete {

		var productGroup PRODUCTGROUP

		j, err0 := bson.MarshalExtJSON(results[0], false, false)

		if err0 != nil {
			respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return false
		}

		err1 := json.Unmarshal([]byte(j), &productGroup)

		if err1 != nil {
			respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return false
		}

		if len(productGroup.Products) == 1 {

			delr := deleteMongoDocument(ExternalDB, pgcol, bson.M{"groupid": p.GroupID})

			if delr == 1 {
				response = deleteESDocumentByID(pgindex, p.GroupID)
			} else {
				response = false
			}

		} else {

			delete(productGroup.Products, p.Sku)

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

			updpg := make(map[string][]interface{})
			for _, valueP := range productGroup.Products {

				for key, value := range valueP.Attributes {

					updpg[key] = append(updpg[key], value)

				}

			}

			for key := range updpg {

				gsetInit()
				addAllInGSet(updpg[key])
				updpg[key] = nil
				updpg[key] = toArrayFromGSet()

			}

			productGroup.Attributes = updpg

			productGroup.Skus = append(productGroup.Skus, p.Sku)

			productGroup.RegularPriceMin = nrpmin
			productGroup.RegularPriceMax = nrpmax
			productGroup.PromotionPriceMin = nppmin
			productGroup.PromotionPriceMax = nppmax
			productGroup.Active = active
			productGroup.Skus = toArrayFromSet()

			setInit()
			for _, prd := range productGroup.Products {
				addAllInSet(prd.SearchKeywords)
			}
			productGroup.SearchKeywords = toArrayFromSet()

			setInit()
			for _, prd := range productGroup.Products {
				addInSet(prd.Size)
			}
			productGroup.Sizes = toArrayFromSet()

			setInit()
			for _, prd := range productGroup.Products {
				addInSet(prd.Color)
			}
			productGroup.Colors = toArrayFromSet()

			setInit()
			for _, prd := range productGroup.Products {
				addInSet(prd.Brand)
			}
			productGroup.Brands = toArrayFromSet()

			setInit()
			for _, prd := range productGroup.Products {
				addAllInSet(prd.Category)
			}
			productGroup.Category = toArrayFromSet()

			result := updateMongoDocument(ExternalDB, pgcol, bson.M{"groupid": p.GroupID}, bson.M{"$set": productGroup})

			if result[0] == 1 && result[1] == 1 {
				productGroup.Products = nil
				response = indexES(pgindex, PGMapping, productGroup, productGroup.GroupID)
			} else {
				response = false
			}

		}

	}

	return response
}

func syncProductGroupFromProducts(w http.ResponseWriter, r *http.Request, skus []string, isPriceUpdate bool) bool {

	rlog.Debug("syncProductGroupFromProducts() handle function invoked ...")

	cidb := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()

	pgcol := cidb + ProductGroupExtension

	pcol := cidb + ProductExtension

	pgindex := cidb + ProductGroupExtension + SearchIndexExtension

	response := true

	for _, sku := range skus {

		results := findMongoDocument(ExternalDB, pcol, bson.M{"sku": sku})

		if len(results) != 1 {
			respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotFound, false)
			return false
		}

		j, err0 := bson.MarshalExtJSON(results[0], false, false)

		if err0 != nil {
			respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return false
		}

		var product PRODUCT

		err1 := json.Unmarshal([]byte(j), &product)

		if err1 != nil {
			respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return false
		}

		result := updateMongoDocument(ExternalDB, pgcol, bson.M{"groupid": product.GroupID}, bson.M{"$set": bson.M{"products." + product.Sku: product}})

		if result[0] == 1 && result[1] == 1 {

			results := findMongoDocument(ExternalDB, pgcol, bson.M{"groupid": product.GroupID})

			var productGroup PRODUCTGROUP

			j, err0 := bson.MarshalExtJSON(results[0], false, false)

			if err0 != nil {
				respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
				return false
			}

			err1 := json.Unmarshal([]byte(j), &productGroup)

			if err1 != nil {
				respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
				return false
			}

			if isPriceUpdate {

				nrpmin := math.MaxFloat64
				var nrpmax float64
				nppmin := math.MaxFloat64
				var nppmax float64

				for _, value := range productGroup.Products {

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

				}

				result := updateMongoDocument(ExternalDB, pgcol, bson.M{"groupid": product.GroupID}, bson.M{"$set": bson.M{"regularpricemin": nrpmin, "regularpricemax": nrpmax, "promotionpricemin": nppmin, "promotionpricemax": nppmax}})

				if result[0] == 1 && result[1] == 1 {

					results := findMongoDocument(ExternalDB, pgcol, bson.M{"groupid": product.GroupID})

					var productGroup PRODUCTGROUP

					j, err0 := bson.MarshalExtJSON(results[0], false, false)

					if err0 != nil {
						respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
						return false
					}

					err1 := json.Unmarshal([]byte(j), &productGroup)

					if err1 != nil {
						respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
						return false
					}

					productGroup.Products = nil
					response = response && indexES(pgindex, PGMapping, productGroup, productGroup.GroupID)

				}

			} else {

				productGroup.Products = nil
				response = response && indexES(pgindex, PGMapping, productGroup, productGroup.GroupID)

			}

			resetProductCacheKeys(nil, &productGroup)

		}

	}

	return response
}
