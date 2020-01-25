package main

import (
	"math"
	"net/http"
	"time"

	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func syncProductGroup(w http.ResponseWriter, r *http.Request, p PRODUCT) bool {

	rlog.Debug("syncProductGroup() handle function invoked ...")

	var response bool

	cidb := getAccessToken(r)

	pgcol := cidb + ProductGroupExtension

	pgindex := cidb + ProductGroupExtension + SearchIndexExtension

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB+cidb, pgcol, bson.M{"GroupID": p.GroupID}, &opts)

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

			m0 := make(map[string]bool)
			addAllInSet(p.SearchKeywords, m0)
			npg.SearchKeywords = toArrayFromSet(m0)

			m1 := make(map[string]bool)
			addInSet(p.Size, m1)
			npg.Sizes = toArrayFromSet(m1)

			m2 := make(map[string]bool)
			addInSet(p.Color, m2)
			npg.Colors = toArrayFromSet(m2)

			m3 := make(map[string]bool)
			addInSet(p.Brand, m3)
			npg.Brands = toArrayFromSet(m3)

			m4 := make(map[string]bool)
			addInSet(p.Sku, m4)
			npg.Skus = toArrayFromSet(m4)

			m5 := make(map[string]bool)
			addAllInSet(p.Category, m5)
			npg.Category = toArrayFromSet(m5)

			npg.Images = p.Images

			npg.Updated = time.Now().UnixNano()

			if !insertMongoDocument(ExternalDB+cidb, pgcol, npg) {
				respondWith(w, r, nil, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
				return false
			}

			npg.Products = nil
			response = indexES(pgindex, PGMapping, npg, npg.GroupID)

		} else {

			var productGroup PRODUCTGROUP

			mapDocument(w, r, &productGroup, results[0])

			m6 := make(map[string]bool)
			addAllInSet(p.SearchKeywords, m6)
			addAllInSet(productGroup.SearchKeywords, m6)
			productGroup.SearchKeywords = toArrayFromSet(m6)

			m7 := make(map[string]bool)
			addInSet(p.Size, m7)
			addAllInSet(productGroup.Sizes, m7)
			productGroup.Sizes = toArrayFromSet(m7)

			m8 := make(map[string]bool)
			addInSet(p.Color, m8)
			addAllInSet(productGroup.Colors, m8)
			productGroup.Colors = toArrayFromSet(m8)

			m9 := make(map[string]bool)
			addInSet(p.Brand, m9)
			addAllInSet(productGroup.Brands, m9)
			productGroup.Brands = toArrayFromSet(m9)

			m10 := make(map[string]bool)
			addAllInSet(p.Category, m10)
			addAllInSet(productGroup.Category, m10)
			productGroup.Category = toArrayFromSet(m10)

			productGroup.Products[p.Sku] = p

			active := false

			m11 := make(map[string]bool)

			for key, value := range productGroup.Products {

				if value.IsMain {
					productGroup.Name = value.Name
					productGroup.Description = value.Description
					productGroup.Images = value.Images
				}

				active = active || value.Active
				addInSet(key, m11)

			}

			prices := computePriceRange(&productGroup)

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

			productGroup.RegularPriceMin = prices[0]
			productGroup.RegularPriceMax = prices[1]
			productGroup.PromotionPriceMin = prices[2]
			productGroup.PromotionPriceMax = prices[3]
			productGroup.Active = active
			productGroup.Skus = toArrayFromSet(m11)

			result := updateMongoDocument(ExternalDB+cidb, pgcol, bson.M{"GroupID": p.GroupID}, bson.M{"$set": productGroup})

			if result[0] == 1 && result[1] == 1 {
				productGroup.Products = nil
				response = indexES(pgindex, PGMapping, productGroup, productGroup.GroupID)
			} else {
				response = false
			}

		}

	} else if r.Method == http.MethodPut {

		var productGroup PRODUCTGROUP

		mapDocument(w, r, &productGroup, results[0])

		m12 := make(map[string]bool)
		addAllInSet(p.SearchKeywords, m12)
		addAllInSet(productGroup.SearchKeywords, m12)
		productGroup.SearchKeywords = toArrayFromSet(m12)

		m13 := make(map[string]bool)
		addInSet(p.Size, m13)
		addAllInSet(productGroup.Sizes, m13)
		productGroup.Sizes = toArrayFromSet(m13)

		m14 := make(map[string]bool)
		addInSet(p.Color, m14)
		addAllInSet(productGroup.Colors, m14)
		productGroup.Colors = toArrayFromSet(m14)

		m15 := make(map[string]bool)
		addInSet(p.Brand, m15)
		addAllInSet(productGroup.Brands, m15)
		productGroup.Brands = toArrayFromSet(m15)

		m16 := make(map[string]bool)
		addAllInSet(p.Category, m16)
		addAllInSet(productGroup.Category, m16)
		productGroup.Category = toArrayFromSet(m16)

		productGroup.Products[p.Sku] = p

		active := false

		m17 := make(map[string]bool)

		for key, value := range productGroup.Products {

			if value.IsMain {
				productGroup.Name = value.Name
				productGroup.Description = value.Description
				productGroup.Images = value.Images
			}

			active = active || value.Active
			addInSet(key, m17)

		}

		prices := computePriceRange(&productGroup)

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

		productGroup.RegularPriceMin = prices[0]
		productGroup.RegularPriceMax = prices[1]
		productGroup.PromotionPriceMin = prices[2]
		productGroup.PromotionPriceMax = prices[3]

		productGroup.Active = active
		productGroup.Skus = toArrayFromSet(m17)

		result := updateMongoDocument(ExternalDB+cidb, pgcol, bson.M{"GroupID": p.GroupID}, bson.M{"$set": productGroup})

		if result[0] == 1 && result[1] == 1 {
			productGroup.Products = nil
			response = indexES(pgindex, PGMapping, productGroup, productGroup.GroupID)
		} else {
			response = false
		}

	} else if r.Method == http.MethodDelete {

		var productGroup PRODUCTGROUP

		mapDocument(w, r, &productGroup, results[0])

		if len(productGroup.Products) == 1 {

			delr := deleteMongoDocument(ExternalDB+cidb, pgcol, bson.M{"GroupID": p.GroupID})

			if delr == 1 {
				response = deleteESDocumentByID(pgindex, p.GroupID)
			} else {
				response = false
			}

		} else {

			delete(productGroup.Products, p.Sku)

			active := false

			m18 := make(map[string]bool)

			for key, value := range productGroup.Products {

				if value.IsMain {
					productGroup.Name = value.Name
					productGroup.Description = value.Description
					productGroup.Images = value.Images
				}

				active = active || value.Active
				addInSet(key, m18)

			}

			prices := computePriceRange(&productGroup)

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

			productGroup.RegularPriceMin = prices[0]
			productGroup.RegularPriceMax = prices[1]
			productGroup.PromotionPriceMin = prices[2]
			productGroup.PromotionPriceMax = prices[3]
			productGroup.Active = active
			productGroup.Skus = toArrayFromSet(m18)

			m19 := make(map[string]bool)
			for _, prd := range productGroup.Products {
				addAllInSet(prd.SearchKeywords, m19)
			}
			productGroup.SearchKeywords = toArrayFromSet(m19)

			m20 := make(map[string]bool)
			for _, prd := range productGroup.Products {
				addInSet(prd.Size, m20)
			}
			productGroup.Sizes = toArrayFromSet(m20)

			m21 := make(map[string]bool)
			for _, prd := range productGroup.Products {
				addInSet(prd.Color, m21)
			}
			productGroup.Colors = toArrayFromSet(m21)

			m22 := make(map[string]bool)
			for _, prd := range productGroup.Products {
				addInSet(prd.Brand, m22)
			}
			productGroup.Brands = toArrayFromSet(m22)

			m23 := make(map[string]bool)
			for _, prd := range productGroup.Products {
				addAllInSet(prd.Category, m23)
			}
			productGroup.Category = toArrayFromSet(m23)

			result := updateMongoDocument(ExternalDB+cidb, pgcol, bson.M{"GroupID": p.GroupID}, bson.M{"$set": productGroup})

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

	cidb := getAccessToken(r)

	pgcol := cidb + ProductGroupExtension

	pcol := cidb + ProductExtension

	pgindex := cidb + ProductGroupExtension + SearchIndexExtension

	response := true

	for _, sku := range skus {

		var opts options.FindOptions

		results := findMongoDocument(ExternalDB+cidb, pcol, bson.M{"Sku": sku}, &opts)

		if len(results) != 1 {
			respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotFound, false)
			return false
		}

		var product PRODUCT

		mapDocument(w, r, &product, results[0])

		result := updateMongoDocument(ExternalDB+cidb, pgcol, bson.M{"GroupID": product.GroupID}, bson.M{"$set": bson.M{"Products." + product.Sku: product}})

		if result[0] == 1 && result[1] == 1 {

			var opts options.FindOptions

			results := findMongoDocument(ExternalDB+cidb, pgcol, bson.M{"GroupID": product.GroupID}, &opts)

			var productGroup PRODUCTGROUP

			mapDocument(w, r, &productGroup, results[0])

			if isPriceUpdate {

				prices := computePriceRange(&productGroup)

				result := updateMongoDocument(ExternalDB+cidb, pgcol, bson.M{"GroupID": product.GroupID}, bson.M{"$set": bson.M{"RegularPriceMin": prices[0], "RegularPriceMax": prices[1], "PromotionPriceMin": prices[2], "PromotionPriceMax": prices[3]}})

				if result[0] == 1 && result[1] == 1 {

					var opts options.FindOptions

					results := findMongoDocument(ExternalDB+cidb, pgcol, bson.M{"GroupID": product.GroupID}, &opts)

					var productGroup PRODUCTGROUP

					mapDocument(w, r, &productGroup, results[0])

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

func computePriceRange(productGroup *PRODUCTGROUP) [4]float64 {

	var result [4]float64

	result[0] = math.MaxFloat64
	result[1] = 0.0
	result[2] = math.MaxFloat64
	result[3] = 0.0

	for _, value := range productGroup.Products {

		if value.RegularPrice < result[0] {
			result[0] = value.RegularPrice
		}
		if value.RegularPrice > result[1] {
			result[1] = value.RegularPrice
		}
		if value.PromotionPrice < result[2] {
			result[2] = value.PromotionPrice
		}
		if value.PromotionPrice > result[3] {
			result[3] = value.PromotionPrice
		}

	}

	return result

}
