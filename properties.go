package main

import "github.com/magiconair/properties"

//MongoURL mongo url for something
var MongoURL string

// MongoPort mongo port
var MongoPort string

//APIPort api port
var APIPort string

//CatalogBasePath base path
var CatalogBasePath string

//APIVersion api version
var APIVersion string

//CatalogPath catalog path
var CatalogPath string

//RedisURL redis url
var RedisURL string

//RedisPort redis port
var RedisPort string

//LoginSuccessMessage login success message
var LoginSuccessMessage string

//LoginFailedMessage login failed message
var LoginFailedMessage string

//ProductAddedMessage product added message
var ProductAddedMessage string

//HTTPBadRequestMessage bad request message
var HTTPBadRequestMessage string

//HTTPInternalServerErrorMessage internal server error message
var HTTPInternalServerErrorMessage string

func loadSystemProperties() {

	p := properties.MustLoadFile(properyFile, properties.UTF8)

	MongoURL = p.GetString("db.mongo.url", "localhost")
	MongoPort = p.GetString("db.mongo.port", "27017")
	APIPort = p.GetString("api.listen.on", "9005")
	CatalogBasePath = p.GetString("api.catalog.base.path", "/catalog/")
	APIVersion = p.GetString("api.version", "v1")
	CatalogPath = CatalogBasePath + APIVersion
	RedisURL = p.GetString("redis.url", "localhost")
	RedisPort = p.GetString("redis.port", "6379")
	LoginSuccessMessage = p.GetString("api.response.message.customer.login.success", "")
	LoginFailedMessage = p.GetString("api.response.message.customer.login.failed", "")
	ProductAddedMessage = p.GetString("api.response.message.product.added", "")
	HTTPBadRequestMessage = p.GetString("api.response.code.message.badRequest", "")
	HTTPInternalServerErrorMessage = p.GetString("api.response.code.message.InternalError", "")

}
