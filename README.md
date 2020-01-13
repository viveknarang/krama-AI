# KRAMA - Introduction

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
    "CustomerID" : "0100456",
    "APIKey"     : "123445"
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Login Successful ...",
    "Time": 1578855509733742430,
    "Response": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjeHMiOiJmZmFhYmJjY2RkIiwiZXhwIjoxNTc4OTM1NTA5LCJpYXQiOjE1Nzg4NTU1MDksIm5iZiI6MTU3ODg1NTQwOSwidWlkIjoiYzU1NWRmMDEtZTlkYy00NDU3LThlMzEtNWQ2ZjU3ODZmOTJiIn0.MQVnZymGzDjR6cW9QHT8SC6HazqkkJRcowqx85C2wrw",
        "validForSeconds": 80000
    }
}
```

> Sample invalid API response (valid token expired or invalid token used):

```json
{
    "Code": 401,
    "Success": true,
    "Message": "You need to either login or your access token is either expired, invalid, or corrupt ...",
    "Time": 1578882700872915416,
    "Response": null
}
```

> Sample invalid API response:

```json
{
    "Code": 401,
    "Success": true,
    "Message": "Login Failed! Please check your credentials and also make sure that you are an active customer ...",
    "Time": 1578882733720622060,
    "Response": null
}
```

This endpoint gets you your API access token. You need to send your customer ID and the API key that we provided you for using our platform. Upon receiving your valid credentials, the API will respond with a token with additional information including the validFor key which tells you how long this access token is valid for. Please set **x-access-token** to the value of the **token**, in the header of your subsequent API calls. 


### HTTP Request URL

`GET http://api.gallao.io/customers/{API version}/login`

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
With the field validFor in response, you can calculate the time after with your servers need to login again to get a new token.
</aside>


# Catalog API

## Add a new product

> Sample HTTP request body:

```json
{
	 "Sku":"1234",
	 "Name": "Test Product 1234",
	 "Images" : ["https://homepages.cae.wisc.edu/~ece533/images/airplane.png"],
	 "Description" : "Test Description 1234",
	 "GroupID" : "55446677", 
	 "SearchKeywords" : ["A","B", "C"],
	 "RegularPrice" : 12.55,
	 "PromotionPrice" : 10.99,
	 "Currency" : "USD",
	 "IsMain" : true,
	 "Quantity" : 45,
	 "Size" : "36D",
	 "Brand" : "VS",
	 "Color" : "Red",
	 "Category" : ["A>B>C>D"],
	 "Active" : true
}
```

> Sample valid API response:

```json
{
    "Code": 201,
    "Success": true,
    "Message": "Product Added ...",
    "Time": 1578884392984514865,
    "Response": {
        "Sku": "1234",
        "Name": "Test Product 1234",
        "GroupID": "55446677",
        "Description": "Test Description 1234",
        "RegularPrice": 12.55,
        "PromotionPrice": 10.99,
        "Images": [
            "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
        ],
        "SearchKeywords": [
            "A",
            "B",
            "C"
        ],
        "Quantity": 45,
        "Category": [
            "A>B>C>D"
        ],
        "Color": "Red",
        "Brand": "VS",
        "Size": "36D",
        "Active": true,
        "Attributes": null,
        "IsMain": true,
        "Currency": "USD",
        "Updated": 1578884392961353657
    }
}
```

Use this API endpoint to add a new product in the products collection. When a product is added in the products collection, this product is also added in product group collection. If the product group with the matching groupID is missing, a new product group is formed. Search index and the cache are also automatically updated with a valid call to this endpoint. 


### HTTP Request URL

`POST http://api.gallao.io/catalog/{API version}/products`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |
|Content-Type       | application/json                             |

### HTTP Request Body Parameters

|    Parameter   |          Constraints         |        Description                                           |
|----------------|------------------------------|--------------------------------------------------------------|
|sku             |   String, Max 50 Characters  |  The SKU of the product                                      |
|name            |   Text, Max 250 Characters   |  The name of the product                                     |
|description     |   Text, Max 2048 Characters  |  The description of the product                              | 
|groupID         |   String, Max 50 Characters  |  The product group ID                                        | 
|regularPrice    |   Float, Greater than 0      |  Everyday price                                              | 
|promotionPrice  |   Float, Greater than 0      |  On-sale price                                               | 
|images          |   URL, Mandatory             |  Product images                                              | 
|searchKeywords  |   Text, Mandatory            |  Keywords that you want this product to be searched with     |
|quantity        |   Integer, Greater than 0    |  Inventory stock quantity                                    | 
|category        |   Text, Mandatory            |  Category breadcrumbs                                        | 
|color           |   Text, Optional             |  Product color                                               |
|brand           |   Text, Optional             |  Product brand                                               | 
|size            |   Text, Optional             |  Product size                                                | 
|active          |   Boolean, Mandatory         |  Is product available for sale?                              |    
|isMain          |   Boolean, Mandatory         |  Is the product main product in the group?                   |


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
    "Success": true,
    "Message": "Product Found ...",
    "Time": 1578884957448467117,
    "Response": {
        "Sku": "1234",
        "Name": "Test Product 1234",
        "GroupID": "55446677",
        "Description": "Test Description 1234",
        "RegularPrice": 12.55,
        "PromotionPrice": 10.99,
        "Images": [
            "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
        ],
        "SearchKeywords": [
            "A",
            "B",
            "C"
        ],
        "Quantity": 45,
        "Category": [
            "A>B>C>D"
        ],
        "Color": "Red",
        "Brand": "VS",
        "Size": "36D",
        "Active": true,
        "Attributes": null,
        "IsMain": true,
        "Currency": "USD",
        "Updated": 1578867116584475564
    }
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": true,
    "Message": "Product Not Found ...",
    "Time": 1578884979555938176,
    "Response": null
}
```

When you want to get a specific product you can use this endpoint. All you need to pass is your access token and the SKU. This endpoint is cached for efficiency but also ensures that updated product data is served when applicable. 


### HTTP Request URL

`GET http://api.gallao.io/catalog/{API version}/products/{SKU}`

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
| sku               | The SKU of the product                                                        |
| name              | The name of the product                                                       |
| description       | The description of the product                                                |   
| groupID           | The product group ID                                                          |
| regularPrice      | Everyday price                                                                |
| promotionPrice    | On-sale price                                                                 |
| images            | Product images                                                                |
| searchKeywords    | Keywords that you want this product to be searched with                       |
| quantity          | Inventory stock quantity                                                      |
| category          | Category breadcrumbs                                                          |
| color             | Product color                                                                 |
| brand             | Product brand                                                                 |
| size              | Product size                                                                  |
| active            | Is product available for sale?                                                |
| isMain            | Is the product main product in the group?                                     |




## Update a product

> Sample HTTP request body:

```json
{
	 "Sku":"1234",
	 "Name": "Test Product 1234",
	 "Images" : ["https://homepages.cae.wisc.edu/~ece533/images/airplane.png"],
	 "Description" : "Test Description 1234",
	 "GroupID" : "55446677", 
	 "SearchKeywords" : ["A","B", "C"],
	 "RegularPrice" : 12.55,
	 "PromotionPrice" : 10.99,
	 "Currency" : "USD",
	 "IsMain" : true,
	 "Quantity" : 45,
	 "Size" : "36D",
	 "Brand" : "VS",
	 "Color" : "Red",
	 "Category" : ["A>B>C>D"],
	 "Active" : true
}
```

> Sample valid API response:

```json
{
    "Code": 202,
    "Success": true,
    "Message": "Product Updated ...",
    "Time": 1578885765895872226,
    "Response": {
        "Sku": "1235",
        "Name": "Test Product 1235",
        "GroupID": "55446677",
        "Description": "Test Description 1235",
        "RegularPrice": 99.99,
        "PromotionPrice": 3.99,
        "Images": [
            "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
        ],
        "SearchKeywords": [
            "D",
            "E",
            "F"
        ],
        "Quantity": 67443,
        "Category": [
            "A>B>C>D"
        ],
        "Color": "Pink",
        "Brand": "VS",
        "Size": "36A",
        "Active": true,
        "Attributes": null,
        "IsMain": true,
        "Currency": "USD",
        "Updated": 0
    }
}
```

> Sample valid API response when there is nothing to update:

```json
{
    "Code": 304,
    "Success": true,
    "Message": "Nothing to Update ...",
    "Time": 1578885774238860804,
    "Response": null
}
```

> Sample invalid API response:

```json
{
    "Code": 304,
    "Success": true,
    "Message": "Product Not Found ...",
    "Time": 1578885669974853990,
    "Response": null
}
```

Use this API endpoint to update your product information in the catalog. For now you need to pass the entire product object with updated parts (this functionality will be improved very soon). When you hit this endpoint, the data in the products collection gets updated, product group data also gets updated automatically, search index is also updated and the cache entry is removed first and updated on the next GET call. 


### HTTP Request URL

`PUT http://api.gallao.io/catalog/{API version}/products/{SKU}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|x-access-token     | The access token that you receive upon login |
|Content-Type       | application/json                             |


### HTTP Body Parameters

|    Parameter   |          Constraints         |        Description                                           |
|----------------|------------------------------|--------------------------------------------------------------|
|sku             |   String, Max 50 Characters  |  The SKU of the product                                      |
|name            |   Text, Max 250 Characters   |  The name of the product                                     |
|description     |   Text, Max 2048 Characters  |  The description of the product                              | 
|groupID         |   String, Max 50 Characters  |  The product group ID                                        | 
|regularPrice    |   Float, Greater than 0      |  Everyday price                                              | 
|promotionPrice  |   Float, Greater than 0      |  On-sale price                                               | 
|images          |   URL, Mandatory             |  Product images                                              | 
|searchKeywords  |   Text, Mandatory            |  Keywords that you want this product to be searched with     |
|quantity        |   Integer, Greater than 0    |  Inventory stock quantity                                    | 
|category        |   Text, Mandatory            |  Category breadcrumbs                                        | 
|color           |   Text, Optional             |  Product color                                               |
|brand           |   Text, Optional             |  Product brand                                               | 
|size            |   Text, Optional             |  Product size                                                | 
|active          |   Boolean, Mandatory         |  Is product available for sale?                              |    
|isMain          |   Boolean, Mandatory         |  Is the product main product in the group?                   |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |
| sku               | The SKU of the product                                                        |
| name              | The name of the product                                                       |
| description       | The description of the product                                                |   
| groupID           | The product group ID                                                          |
| regularPrice      | Everyday price                                                                |
| promotionPrice    | On-sale price                                                                 |
| images            | Product images                                                                |
| searchKeywords    | Keywords that you want this product to be searched with                       |
| quantity          | Inventory stock quantity                                                      |
| category          | Category breadcrumbs                                                          |
| color             | Product color                                                                 |
| brand             | Product brand                                                                 |
| size              | Product size                                                                  |
| active            | Is product available for sale?                                                |
| isMain            | Is the product main product in the group?                                     |



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
    "Success": true,
    "Message": "Product Not Found ...",
    "Time": 1578886152262487961,
    "Response": null
}
```

Use this API endpoint to remove a product from the catalog. When you hit this endpoint with a valid request, the product in the products collection gets removed, the productgroups collection is also automatically updated and the cache and search index is also updated. If there was only one product in the product group the product group object is also removed from the productgroups collection. 
 

### HTTP Request URL

`DELETE http://api.gallao.io/catalog/{API version}/products/{SKU}`

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
    "Time": 1578886406550794133,
    "Response": {
        "GroupID": "55446677",
        "Name": "Test Product 1235",
        "Description": "Test Description 1235",
        "RegularPriceMin": 99.99,
        "RegularPriceMax": 99.99,
        "PromotionPriceMin": 3.99,
        "PromotionPriceMax": 3.99,
        "Skus": [
            "1235"
        ],
        "Images": [
            "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
        ],
        "SearchKeywords": [
            "D",
            "E",
            "F"
        ],
        "Category": [
            "A>B>C>D"
        ],
        "Colors": [
            "Pink"
        ],
        "Brands": [
            "VS"
        ],
        "Sizes": [
            "36A"
        ],
        "Active": true,
        "Currency": "USD",
        "Updated": 1578884392962398268,
        "Products": {
            "1235": {
                "Sku": "1235",
                "Name": "Test Product 1235",
                "GroupID": "55446677",
                "Description": "Test Description 1235",
                "RegularPrice": 99.99,
                "PromotionPrice": 3.99,
                "Images": [
                    "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
                ],
                "SearchKeywords": [
                    "D",
                    "E",
                    "F"
                ],
                "Quantity": 67443,
                "Category": [
                    "A>B>C>D"
                ],
                "Color": "Pink",
                "Brand": "VS",
                "Size": "36A",
                "Active": true,
                "Attributes": null,
                "IsMain": true,
                "Currency": "USD",
                "Updated": 0
            }
        }
    }
}
```

> Sample valid API response:

```json
{
    "Code": 404,
    "Success": true,
    "Message": "Product Group Not Found ...",
    "Time": 1578886448802758717,
    "Response": null
}
```


This API endpoint gets a specific product group by product group ID. This endpoint is cached for efficiency and speed. 


### HTTP Request URL

`GET http://api.gallao.io/catalog/{API version}/productgroups/{PGID}`

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
    "Success": true,
    "Message": "Product Group Not Found ...",
    "Time": 1578886787090095342,
    "Response": null
}
```

Use this API endpoint to remove a product group from the productgroups collection in the database. This call also updates the search index. When a product group is deleted, entries of related products in the products collection are also removed. 
 

### HTTP Request URL

`DELETE http://api.gallao.io/catalog/{API version}/productgroups/{PGID}`

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
  "PaymentAmount": 10.99,
  "Currency": "CAD",
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
    "1234": {
      "Sku": "1234",
      "Name": "Test Product 1234",
      "GroupID": "55446677",
      "Description": "Test Description 1234",
      "RegularPrice": 12.55,
      "PromotionPrice": 10.99,
      "Images": [
        "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
      ],
      "SearchKeywords": [
        "A",
        "B",
        "C"
      ],
      "Quantity": 45,
      "Category": [
        "A>B>C>D"
      ],
      "Color": "Red",
      "Brand": "VS",
      "Size": "36D",
      "Active": true,
      "Attributes": null,
      "IsMain": true,
      "Currency": "USD",
      "Updated": 1578888029190398500
    }
  },
  "ProductQuantity": {
    "1234": 1
  }
}
```


> Sample valid API response:

```json
{
    "Code": 201,
    "Success": true,
    "Message": "Order Created ...",
    "Time": 1578888156909048030,
    "Response": {
        "OrderID": "1234567",
        "OrderCreationDate": 1578888156908373526,
        "OrderUpdateDate": 0,
        "CustomerID": "55348",
        "Products": {
            "1234": {
                "Sku": "1234",
                "Name": "Test Product 1234",
                "GroupID": "55446677",
                "Description": "Test Description 1234",
                "RegularPrice": 12.55,
                "PromotionPrice": 10.99,
                "Images": [
                    "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
                ],
                "SearchKeywords": [
                    "A",
                    "B",
                    "C"
                ],
                "Quantity": 45,
                "Category": [
                    "A>B>C>D"
                ],
                "Color": "Red",
                "Brand": "VS",
                "Size": "36D",
                "Active": true,
                "Attributes": null,
                "IsMain": true,
                "Currency": "USD",
                "Updated": 1578888029190398500
            }
        },
        "ProductQuantity": {
            "1234": 1
        },
        "PaymentStatus": "PAYMENT_PENDING",
        "PaymentAmount": 10.99,
        "Currency": "CAD",
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

`POST http://api.gallao.io/orders/{API version}/orders`

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
    "Time": 1578889591913632532,
    "Response": {
        "OrderID": "1234567",
        "OrderCreationDate": 1578888156908373526,
        "OrderUpdateDate": 0,
        "CustomerID": "55348",
        "Products": {
            "1234": {
                "Sku": "1234",
                "Name": "Test Product 1234",
                "GroupID": "55446677",
                "Description": "Test Description 1234",
                "RegularPrice": 12.55,
                "PromotionPrice": 10.99,
                "Images": [
                    "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
                ],
                "SearchKeywords": [
                    "A",
                    "B",
                    "C"
                ],
                "Quantity": 45,
                "Category": [
                    "A>B>C>D"
                ],
                "Color": "Red",
                "Brand": "VS",
                "Size": "36D",
                "Active": true,
                "Attributes": null,
                "IsMain": true,
                "Currency": "USD",
                "Updated": 1578888029190398500
            }
        },
        "ProductQuantity": {
            "1234": 1
        },
        "PaymentStatus": "PAYMENT_PENDING",
        "PaymentAmount": 10.99,
        "Currency": "CAD",
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
    "Success": true,
    "Message": "Order Not Found ...",
    "Time": 1578889659887108977,
    "Response": null
}
```

Use this endpoint to get all the orders associated with a customer using the customer ID. 


### HTTP Request URL

`GET http://api.gallao.io/orders/{API version}/orders/customer/{CID}`

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
    "Time": 1578889591913632532,
    "Response": {
        "OrderID": "1234567",
        "OrderCreationDate": 1578888156908373526,
        "OrderUpdateDate": 0,
        "CustomerID": "55348",
        "Products": {
            "1234": {
                "Sku": "1234",
                "Name": "Test Product 1234",
                "GroupID": "55446677",
                "Description": "Test Description 1234",
                "RegularPrice": 12.55,
                "PromotionPrice": 10.99,
                "Images": [
                    "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
                ],
                "SearchKeywords": [
                    "A",
                    "B",
                    "C"
                ],
                "Quantity": 45,
                "Category": [
                    "A>B>C>D"
                ],
                "Color": "Red",
                "Brand": "VS",
                "Size": "36D",
                "Active": true,
                "Attributes": null,
                "IsMain": true,
                "Currency": "USD",
                "Updated": 1578888029190398500
            }
        },
        "ProductQuantity": {
            "1234": 1
        },
        "PaymentStatus": "PAYMENT_PENDING",
        "PaymentAmount": 10.99,
        "Currency": "CAD",
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
    "Success": true,
    "Message": "Order Not Found ...",
    "Time": 1578889659887108977,
    "Response": null
}
```

Use this endpoint to get all the orders associated with a customer using the customer ID. 


### HTTP Request URL

`GET http://api.gallao.io/orders/{API version}/orders/{OID}`

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
  "PaymentAmount": 10.99,
  "Currency": "CAD",
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
    "1234": {
      "Sku": "1234",
      "Name": "Test Product 1234",
      "GroupID": "55446677",
      "Description": "Test Description 1234",
      "RegularPrice": 12.55,
      "PromotionPrice": 10.99,
      "Images": [
        "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
      ],
      "SearchKeywords": [
        "A",
        "B",
        "C"
      ],
      "Quantity": 45,
      "Category": [
        "A>B>C>D"
      ],
      "Color": "Red",
      "Brand": "VS",
      "Size": "36D",
      "Active": true,
      "Attributes": null,
      "IsMain": true,
      "Currency": "USD",
      "Updated": 1578888029190398500
    }
  },
  "ProductQuantity": {
    "1234": 1
  }
}
```

> Sample valid API response:

```json
{
    "Code": 202,
    "Success": true,
    "Message": "Order Updated ...",
    "Time": 1578890371677332692,
    "Response": {
        "OrderID": "1234567",
        "OrderCreationDate": 0,
        "OrderUpdateDate": 0,
        "CustomerID": "55348",
        "Products": {
            "1234": {
                "Sku": "1234",
                "Name": "Test Product 1234",
                "GroupID": "55446677",
                "Description": "Test Description 1234",
                "RegularPrice": 12.55,
                "PromotionPrice": 10.99,
                "Images": [
                    "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
                ],
                "SearchKeywords": [
                    "A",
                    "B",
                    "C"
                ],
                "Quantity": 45,
                "Category": [
                    "A>B>C>D"
                ],
                "Color": "Red",
                "Brand": "VS",
                "Size": "36D",
                "Active": true,
                "Attributes": null,
                "IsMain": true,
                "Currency": "USD",
                "Updated": 1578888029190398500
            }
        },
        "ProductQuantity": {
            "1234": 1
        },
        "PaymentStatus": "PAID",
        "PaymentAmount": 10.99,
        "Currency": "CAD",
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

> Sample invalid API response (when nothing to update ...):

```json
{
    "Code": 304,
    "Success": true,
    "Message": "Order Not Updated ...",
    "Time": 1578890430356224891,
    "Response": null
}
```

> Sample invalid API response (when an order with order ID does not exist):

```json
{
    "Code": 304,
    "Success": true,
    "Message": "Order Not Found ...",
    "Time": 1578890459696054203,
    "Response": null
}
```


Use this endpoint to update the order object. 


### HTTP Request URL

`PUT http://api.gallao.io/orders/{API version}/orders/{ID}`


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
    "Success": true,
    "Message": "Order Not Found ...",
    "Time": 1578890926357243724,
    "Response": null
}
```

Use this endpoint to delete an order using the order ID. Recommended to use this endpoint carefully. 


### HTTP Request

`DELETE http://api.gallao.io/orders/{API version}/orders/{ID}`

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
  "Q" : "1234",
  "Fields" : ["Name", "Skus"]
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Search Result ...",
    "Time": 1578891074650533153,
    "Response": {
        "took": 8,
        "hits": {
            "total": {
                "value": 1,
                "relation": "eq"
            },
            "max_score": 0.2876821,
            "hits": [
                {
                    "_score": 0.2876821,
                    "_index": "ffaabbccdd.productgroup.index",
                    "_type": "_doc",
                    "_id": "55446677",
                    "_seq_no": null,
                    "_primary_term": null,
                    "_source": {
                        "GroupID": "55446677",
                        "Name": "Test Product 1234",
                        "Description": "Test Description 1234",
                        "RegularPriceMin": 12.55,
                        "RegularPriceMax": 12.55,
                        "PromotionPriceMin": 10.99,
                        "PromotionPriceMax": 10.99,
                        "Skus": [
                            "1234"
                        ],
                        "Images": [
                            "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
                        ],
                        "SearchKeywords": [
                            "B",
                            "C",
                            "A"
                        ],
                        "Category": [
                            "A>B>C>D"
                        ],
                        "Colors": [
                            "Red"
                        ],
                        "Brands": [
                            "VS"
                        ],
                        "Sizes": [
                            "36D"
                        ],
                        "Active": true,
                        "Currency": "USD",
                        "Updated": 1578888029191572453,
                        "Products": {
                            "1234": {
                                "Sku": "1234",
                                "Name": "Test Product 1234",
                                "GroupID": "55446677",
                                "Description": "Test Description 1234",
                                "RegularPrice": 12.55,
                                "PromotionPrice": 10.99,
                                "Images": [
                                    "https://homepages.cae.wisc.edu/~ece533/images/airplane.png"
                                ],
                                "SearchKeywords": [
                                    "A",
                                    "B",
                                    "C"
                                ],
                                "Quantity": 45,
                                "Category": [
                                    "A>B>C>D"
                                ],
                                "Color": "Red",
                                "Brand": "VS",
                                "Size": "36D",
                                "Active": true,
                                "Attributes": null,
                                "IsMain": true,
                                "Currency": "USD",
                                "Updated": 1578888029190398408
                            }
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
                        "key": "vs",
                        "doc_count": 1
                    }
                ]
            },
            "Colors": {
                "doc_count_error_upper_bound": 0,
                "sum_other_doc_count": 0,
                "buckets": [
                    {
                        "key": "red",
                        "doc_count": 1
                    }
                ]
            },
            "Sizes": {
                "doc_count_error_upper_bound": 0,
                "sum_other_doc_count": 0,
                "buckets": [
                    {
                        "key": "36d",
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

`GET http://api.gallao.io/search/{API version}/productgroups/search`


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