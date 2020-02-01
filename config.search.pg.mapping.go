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
		  "type": "keyword"
		},
		"Category": {
		  "type": "text"
		},
		"Colors": {
		  "type": "keyword"
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
		  "type": "keyword"
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
			"Attributes_text": {
				"path_match":   "Attributes.*",
				"match_mapping_type": "string",
				"mapping": {
				  "type": "keyword"
				}
			}
		},
		{
		  "Attributes_long": {
			  "path_match":   "Attributes.*",
			  "match_mapping_type": "long",
			  "mapping": {
				"type": "long"
			  }
		  }
		},
		{
		  "Attributes_boolean": {
			  "path_match":   "Attributes.*",
			  "match_mapping_type": "boolean",
			  "mapping": {
				"type": "boolean"
			  }
		  }
		},
		{
		  "Attributes_date": {
			  "path_match":   "Attributes.*",
			  "match_mapping_type": "date",
			  "mapping": {
				"type": "date"
			  }
		  }
		},
		{
		  "Attributes_double": {
			  "path_match":   "Attributes.*",
			  "match_mapping_type": "double",
			  "mapping": {
				"type": "double"
			  }
		  }
		}
	  ]
	}
  }`
