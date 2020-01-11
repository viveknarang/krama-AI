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

//InvalidSessionMessage invalid session message
var InvalidSessionMessage string

//LoginSessionDuration login session duration
var LoginSessionDuration int64

//JWTSecret JWT Secret
var JWTSecret string

//InternalDB Internal db
var InternalDB string

//ExternalDB External db
var ExternalDB string

//ProductExtension product extension
var ProductExtension string

//ProductGroupExtension product group extension
var ProductGroupExtension string

//CustomersDB customers database
var CustomersDB string

//ProductFoundMessage product found message
var ProductFoundMessage string

//ProductNotFoundMessage product not found message
var ProductNotFoundMessage string

//ProductUpdatedMessage product updated message
var ProductUpdatedMessage string

//ProductNotUpdatedMessage product not updated message
var ProductNotUpdatedMessage string

//ProductDeletedMessage product deleted message
var ProductDeletedMessage string

//ProductGroupDeletedMessage product deleted message
var ProductGroupDeletedMessage string

//ProductGroupFoundMessage product found message
var ProductGroupFoundMessage string

//ProductGroupNotFoundMessage product not found message
var ProductGroupNotFoundMessage string

//ProductNotAddedMessage product not added message
var ProductNotAddedMessage string

//ProductNotDeletedMessage product not deleted message
var ProductNotDeletedMessage string

const properyFile = "/home/narang/work/src/github.com/viveknarang/kramaAPI/api.properties"

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
	InvalidSessionMessage = p.GetString("api.response.message.login.invalidSession", "")
	LoginSessionDuration = p.GetInt64("api.login.session.duration", 80000)
	JWTSecret = p.GetString("jwt.secret", "")
	InternalDB = p.GetString("db.mongo.internal", "Internal")
	ExternalDB = p.GetString("db.mongo.external", "External")
	ProductExtension = p.GetString("db.mongo.external.product.extension", "")
	ProductGroupExtension = p.GetString("db.mongo.external.productgroup.extension", "")
	CustomersDB = p.GetString("db.mongo.internal.customers.collection", "")
	ProductFoundMessage = p.GetString("api.response.message.product.found", "")
	ProductNotFoundMessage = p.GetString("api.response.message.product.notfound", "")
	ProductUpdatedMessage = p.GetString("api.response.message.product.updated", "")
	ProductNotUpdatedMessage = p.GetString("api.response.message.product.notupdated", "")
	ProductDeletedMessage = p.GetString("api.response.message.product.deleted", "")
	ProductGroupDeletedMessage = p.GetString("api.response.message.productgroup.deleted", "")
	ProductGroupFoundMessage = p.GetString("api.response.message.productgroup.found", "")
	ProductGroupNotFoundMessage = p.GetString("api.response.message.productgroup.notfound", "")
	ProductNotAddedMessage = p.GetString("api.response.message.product.notadded", "")
	ProductNotDeletedMessage = p.GetString("api.response.message.product.notdeleted", "")

}
