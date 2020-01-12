package main

//PGMapping product group indexing configuration
const PGMapping = `
{
	"mappings": {
	  "properties": {
		"active": {
		  "type": "boolean"
		},
		"brands": {
		  "type": "text",
		  "fielddata": true
		},
		"category": {
		  "type": "text"
		},
		"colors": {
		  "type": "text",
		  "fielddata": true
		},
		"description": {
		  "type": "text"
		},
		"groupID": {
		  "type": "keyword"
		},
		"images": {
		  "type": "text"
		},
		"name": {
		  "type": "text"
		},
		"skus": {
		  "type": "keyword"
		},
		"products": {
		  "type": "object"
		},
		"promotionPriceMax": {
		  "type": "float"
		},
		"promotionPriceMin": {
		  "type": "float"
		},
		"regularPriceMax": {
		  "type": "float"
		},
		"regularPriceMin": {
		  "type": "float"
		},
		"searchKeywords": {
		  "type": "keyword"
		},
		"sizes": {
		  "type": "text",
		  "fielddata": true
		}
	  }
	}
  }`
