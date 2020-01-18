package main

import (
	"github.com/magiconair/properties"
	"github.com/romana/rlog"
)

//MongoURL mongo url for something
var MongoURL string

// MongoPort mongo port
var MongoPort string

//APIPort api port
var APIPort string

//CatalogBasePath base path
var CatalogBasePath string

//OrdersBasePath base path
var OrdersBasePath string

//APIVersion api version
var APIVersion string

//CatalogPath catalog path
var CatalogPath string

//OrdersPath catalog path
var OrdersPath string

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

//OrdersExtension orders extension
var OrdersExtension string

//OrderCreatedMessage order created message
var OrderCreatedMessage string

//OrderUpdatedMessage order updated message
var OrderUpdatedMessage string

//OrderDeletedMessage order deleted message
var OrderDeletedMessage string

//OrderNotFoundMessage order not found message
var OrderNotFoundMessage string

//OrderFoundMessage order not found message
var OrderFoundMessage string

//OrderNotUpdatedMessage order updated message
var OrderNotUpdatedMessage string

//ElasticURL elastic URL
var ElasticURL string

//ElasticPort elastic port
var ElasticPort string

//SearchIndexExtension search index extension
var SearchIndexExtension string

//SearchBasePath search base path
var SearchBasePath string

//SearchPath search api path
var SearchPath string

//MissingAccessToken missing access token
var MissingAccessToken string

//MissingContentType missing content type
var MissingContentType string

//ServiceDownMessage api service down message
var ServiceDownMessage string

//CustomersBasePath customers base path
var CustomersBasePath string

//CustomersPath customers path
var CustomersPath string

//CustomersAddedMessage product found message
var CustomersAddedMessage string

//CustomersFoundMessage product found message
var CustomersFoundMessage string

//CustomersNotFoundMessage product not found message
var CustomersNotFoundMessage string

//CustomersUpdatedMessage product updated message
var CustomersUpdatedMessage string

//CustomersNotUpdatedMessage product not updated message
var CustomersNotUpdatedMessage string

//CustomersDeletedMessage product deleted message
var CustomersDeletedMessage string

//CustomersNotAddedMessage product not added message
var CustomersNotAddedMessage string

//CustomersNotDeletedMessage product not deleted message
var CustomersNotDeletedMessage string

//CustomersCollectionExtension customers collection extension
var CustomersCollectionExtension string

//ProductAlreadyExistsMessage product already exists message
var ProductAlreadyExistsMessage string

//CustomerAlreadyExistsMessage customer already exists message
var CustomerAlreadyExistsMessage string

//ShoppingCartBasePath shopping cart base path
var ShoppingCartBasePath string

//ShoppingCartPath shopping cart path
var ShoppingCartPath string

//ShoppingCartLife shopping cart life duration
var ShoppingCartLife string

func loadSystemProperties() bool {

	rlog.Debug("loadSystemProperties() handle function invoked ...")

	if !fileExists(PROPERTYFILE) {

		rlog.Error("loadSystemProperties() Error property file: " + PROPERTYFILE + " does not exist...")
		return false

	} else {

		rlog.Debug("loadSystemProperties() loading property file: " + PROPERTYFILE)

	}

	p := properties.MustLoadFile(PROPERTYFILE, properties.UTF8)

	MongoURL = p.GetString("db.mongo.url", "localhost")
	MongoPort = p.GetString("db.mongo.port", "27017")
	APIPort = p.GetString("api.listen.on", "9005")
	CatalogBasePath = p.GetString("api.catalog.base.path", "/catalog/")
	OrdersBasePath = p.GetString("api.orders.base.path", "/orders/")
	SearchBasePath = p.GetString("pi.search.base.path", "/search/")
	CustomersBasePath = p.GetString("api.customers.base.path", "/customers/")
	ShoppingCartBasePath = p.GetString("api.shoppingcart.base.path", "/shoppingcart/")

	APIVersion = p.GetString("api.version", "v1")
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
	OrdersExtension = p.GetString("db.mongo.external.orders.extension", "")
	OrderCreatedMessage = p.GetString("api.response.message.orders.ordercreated", "")
	OrderUpdatedMessage = p.GetString("api.response.message.orders.orderupdated", "")
	OrderNotUpdatedMessage = p.GetString("api.response.message.orders.ordernotupdated", "")
	OrderDeletedMessage = p.GetString("api.response.message.orders.orderdeleted", "")
	OrderFoundMessage = p.GetString("api.response.message.orders.orderfound", "")
	OrderNotFoundMessage = p.GetString("api.response.message.orders.ordernotfound", "")
	ElasticURL = p.GetString("search.elastic.url", "")
	ElasticPort = p.GetString("search.elastic.port", "")
	SearchIndexExtension = p.GetString("search.elastic.index.extension", "")
	MissingAccessToken = p.GetString("api.response.message.exception.missingtoken", "")
	MissingContentType = p.GetString("api.response.message.exception.contentTypeMissing", "")
	ServiceDownMessage = p.GetString("api.response.message.exception.serviceUnavailable", "")
	CustomersAddedMessage = p.GetString("api.response.message.customers.added", "")
	CustomersFoundMessage = p.GetString("api.response.message.customers.found", "")
	CustomersNotFoundMessage = p.GetString("api.response.message.customers.notfound", "")
	CustomersUpdatedMessage = p.GetString("api.response.message.customers.updated", "")
	CustomersNotUpdatedMessage = p.GetString("api.response.message.customers.notupdated", "")
	CustomersDeletedMessage = p.GetString("api.response.message.customers.deleted", "")
	CustomersNotAddedMessage = p.GetString("api.response.message.customers.notadded", "")
	CustomersNotDeletedMessage = p.GetString("api.response.message.customers.notdeleted", "")
	CustomersCollectionExtension = p.GetString("db.mongo.external.customers.extension", "")
	ProductAlreadyExistsMessage = p.GetString("api.response.message.product.exists", "")
	CustomerAlreadyExistsMessage = p.GetString("api.response.message.customers.exists", "")
	ShoppingCartLife = p.GetString("api.shoppingcart.life", "")

	CatalogPath = CatalogBasePath + APIVersion
	OrdersPath = OrdersBasePath + APIVersion
	SearchPath = SearchBasePath + APIVersion
	CustomersPath = CustomersBasePath + APIVersion
	ShoppingCartPath = ShoppingCartBasePath + APIVersion

	return true

}
