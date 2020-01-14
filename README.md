
[![Build Status](https://travis-ci.org/viveknarang/kramaAPI.svg?branch=master)](https://travis-ci.org/viveknarang/kramaAPI)

# Krama AI - A blazing fast E-commerce AI platform

## Introduction

This API provides you with several options to maintain your product catalog and search products in your catalog. 

With this API, you can create, update and delete products in your product catalog. The API automatically manages the product groups for you. The product group is identified by the groupID field (i.e. all the products with the same groupID are combined in a single product group). While the API allows you to get and delete product groups by groupID, the API does not allow you to directly modify the product groups. The only way to modify product groups is to use API endpoints for individual products. With this approach, the API tries to ensure that the data is not corrupted. When you delete a product group, all the products in the group are also automatically deleted. 

<aside class="notice">
It is important to understand the concept of product groups. For search quality and other advanced features provided by our platform, products are grouped to form product groups. These product groups are essentially a product with different variations. Example: A shirt can be of multiple sizes and/or colors. So all of the variations of this shirt are grouped to form a product group.  
</aside>

The search API is fairly powerful too. It allows features like search on a specific field or a set of fields. The search API responds with product groups where the query matches certain fields. Search API also allows you to select the standard facets to be included in the API response. In addtion to the features mentioned above, the API automatically syncs the search index with changes in products/product groups, **in real-time**. Also, for efficiency and speed, the search and GET product/productgroup endpoints are cached. Upon any updates, the cache is updated as well. 

We are continuously adding new features and improving this API, if you have any suggestions please reach out to us [here](mailto:vivek.narang10@gmail.com)

API Powered by

- Golang            
- Redis             
- Elasticsearch     
- MongoDB           



<aside class="success">
The current API version is: v1 Please replace {API version} with v1 in your API calls
</aside>

# API login

## Get API access token

> Sample HTTP request body:

```json
{
  "CustomerID": "6476154099",
  "APIKey": "zaCELgL.0imfnc8mVLWwsAawjYr4Rx-Af50DDqtlx"
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Login Successful ...",
    "Time": 1579026954047130825,
    "Response": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjeHMiOiJDb250ZUFtZXJpY2EiLCJleHAiOjE1NzkxMDY5NTQsImlhdCI6MTU3OTAyNjk1NCwibmJmIjoxNTc5MDI2ODU0LCJ1aWQiOiIwMjQ0Zjg1NS1jMWQ3LTQyNGYtOWI5OS04NGZmYWNiYzYwOGUifQ.6IhX3X321NlZFtSSf3JUPisD7fTxqeVrCpHQ6WDDgIk",
        "validForSeconds": 80000
    }
}
```

> Sample invalid API response (valid token expired or invalid token used):

```json
{
    "Code": 401,
    "Success": false,
    "Message": "You need to either login or your access token is either expired, invalid, or corrupt ...",
    "Time": 1579029101079922462,
    "Response": null
}
```

> Sample invalid API response:

```json
{
    "Code": 401,
    "Success": false,
    "Message": "Login Failed! Please check your credentials and also make sure that you are an active customer ...",
    "Time": 1579027031268672037,
    "Response": null
}
```

This endpoint gets you your API access token. You need to send your customer ID and the API key that we provided you for using our platform. Upon receiving your valid credentials, the API will respond with a token with additional information including the validFor key which tells you how long this access token is valid for. Please set **x-access-token** to the value of the **token**, in the header of your subsequent API calls. 


### HTTP Request URL

`GET https://api.krama.ai/customers/{API version}/login`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Content-Type       | application/json                             |

### HTTP Request Body Parameters

| Parameter         | Description                     |
|-------------------| --------------------------------|
| CustomerID        | Your customer ID provided by us |
| APIKey            | The API key that is sent by us  |

<aside class="warning">
You do not need to invoke login too often. Please include the token that you receive upon a successful login in your subsequent API calls, until the token expires.
</aside>

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |
| token             | Included in response object that should be included in subsequent API calls   |
| validForSeconds   | Included in response object that tells the validity of access token in seconds|


<aside class="notice">
With the field validForSeconds in response, you can calculate the time after with your servers need to login again to get a new token.
</aside>


# Catalog API

## Add a new product

> Sample HTTP request body:

```json
{
  "Sku": "B07K3BHGL3",
  "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
  "Images": [
    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
  ],
  "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
  "GroupID": "MSLAPS2",
  "SearchKeywords": [
    "Laptop",
    "Microsoft",
    "Surface"
  ],
  "RegularPrice": 2799,
  "PromotionPrice": 2600,
  "Currency": "CDN",
  "IsMain": true,
  "Quantity": 200,
  "Size": "13.5 inches",
  "Brand": "Microsoft",
  "Color": "Black",
  "Category": [
    "Computers & Tablets>Laptops"
  ],
  "Active": true,
  "Attributes": {
    "Display Size": "13.5 inches",
    "RAM": "16 GB",
    "Memory Speed": "1 GHz",
    "Wireless Standard": "802.11ac",
    "Number of USB 2.0 Ports": "1",
    "Series": "Surface Laptop 2",
    "Item model number": "DAL-00092",
    "Operating System": "Windows 10 Home",
    "Item Weight": "1.28 Kg",
    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
    "Color": "Black",
    "Processor Count": "16",
    "Flash Memory Size": "512.00",
    "Batteries": "1",
    "ASIN": "B07K3BHGL3",
    "Shipping Weight": "2.2 kg",
    "Date First Available": "Nov. 4 2018"
  }
}
```

> Sample valid API response:

```json
{
    "Code": 201,
    "Success": true,
    "Message": "Product Added ...",
    "Time": 1579028967992159981,
    "Response": {
        "Sku": "B07K3BHGL3",
        "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
        "GroupID": "MSLAPS2",
        "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
        "RegularPrice": 2799,
        "PromotionPrice": 2600,
        "Images": [
            "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
            "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
        ],
        "SearchKeywords": [
            "Laptop",
            "Microsoft",
            "Surface"
        ],
        "Quantity": 200,
        "Category": [
            "Computers & Tablets>Laptops"
        ],
        "Color": "Black",
        "Brand": "Microsoft",
        "Size": "13.5 inches",
        "Active": true,
        "Attributes": {
            "ASIN": "B07K3BHGL3",
            "Batteries": "1",
            "Color": "Black",
            "Date First Available": "Nov. 4 2018",
            "Display Size": "13.5 inches",
            "Flash Memory Size": "512.00",
            "Item Weight": "1.28 Kg",
            "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
            "Item model number": "DAL-00092",
            "Memory Speed": "1 GHz",
            "Number of USB 2.0 Ports": "1",
            "Operating System": "Windows 10 Home",
            "Processor Count": "16",
            "RAM": "16 GB",
            "Series": "Surface Laptop 2",
            "Shipping Weight": "2.2 kg",
            "Wireless Standard": "802.11ac"
        },
        "IsMain": true,
        "Currency": "CDN",
        "Updated": 1579028967412667337
    }
}
```

Use this API endpoint to add a new product in the products collection. When a product is added in the products collection, this product is also added in product group collection. If the product group with the matching groupID is missing, a new product group is formed. Search index and the cache are also automatically updated with a valid call to this endpoint. 


### HTTP Request URL

`POST https://api.krama.ai/catalog/{API version}/products`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |
|Content-Type       | application/json                             |

### HTTP Request Body Parameters

|    Parameter   |          Constraints         |        Description                                           |
|----------------|------------------------------|--------------------------------------------------------------|
|Sku             |   String, Max 50 Characters  |  The SKU of the product                                      |
|Name            |   Text, Max 100 Characters   |  The name of the product                                     |
|Description     |   Text, Max 10240 Characters |  The description of the product                              | 
|GroupID         |   String, Max 50 Characters  |  The product group ID                                        | 
|RegularPrice    |   Float, Greater than 0      |  Everyday price                                              | 
|PromotionPrice  |   Float, Greater than 0      |  On-sale price                                               | 
|Images          |   Valid URL, Mandatory       |  Product images                                              | 
|SearchKeywords  |   Text[], Mandatory          |  Keywords that you want this product to be searched with     |
|Quantity        |   Integer, Greater than 0    |  Inventory stock quantity                                    | 
|Category        |   Text[], Mandatory          |  Category breadcrumbs array                                  | 
|Color           |   Text, Optional             |  Product color                                               |
|Brand           |   Text, Optional             |  Product brand                                               | 
|Size            |   Text, Optional             |  Product size                                                | 
|Active          |   Boolean, Mandatory         |  Flag to set the product availble for sale                   |    
|IsMain          |   Boolean, Mandatory         |  Is the product main product in the product group?           |
|Attributes      |   Details in another section |  Additional field to pass in misc. product attributes        |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |




## Get a specific product

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": false,
    "Message": "Product Found ...",
    "Time": 1579029290287576701,
    "Response": {
        "Sku": "B07K3BHGL3",
        "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
        "GroupID": "MSLAPS2",
        "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
        "RegularPrice": 2799,
        "PromotionPrice": 2600,
        "Images": [
            "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
            "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
        ],
        "SearchKeywords": [
            "Laptop",
            "Microsoft",
            "Surface"
        ],
        "Quantity": 200,
        "Category": [
            "Computers & Tablets>Laptops"
        ],
        "Color": "Black",
        "Brand": "Microsoft",
        "Size": "13.5 inches",
        "Active": true,
        "Attributes": {
            "ASIN": "B07K3BHGL3",
            "Batteries": "1",
            "Color": "Black",
            "Date First Available": "Nov. 4 2018",
            "Display Size": "13.5 inches",
            "Flash Memory Size": "512.00",
            "Item Weight": "1.28 Kg",
            "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
            "Item model number": "DAL-00092",
            "Memory Speed": "1 GHz",
            "Number of USB 2.0 Ports": "1",
            "Operating System": "Windows 10 Home",
            "Processor Count": "16",
            "RAM": "16 GB",
            "Series": "Surface Laptop 2",
            "Shipping Weight": "2.2 kg",
            "Wireless Standard": "802.11ac"
        },
        "IsMain": true,
        "Currency": "CDN",
        "Updated": 1579028967412667337
    }
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Product Not Found ...",
    "Time": 1579029354292991837,
    "Response": null
}
```

When you want to get a specific product you can use this endpoint. All you need to pass is your access token and the SKU. This endpoint is cached for efficiency but also ensures that updated product data is served when applicable. 


### HTTP Request URL

`GET https://api.krama.ai/catalog/{API version}/products/{SKU}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |

### HTTP Request URL Parameters

| Parameter    | Description         |
|--------------|---------------------|
|SKU           | The product SKU     |

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |
| Sku               | The SKU of the product                                                        |
| Name              | The name of the product                                                       |
| Description       | The description of the product                                                |   
| GroupID           | The product group ID                                                          |
| RegularPrice      | Everyday price                                                                |
| PromotionPrice    | On-sale price                                                                 |
| Images            | Product images                                                                |
| SearchKeywords    | Keywords that you want this product to be searched with                       |
| Quantity          | Inventory stock quantity                                                      |
| Category          | Category breadcrumbs                                                          |
| Color             | Product color                                                                 |
| Brand             | Product brand                                                                 |
| Size              | Product size                                                                  |
| Active            | Is product available for sale?                                                |
| IsMain            | Is the product main product in the group?                                     |
| Attributes        | Field to define additional product attributes                                 |




## Update a product

> Sample HTTP request body:

```json
{
  "Sku": "B07K3BHGL3",
  "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
  "Images": [
    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
  ],
  "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
  "GroupID": "MSLAPS2",
  "SearchKeywords": [
    "Laptop",
    "Microsoft",
    "Surface"
  ],
  "RegularPrice": 2999,
  "PromotionPrice": 2600,
  "Currency": "CDN",
  "IsMain": true,
  "Quantity": 200,
  "Size": "13.5 inches",
  "Brand": "Microsoft",
  "Color": "Black",
  "Category": [
    "Computers & Tablets>Laptops"
  ],
  "Active": true,
  "Attributes": {
    "Display Size": "13.5 inches",
    "RAM": "16 GB",
    "Memory Speed": "1 GHz",
    "Wireless Standard": "802.11ac",
    "Number of USB 2.0 Ports": "1",
    "Series": "Surface Laptop 2",
    "Item model number": "DAL-00092",
    "Operating System": "Windows 10 Home",
    "Item Weight": "1.28 Kg",
    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
    "Color": "Black",
    "Processor Count": "16",
    "Flash Memory Size": "512.00",
    "Batteries": "1",
    "ASIN": "B07K3BHGL3",
    "Shipping Weight": "2.2 kg",
    "Date First Available": "Nov. 4 2018"
  }
}
```

> Sample valid API response:

```json
{
    "Code": 202,
    "Success": true,
    "Message": "Product Updated ...",
    "Time": 1579029475468110631,
    "Response": {
        "Sku": "B07K3BHGL3",
        "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
        "GroupID": "MSLAPS2",
        "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
        "RegularPrice": 2999,
        "PromotionPrice": 2600,
        "Images": [
            "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
            "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
        ],
        "SearchKeywords": [
            "Laptop",
            "Microsoft",
            "Surface"
        ],
        "Quantity": 200,
        "Category": [
            "Computers & Tablets>Laptops"
        ],
        "Color": "Black",
        "Brand": "Microsoft",
        "Size": "13.5 inches",
        "Active": true,
        "Attributes": {
            "ASIN": "B07K3BHGL3",
            "Batteries": "1",
            "Color": "Black",
            "Date First Available": "Nov. 4 2018",
            "Display Size": "13.5 inches",
            "Flash Memory Size": "512.00",
            "Item Weight": "1.28 Kg",
            "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
            "Item model number": "DAL-00092",
            "Memory Speed": "1 GHz",
            "Number of USB 2.0 Ports": "1",
            "Operating System": "Windows 10 Home",
            "Processor Count": "16",
            "RAM": "16 GB",
            "Series": "Surface Laptop 2",
            "Shipping Weight": "2.2 kg",
            "Wireless Standard": "802.11ac"
        },
        "IsMain": true,
        "Currency": "CDN",
        "Updated": 0
    }
}
```

> Sample invalid API response:

```json
{
    "Code": 304,
    "Success": false,
    "Message": "Product Not Found ...",
    "Time": 1579029677035327895,
    "Response": null
}
```

Use this API endpoint to update your product information in the catalog. For now you need to pass the entire product object with updated parts (this functionality will be improved very soon). When you hit this endpoint, the data in the products collection gets updated, product group data also gets updated automatically, search index is also updated and the cache entry is removed first and updated on the next GET call. 


### HTTP Request URL

`PUT https://api.krama.ai/catalog/{API version}/products/{SKU}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |
|Content-Type       | application/json                             |


### HTTP Body Parameters

|    Parameter   |          Constraints         |        Description                                           |
|----------------|------------------------------|--------------------------------------------------------------|
|Sku             |   String, Max 50 Characters  |  The SKU of the product                                      |
|Name            |   Text, Max 100 Characters   |  The name of the product                                     |
|Description     |   Text, Max 10240 Characters |  The description of the product                              | 
|GroupID         |   String, Max 50 Characters  |  The product group ID                                        | 
|RegularPrice    |   Float, Greater than 0      |  Everyday price                                              | 
|PromotionPrice  |   Float, Greater than 0      |  On-sale price                                               | 
|Images          |   Valid URL, Mandatory       |  Product images                                              | 
|SearchKeywords  |   Text[], Mandatory          |  Keywords that you want this product to be searched with     |
|Quantity        |   Integer, Greater than 0    |  Inventory stock quantity                                    | 
|Category        |   Text[], Mandatory          |  Category breadcrumbs array                                  | 
|Color           |   Text, Optional             |  Product color                                               |
|Brand           |   Text, Optional             |  Product brand                                               | 
|Size            |   Text, Optional             |  Product size                                                | 
|Active          |   Boolean, Mandatory         |  Flag to set the product availble for sale                   |    
|IsMain          |   Boolean, Mandatory         |  Is the product main product in the product group?           |
|Attributes      |   Details in another section |  Additional field to pass in misc. product attributes        |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |
| Sku               | The SKU of the product                                                        |
| Name              | The name of the product                                                       |
| Description       | The description of the product                                                |   
| GroupID           | The product group ID                                                          |
| RegularPrice      | Everyday price                                                                |
| PromotionPrice    | On-sale price                                                                 |
| Images            | Product images                                                                |
| SearchKeywords    | Keywords that you want this product to be searched with                       |
| Quantity          | Inventory stock quantity                                                      |
| Category          | Category breadcrumbs                                                          |
| Color             | Product color                                                                 |
| Brand             | Product brand                                                                 |
| Size              | Product size                                                                  |
| Active            | Is product available for sale?                                                |
| IsMain            | Is the product main product in the group?                                     |
| Attributes        | Field to define additional product attributes                                 |




## Update product price (Bulk)

> Sample HTTP request body:

```json
{
  "Prices": {
    "B07K3BHGL3": {
      "RegularPrice": 2700,
      "PromotionPrice": 2500
    },
    "B07K3BHGL4": {
      "RegularPrice": 99,
      "PromotionPrice": 45
    }
  }
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Prices Updated ...",
    "Time": 1579029846231436883,
    "Response": {
        "Products Not Found": [
            "B07K3BHGL4"
        ],
        "Products Not Updated": null,
        "Products Updated": [
            "B07K3BHGL3"
        ]
    }
}
```

Use this API endpoint to update your product prices in the catalog. You can use this endpoint to submit a map of skus of prices. Each sku in the map
is associated with the regular and the promotion prices. The new prices cannot not be negative values. As with other features the API takes care of 
ensuring that the product groups and the search index are synced as well. The response object contains three lists - the list of updated skus 
(essentially the skus for which the prices were updated), a list of skus which were not updated (most likely that you submitted same old prices), and
finally a list of skus that were not found. You can use this information to check if your update request was executed as per your expectations. 


### HTTP Request URL

`PUT https://api.krama.ai/catalog/{API version}/products/price/update`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |
|Content-Type       | application/json                             |


### HTTP Body Parameters

|    Parameter   |          Constraints         |        Description                                           |
|----------------|------------------------------|--------------------------------------------------------------|
|Prices          |                              |  The prices map that contains the sku-price maps             |
|PromotionPrice  |   Float, non-negative        |  The promotion price of the mapped sku                       |
|RegularPrice    |   Float, non-negative        |  The regular price associated with the sku                   | 


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |





## Update product inventory (Bulk)

> Sample HTTP request body:

```json
{
  "Quantity": {
    "B07K3BHGL3": 500,
    "B07K3BHGL4": 341
  }
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Inventory Updated ...",
    "Time": 1579030006188444624,
    "Response": {
        "Products Not Found": [
            "B07K3BHGL4"
        ],
        "Products Not Updated": null,
        "Products Updated": [
            "B07K3BHGL3"
        ]
    }
}
```

Use this API endpoint to update your product quantity. You can use this endpoint to updates inventory for multiple products. Please find the API request details below. The API updates product groups as well as search index. The response object contains three lists - a list that gives you the skus of those
products where the product quantities were updated. The response object also provides you with list of those skus where there was no change and also
those skus which were not found in the catalog. 


### HTTP Request URL

`PUT https://api.krama.ai/catalog/{API version}/products/inventory/update`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |
|Content-Type       | application/json                             |


### HTTP Body Parameters

|    Parameter    |          Constraints                    |        Description                                           |
|-----------------|-----------------------------------------|--------------------------------------------------------------|
|Quantity         |                                         |  The map containing skus-quantity mappings                   |
|{sku : quantity} | {unique identifier, float non-negative} |  sku-quantity mappings in the Quantity object                |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |



## Delete a product

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Product Deleted ...",
    "Time": 1578886136908917901,
    "Response": null
}
```

> If the SKU does not exist, the response will look like the following:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Product Not Found ...",
    "Time": 1578886152262487961,
    "Response": null
}
```

Use this API endpoint to remove a product from the catalog. When you hit this endpoint with a valid request, the product in the products collection gets removed, the productgroups collection is also automatically updated and the cache and search index is also updated. If there was only one product in the product group the product group object is also removed from the productgroups collection. 
 

### HTTP Request URL

`DELETE https://api.krama.ai/catalog/{API version}/products/{SKU}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |

### HTTP Request URL Parameters

| Parameter    | Description         |
|--------------|---------------------|
|SKU           | The product SKU     |

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |


## Get a product group

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Product Group Found ...",
    "Time": 1579030135545676974,
    "Response": {
        "GroupID": "MSLAPS2",
        "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
        "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
        "RegularPriceMin": 2700,
        "RegularPriceMax": 2700,
        "PromotionPriceMin": 2500,
        "PromotionPriceMax": 2500,
        "Skus": [
            "B07K3BHGL3"
        ],
        "Images": [
            "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
            "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
        ],
        "SearchKeywords": [
            "Laptop",
            "Microsoft",
            "Surface"
        ],
        "Category": [
            "Computers & Tablets>Laptops"
        ],
        "Colors": [
            "Black"
        ],
        "Brands": [
            "Microsoft"
        ],
        "Sizes": [
            "13.5 inches"
        ],
        "Active": true,
        "Currency": "CDN",
        "Updated": 1579028967414440899,
        "Products": {
            "B07K3BHGL3": {
                "Sku": "B07K3BHGL3",
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                "GroupID": "MSLAPS2",
                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                "RegularPrice": 2700,
                "PromotionPrice": 2500,
                "Images": [
                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                ],
                "SearchKeywords": [
                    "Laptop",
                    "Microsoft",
                    "Surface"
                ],
                "Quantity": 500,
                "Category": [
                    "Computers & Tablets>Laptops"
                ],
                "Color": "Black",
                "Brand": "Microsoft",
                "Size": "13.5 inches",
                "Active": true,
                "Attributes": {
                    "ASIN": "B07K3BHGL3",
                    "Batteries": "1",
                    "Color": "Black",
                    "Date First Available": "Nov. 4 2018",
                    "Display Size": "13.5 inches",
                    "Flash Memory Size": "512.00",
                    "Item Weight": "1.28 Kg",
                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                    "Item model number": "DAL-00092",
                    "Memory Speed": "1 GHz",
                    "Number of USB 2.0 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "RAM": "16 GB",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac"
                },
                "IsMain": true,
                "Currency": "CDN",
                "Updated": 0
            }
        },
        "Attributes": {
            "ASIN": [
                "B07K3BHGL3"
            ],
            "Batteries": [
                "1"
            ],
            "Color": [
                "Black"
            ],
            "Date First Available": [
                "Nov. 4 2018"
            ],
            "Display Size": [
                "13.5 inches"
            ],
            "Flash Memory Size": [
                "512.00"
            ],
            "Item Weight": [
                "1.28 Kg"
            ],
            "Item dimensions L x W x H": [
                "17.8 x 12.7 x 15.2 cm"
            ],
            "Item model number": [
                "DAL-00092"
            ],
            "Memory Speed": [
                "1 GHz"
            ],
            "Number of USB 2.0 Ports": [
                "1"
            ],
            "Operating System": [
                "Windows 10 Home"
            ],
            "Processor Count": [
                "16"
            ],
            "RAM": [
                "16 GB"
            ],
            "Series": [
                "Surface Laptop 2"
            ],
            "Shipping Weight": [
                "2.2 kg"
            ],
            "Wireless Standard": [
                "802.11ac"
            ]
        }
    }
}
```

> Sample valid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Product Group Not Found ...",
    "Time": 1578886448802758717,
    "Response": null
}
```


This API endpoint gets a specific product group by product group ID. This endpoint is cached for efficiency and speed. 


### HTTP Request URL

`GET https://api.krama.ai/catalog/{API version}/productgroups/{PGID}`

### HTTP Requesr Header

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |

### HTTP Request URL Parameters

| Parameter     | Description          |
|---------------|----------------------|
|PGID           | The product Group ID |

### HTTP Response

|  Key                  |    Description                                                                |
|-----------------------|-------------------------------------------------------------------------------|
| Code                  | Response code for the request                                                 |
| Success               | Flag that tells if the request was successful                                 |
| Message               | Message for additional information                                            |
| Time                  | Unix timestamp of the response                                                |
| Response              | Response object containing response information                               |
| Skus                  |  A list of SKUs of all the products in the group                              |
| Colors                |  A list of all the colors from all the products in the group                  |
| Brands                |  A list of all the brands from all the products in the group                  |
| Sizes                 |  A list of all the sizes from all the products in the group                   |
| Images                |  A list of all the images from the main product in the group                  |
| SearchKeywords        |  A list of all the searchKeywords from all the products in the group          |
| Category              |  Category from the main product in the group                                  |
| GroupID               |  Product group ID to uniquely identify this product group                     |
| Name                  |  Name from the main product in the group                                      |
| Description           |  Product description                                                          |
| RegularPriceMin       |  Minimum regular price computed in the group                                  |
| RegularPriceMax       |  Maximum regular price computed in the group                                  |
| PromotionPriceMin     |  Minimum promotion price computed in the group                                |
| PromotionPriceMax     |  Maximum promotion price computed in the group                                |
| Active                |  Active flag to indicate if the product is available for sale                 |
| Products              |  List of all the product objects for reference                                |
| Attributes            |  Field that maps additional attributes to the product group                   |



## Delete a product group

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Product Group Deleted ...",
    "Time": 1578886773436347784,
    "Response": null
}
```

> If the PGID does not exist, the response will look like the following:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Product Group Not Found ...",
    "Time": 1578886787090095342,
    "Response": null
}
```

Use this API endpoint to remove a product group from the productgroups collection in the database. This call also updates the search index. When a product group is deleted, entries of related products in the products collection are also removed. 
 

### HTTP Request URL

`DELETE https://api.krama.ai/catalog/{API version}/productgroups/{PGID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |

### HTTP Request URL Parameters

| Parameter     | Description          |
|---------------|----------------------|
|PGID           | The product Group ID |

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |



# Order API

## Create an order


> Sample HTTP request body:

```json
{
  "OrderID": "1234567",
  "CustomerID": "55348",
  "PaymentStatus": "PAYMENT_PENDING",
  "PaymentAmount": 2600,
  "Currency": "CDN",
  "OrderStatus": "ORDERED",
  "ShippingAddress": {
    "FirstName": "Vivek",
    "LastName": "Narang",
    "AddressLineOne": "111 Edgar Ave",
    "AddressLineTwo": "",
    "City": "Richmond Hill",
    "State": "Ontario",
    "Country": "Canada",
    "Pincode": "L4C 6K3"
  },
  "Products": {
    "B07K3BHGL3": {
      "Sku": "B07K3BHGL3",
      "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
      "GroupID": "MSLAPS2",
      "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
      "RegularPrice": 2999,
      "PromotionPrice": 2600,
      "Images": [
        "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
        "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
      ],
      "SearchKeywords": [
        "Laptop",
        "Microsoft",
        "Surface"
      ],
      "Quantity": 200,
      "Category": [
        "Computers & Tablets>Laptops"
      ],
      "Color": "Black",
      "Brand": "Microsoft",
      "Size": "13.5 inches",
      "Active": true,
      "Attributes": {
        "ASIN": "B07K3BHGL3",
        "Batteries": "1",
        "Color": "Black",
        "Date First Available": "Nov. 4 2018",
        "Display Size": "13.5 inches",
        "Flash Memory Size": "512.00",
        "Item Weight": "1.28 Kg",
        "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
        "Item model number": "DAL-00092",
        "Memory Speed": "1 GHz",
        "Number of USB 2.0 Ports": "1",
        "Operating System": "Windows 10 Home",
        "Processor Count": "16",
        "RAM": "16 GB",
        "Series": "Surface Laptop 2",
        "Shipping Weight": "2.2 kg",
        "Wireless Standard": "802.11ac"
      },
      "IsMain": true,
      "Currency": "CDN",
      "Updated": 0
    }
  },
  "ProductQuantity": {
    "B07K3BHGL3": 1
  }
}
```


> Sample valid API response:

```json
{
    "Code": 201,
    "Success": true,
    "Message": "Order Created ...",
    "Time": 1579030377363245763,
    "Response": {
        "OrderID": "1234567",
        "OrderCreationDate": 1579030377359016931,
        "OrderUpdateDate": 0,
        "CustomerID": "55348",
        "Products": {
            "B07K3BHGL3": {
                "Sku": "B07K3BHGL3",
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                "GroupID": "MSLAPS2",
                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                "RegularPrice": 2999,
                "PromotionPrice": 2600,
                "Images": [
                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                ],
                "SearchKeywords": [
                    "Laptop",
                    "Microsoft",
                    "Surface"
                ],
                "Quantity": 200,
                "Category": [
                    "Computers & Tablets>Laptops"
                ],
                "Color": "Black",
                "Brand": "Microsoft",
                "Size": "13.5 inches",
                "Active": true,
                "Attributes": {
                    "ASIN": "B07K3BHGL3",
                    "Batteries": "1",
                    "Color": "Black",
                    "Date First Available": "Nov. 4 2018",
                    "Display Size": "13.5 inches",
                    "Flash Memory Size": "512.00",
                    "Item Weight": "1.28 Kg",
                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                    "Item model number": "DAL-00092",
                    "Memory Speed": "1 GHz",
                    "Number of USB 2.0 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "RAM": "16 GB",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac"
                },
                "IsMain": true,
                "Currency": "CDN",
                "Updated": 0
            }
        },
        "ProductQuantity": {
            "B07K3BHGL3": 1
        },
        "PaymentStatus": "PAYMENT_PENDING",
        "PaymentAmount": 2600,
        "Currency": "CDN",
        "OrderStatus": "ORDERED",
        "ShippingAddress": {
            "FirstName": "Vivek",
            "LastName": "Narang",
            "AddressLineOne": "111 Edgar Ave",
            "AddressLineTwo": "",
            "City": "Richmond Hill",
            "State": "Ontario",
            "Country": "Canada",
            "Pincode": "L4C 6K3"
        },
        "Attributes": null
    }
}
```

Use this endpoint to create an order entry in the database. Details on the fields and the constraints is listed below. 


### HTTP Request URL

`POST https://api.krama.ai/orders/{API version}/orders`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |
|Content-Type       | application/json                             |

### HTTP Body Request Parameters

|         Parameter                                |               Description                                                           |
|--------------------------------------------------|-------------------------------------------------------------------------------------|
| OrderID                                          | string, order id to uniquely identify an order in the system                        |
| CustomerID                                       | string, customer id to bind an order with a customer                                |
| PaymentStatus                                    | string, payent status identifier of an order                                        |
| PaymentAmount                                    | float,  total order amount                                                          |
| Currency                                         | string, payment currency like CAD, USD etc ...                                      |
| OrderStatus                                      | string, order status identifier of an order                                         |
| ShippingAddress.FirstName                        | string, first name of the customer as mentioned on the shipping address             |
| ShippingAddress.LastName                         | string, last name of the customer as mentioned on the shipping address              |
| ShippingAddress.AddressLineOne                   | string, address line one of the shipping address                                    |
| ShippingAddress.AddressLineTwo                   | string, address line two of the shipping address                                    |
| ShippingAddress.City                             | string, city of the shipping address                                                |
| ShippingAddress.State                            | string, state of the shipping address                                               |
| ShippingAddress.Country                          | string, country of the shipping address                                             |
| ShippingAddress.Pincode                          | string, pincode of the shipping address                                             |
| Products                                         | Map[SKU]=>{Product Object} Product map for reference                                |
| ProductQuantity                                  | Map[SKU]=>Quantity, Integer, product quantity map                                   |

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |


## Get orders for a customer

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Order Found ...",
    "Time": 1579030536721609239,
    "Response": {
        "OrderID": "1234567",
        "OrderCreationDate": 1579030377359016931,
        "OrderUpdateDate": 0,
        "CustomerID": "55348",
        "Products": {
            "B07K3BHGL3": {
                "Sku": "B07K3BHGL3",
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                "GroupID": "MSLAPS2",
                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                "RegularPrice": 2999,
                "PromotionPrice": 2600,
                "Images": [
                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                ],
                "SearchKeywords": [
                    "Laptop",
                    "Microsoft",
                    "Surface"
                ],
                "Quantity": 200,
                "Category": [
                    "Computers & Tablets>Laptops"
                ],
                "Color": "Black",
                "Brand": "Microsoft",
                "Size": "13.5 inches",
                "Active": true,
                "Attributes": {
                    "ASIN": "B07K3BHGL3",
                    "Batteries": "1",
                    "Color": "Black",
                    "Date First Available": "Nov. 4 2018",
                    "Display Size": "13.5 inches",
                    "Flash Memory Size": "512.00",
                    "Item Weight": "1.28 Kg",
                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                    "Item model number": "DAL-00092",
                    "Memory Speed": "1 GHz",
                    "Number of USB 2.0 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "RAM": "16 GB",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac"
                },
                "IsMain": true,
                "Currency": "CDN",
                "Updated": 0
            }
        },
        "ProductQuantity": {
            "B07K3BHGL3": 1
        },
        "PaymentStatus": "PAYMENT_PENDING",
        "PaymentAmount": 2600,
        "Currency": "CDN",
        "OrderStatus": "ORDERED",
        "ShippingAddress": {
            "FirstName": "Vivek",
            "LastName": "Narang",
            "AddressLineOne": "111 Edgar Ave",
            "AddressLineTwo": "",
            "City": "Richmond Hill",
            "State": "Ontario",
            "Country": "Canada",
            "Pincode": "L4C 6K3"
        },
        "Attributes": null
    }
}
```

> Sample API response for non existing order ID

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Order Not Found ...",
    "Time": 1578889659887108977,
    "Response": null
}
```

Use this endpoint to get all the orders associated with a customer using the customer ID. 


### HTTP Request URL

`GET https://api.krama.ai/orders/{API version}/orders/customer/{CID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |

### HTTP Request URL Parameters

| Parameter     | Description          |
|---------------|----------------------|
|CID            | The Customer ID      |

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |


## Get order by order ID

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Order Found ...",
    "Time": 1579030494721736586,
    "Response": {
        "OrderID": "1234567",
        "OrderCreationDate": 1579030377359016931,
        "OrderUpdateDate": 0,
        "CustomerID": "55348",
        "Products": {
            "B07K3BHGL3": {
                "Sku": "B07K3BHGL3",
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                "GroupID": "MSLAPS2",
                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                "RegularPrice": 2999,
                "PromotionPrice": 2600,
                "Images": [
                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                ],
                "SearchKeywords": [
                    "Laptop",
                    "Microsoft",
                    "Surface"
                ],
                "Quantity": 200,
                "Category": [
                    "Computers & Tablets>Laptops"
                ],
                "Color": "Black",
                "Brand": "Microsoft",
                "Size": "13.5 inches",
                "Active": true,
                "Attributes": {
                    "ASIN": "B07K3BHGL3",
                    "Batteries": "1",
                    "Color": "Black",
                    "Date First Available": "Nov. 4 2018",
                    "Display Size": "13.5 inches",
                    "Flash Memory Size": "512.00",
                    "Item Weight": "1.28 Kg",
                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                    "Item model number": "DAL-00092",
                    "Memory Speed": "1 GHz",
                    "Number of USB 2.0 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "RAM": "16 GB",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac"
                },
                "IsMain": true,
                "Currency": "CDN",
                "Updated": 0
            }
        },
        "ProductQuantity": {
            "B07K3BHGL3": 1
        },
        "PaymentStatus": "PAYMENT_PENDING",
        "PaymentAmount": 2600,
        "Currency": "CDN",
        "OrderStatus": "ORDERED",
        "ShippingAddress": {
            "FirstName": "Vivek",
            "LastName": "Narang",
            "AddressLineOne": "111 Edgar Ave",
            "AddressLineTwo": "",
            "City": "Richmond Hill",
            "State": "Ontario",
            "Country": "Canada",
            "Pincode": "L4C 6K3"
        },
        "Attributes": null
    }
}
```

> Sample API response for non existing order ID

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Order Not Found ...",
    "Time": 1578889659887108977,
    "Response": null
}
```

Use this endpoint to get all the orders associated with a customer using the customer ID. 


### HTTP Request URL

`GET https://api.krama.ai/orders/{API version}/orders/{OID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |

### HTTP Request URL Parameters

| Parameter     | Description          |
|---------------|----------------------|
|OID            | The Order ID         |

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |



## Update order

> Sample HTTP Request body:

```json
{
  "OrderID": "1234567",
  "CustomerID": "55348",
  "PaymentStatus": "PAID",
  "PaymentAmount": 2600.0,
  "Currency": "CDN",
  "OrderStatus": "ORDERED",
  "ShippingAddress": {
    "FirstName": "Vivek",
    "LastName": "Narang",
    "AddressLineOne": "111 Edgar Ave",
    "AddressLineTwo": "",
    "City": "Richmond Hill",
    "State": "Ontario",
    "Country": "Canada",
    "Pincode": "L4C 6K3"
  },
  "Products": {
    "B07K3BHGL3": {
        "Sku": "B07K3BHGL3",
        "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
        "GroupID": "MSLAPS2",
        "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
        "RegularPrice": 2999,
        "PromotionPrice": 2600,
        "Images": [
            "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
            "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
        ],
        "SearchKeywords": [
            "Laptop",
            "Microsoft",
            "Surface"
        ],
        "Quantity": 200,
        "Category": [
            "Computers & Tablets>Laptops"
        ],
        "Color": "Black",
        "Brand": "Microsoft",
        "Size": "13.5 inches",
        "Active": true,
        "Attributes": {
            "ASIN": "B07K3BHGL3",
            "Batteries": "1",
            "Color": "Black",
            "Date First Available": "Nov. 4 2018",
            "Display Size": "13.5 inches",
            "Flash Memory Size": "512.00",
            "Item Weight": "1.28 Kg",
            "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
            "Item model number": "DAL-00092",
            "Memory Speed": "1 GHz",
            "Number of USB 2.0 Ports": "1",
            "Operating System": "Windows 10 Home",
            "Processor Count": "16",
            "RAM": "16 GB",
            "Series": "Surface Laptop 2",
            "Shipping Weight": "2.2 kg",
            "Wireless Standard": "802.11ac"
        },
        "IsMain": true,
        "Currency": "CDN",
        "Updated": 0
    }
  },
  "ProductQuantity": {
    "B07K3BHGL3": 1
  }
}
```

> Sample valid API response:

```json
{
    "Code": 202,
    "Success": true,
    "Message": "Order Updated ...",
    "Time": 1579030671407502692,
    "Response": {
        "OrderID": "1234567",
        "OrderCreationDate": 0,
        "OrderUpdateDate": 0,
        "CustomerID": "55348",
        "Products": {
            "B07K3BHGL3": {
                "Sku": "B07K3BHGL3",
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                "GroupID": "MSLAPS2",
                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                "RegularPrice": 2999,
                "PromotionPrice": 2600,
                "Images": [
                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                ],
                "SearchKeywords": [
                    "Laptop",
                    "Microsoft",
                    "Surface"
                ],
                "Quantity": 200,
                "Category": [
                    "Computers & Tablets>Laptops"
                ],
                "Color": "Black",
                "Brand": "Microsoft",
                "Size": "13.5 inches",
                "Active": true,
                "Attributes": {
                    "ASIN": "B07K3BHGL3",
                    "Batteries": "1",
                    "Color": "Black",
                    "Date First Available": "Nov. 4 2018",
                    "Display Size": "13.5 inches",
                    "Flash Memory Size": "512.00",
                    "Item Weight": "1.28 Kg",
                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                    "Item model number": "DAL-00092",
                    "Memory Speed": "1 GHz",
                    "Number of USB 2.0 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "RAM": "16 GB",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac"
                },
                "IsMain": true,
                "Currency": "CDN",
                "Updated": 0
            }
        },
        "ProductQuantity": {
            "B07K3BHGL3": 1
        },
        "PaymentStatus": "PAID",
        "PaymentAmount": 2600,
        "Currency": "CDN",
        "OrderStatus": "ORDERED",
        "ShippingAddress": {
            "FirstName": "Vivek",
            "LastName": "Narang",
            "AddressLineOne": "111 Edgar Ave",
            "AddressLineTwo": "",
            "City": "Richmond Hill",
            "State": "Ontario",
            "Country": "Canada",
            "Pincode": "L4C 6K3"
        },
        "Attributes": null
    }
}
```

> Sample invalid API response (when an order with order ID does not exist):

```json
{
    "Code": 304,
    "Success": false,
    "Message": "Order Not Found ...",
    "Time": 1578890459696054203,
    "Response": null
}
```


Use this endpoint to update the order object. 


### HTTP Request URL

`PUT https://api.krama.ai/orders/{API version}/orders/{ID}`


### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |
|Content-Type       | application/json                             |


### URL Parameters

|Parameter                |               Description                            |
|-------------------------|------------------------------------------------------|
|ID                       |  Order ID                                            |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |




## Delete order by order ID

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Order Deleted ...",
    "Time": 1578890915789973237,
    "Response": null
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Order Not Found ...",
    "Time": 1578890926357243724,
    "Response": null
}
```

Use this endpoint to delete an order using the order ID. Recommended to use this endpoint carefully. 


### HTTP Request

`DELETE https://api.krama.ai/orders/{API version}/orders/{ID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |

### HTTP Request URL Parameters

| Parameter     | Description          |
|---------------|----------------------|
| ID            | The Order ID         |

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |




# Search API

## Basic search

> Sample HTTP request body:

```json
{
  "Q": "Surface Laptop 2",
  "Fields": [
    "Name",
    "Skus"
  ]
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Search Result ...",
    "Time": 1579030869519210940,
    "Response": {
        "took": 88,
        "hits": {
            "total": {
                "value": 1,
                "relation": "eq"
            },
            "max_score": 0.26103413,
            "hits": [
                {
                    "_score": 0.26103413,
                    "_index": "ffaabbccdd.productgroup.index",
                    "_type": "_doc",
                    "_id": "MSLAPS2",
                    "_seq_no": null,
                    "_primary_term": null,
                    "_source": {
                        "GroupID": "MSLAPS2",
                        "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                        "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                        "RegularPriceMin": 2700,
                        "RegularPriceMax": 2700,
                        "PromotionPriceMin": 2500,
                        "PromotionPriceMax": 2500,
                        "Skus": [
                            "B07K3BHGL3"
                        ],
                        "Images": [
                            "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                            "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                        ],
                        "SearchKeywords": [
                            "Laptop",
                            "Microsoft",
                            "Surface"
                        ],
                        "Category": [
                            "Computers & Tablets>Laptops"
                        ],
                        "Colors": [
                            "Black"
                        ],
                        "Brands": [
                            "Microsoft"
                        ],
                        "Sizes": [
                            "13.5 inches"
                        ],
                        "Active": true,
                        "Currency": "CDN",
                        "Updated": 1579028967414440899,
                        "Products": {
                            "B07K3BHGL3": {
                                "Sku": "B07K3BHGL3",
                                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                                "GroupID": "MSLAPS2",
                                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                                "RegularPrice": 2700,
                                "PromotionPrice": 2500,
                                "Images": [
                                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                                ],
                                "SearchKeywords": [
                                    "Laptop",
                                    "Microsoft",
                                    "Surface"
                                ],
                                "Quantity": 500,
                                "Category": [
                                    "Computers & Tablets>Laptops"
                                ],
                                "Color": "Black",
                                "Brand": "Microsoft",
                                "Size": "13.5 inches",
                                "Active": true,
                                "Attributes": {
                                    "ASIN": "B07K3BHGL3",
                                    "Batteries": "1",
                                    "Color": "Black",
                                    "Date First Available": "Nov. 4 2018",
                                    "Display Size": "13.5 inches",
                                    "Flash Memory Size": "512.00",
                                    "Item Weight": "1.28 Kg",
                                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                                    "Item model number": "DAL-00092",
                                    "Memory Speed": "1 GHz",
                                    "Number of USB 2.0 Ports": "1",
                                    "Operating System": "Windows 10 Home",
                                    "Processor Count": "16",
                                    "RAM": "16 GB",
                                    "Series": "Surface Laptop 2",
                                    "Shipping Weight": "2.2 kg",
                                    "Wireless Standard": "802.11ac"
                                },
                                "IsMain": true,
                                "Currency": "CDN",
                                "Updated": 0
                            }
                        },
                        "Attributes": {
                            "ASIN": [
                                "B07K3BHGL3"
                            ],
                            "Batteries": [
                                "1"
                            ],
                            "Color": [
                                "Black"
                            ],
                            "Date First Available": [
                                "Nov. 4 2018"
                            ],
                            "Display Size": [
                                "13.5 inches"
                            ],
                            "Flash Memory Size": [
                                "512.00"
                            ],
                            "Item Weight": [
                                "1.28 Kg"
                            ],
                            "Item dimensions L x W x H": [
                                "17.8 x 12.7 x 15.2 cm"
                            ],
                            "Item model number": [
                                "DAL-00092"
                            ],
                            "Memory Speed": [
                                "1 GHz"
                            ],
                            "Number of USB 2.0 Ports": [
                                "1"
                            ],
                            "Operating System": [
                                "Windows 10 Home"
                            ],
                            "Processor Count": [
                                "16"
                            ],
                            "RAM": [
                                "16 GB"
                            ],
                            "Series": [
                                "Surface Laptop 2"
                            ],
                            "Shipping Weight": [
                                "2.2 kg"
                            ],
                            "Wireless Standard": [
                                "802.11ac"
                            ]
                        }
                    }
                }
            ]
        },
        "aggregations": {
            "Brands": {
                "doc_count_error_upper_bound": 0,
                "sum_other_doc_count": 0,
                "buckets": [
                    {
                        "key": "microsoft",
                        "doc_count": 1
                    }
                ]
            },
            "Colors": {
                "doc_count_error_upper_bound": 0,
                "sum_other_doc_count": 0,
                "buckets": [
                    {
                        "key": "black",
                        "doc_count": 1
                    }
                ]
            },
            "Sizes": {
                "doc_count_error_upper_bound": 0,
                "sum_other_doc_count": 0,
                "buckets": [
                    {
                        "key": "13.5",
                        "doc_count": 1
                    }
                ]
            }
        },
        "_shards": {
            "total": 1,
            "successful": 1,
            "failed": 0
        }
    }
}
```

Use this API endpoint to search a product group in the search index. Please note that product group objects are only returned in the response. 


### HTTP Request

`GET https://api.krama.ai/search/{API version}/productgroups/search`


### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| Q                     |  String, Query string                               |
| Fields                |  []String, Multiple, query fields                   |

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |




# Response Codes

This API uses the following HTTP Response codes:

| Response Code                      | Meaning                                                                 |
|------------------------------------|-------------------------------------------------------------------------|
| 200  **StatusOK**                  | The API call was successful                                             |
| 201  **StatusCreated**             | The resource was successfully created                                   |
| 202  **StatusAccepted**            | The API request was accepted                                            |
| 304  **StatusNotModified**         | The request to modify the resource failed for some reason               |
| 400  **StatusBadRequest**          | The API call was malformed                                              |
| 401  **StatusUnauthorized**        | The API request is not authorized. Please check your access credentials |
| 404  **StatusNotFound**            | The requested/referenced resource was not found                         |
| 500  **StatusInternalServerError** | Something went wrong on our server                                      |
| 503  **StatusServiceUnavailable**  | The API service is down                                                 |