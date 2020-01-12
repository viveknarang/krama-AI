package main

//PGMapping product group indexing configuration
const PGMapping = `
{
	"mappings": {
	  "properties": {
		"Active": {
		  "type": "boolean"
		},
		"Brands": {
		  "type": "text",
		  "fielddata": true
		},
		"Category": {
		  "type": "text"
		},
		"Colors": {
		  "type": "text",
		  "fielddata": true
		},
		"Description": {
		  "type": "text"
		},
		"GroupID": {
		  "type": "keyword"
		},
		"Images": {
		  "type": "text"
		},
		"Name": {
		  "type": "text"
		},
		"Skus": {
		  "type": "keyword"
		},
		"PromotionPriceMax": {
		  "type": "float"
		},
		"PromotionPriceMin": {
		  "type": "float"
		},
		"RegularPriceMax": {
		  "type": "float"
		},
		"RegularPriceMin": {
		  "type": "float"
		},
		"SearchKeywords": {
		  "type": "keyword"
		},
		"Sizes": {
		  "type": "text",
		  "fielddata": true
		},
		"Currency": {
		  "type": "keyword"
		},
		"Updated": {
		  "type": "long"
		}
	  },
	  "dynamic_templates": [
		{
		  "Products_objects": {
			"mapping": {
			  "type": "object",
			  "properties": {
				"Sku": {
				  "type": "keyword"
				},
				"Name": {
				  "type": "text"
				},
				"GroupID": {
				  "type": "keyword"
				},
				"Description": {
				  "type": "text"
				},
				"RegularPrice": {
				  "type": "float"
				},
				"PromotionPrice": {
				  "type": "float"
				},
				"Images": {
				  "type": "text"
				},
				"SearchKeywords": {
				  "type": "text"
				},
				"Quantity": {
				  "type": "long"
				},
				"Category": {
				  "type": "text"
				},
				"Color": {
				  "type": "text"
				},
				"Brand": {
				  "type": "text"
				},
				"Size": {
				  "type": "text"
				},
				"Active": {
				  "type": "boolean"
				},
				"IsMain": {
				  "type": "boolean"
				},
				"Currency": {
				  "type": "keyword"
				},
				"Updated": {
				  "type": "long"
				}
			  }
			},
			"match_mapping_type": "object",
			"path_match":   "Products.*"
		  }
		}
	  ]
	}
  }`
