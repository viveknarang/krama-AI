
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/67133763de0b42ad917afae5b28f9bd4)](https://app.codacy.com/manual/viveknarang/krama-AI?utm_source=github.com&utm_medium=referral&utm_content=viveknarang/krama-AI&utm_campaign=Badge_Grade_Dashboard)
[![Build Status](https://travis-ci.org/viveknarang/krama-AI.svg?branch=master)](https://travis-ci.org/viveknarang/krama-AI)  [![Go Report Card](https://goreportcard.com/badge/github.com/viveknarang/krama-AI)](https://goreportcard.com/report/github.com/viveknarang/krama-AI)

# Krama AI - A blazing fast E-commerce AI platform

# Introduction

Krama AI is an e-commerce AI platform that provides a portfolio of novel and powerful features to build an online e-commerce store. The headless, API-first approach allows our customers to utilize platform features to build online stores with exceptional flexibility. Using Krama AI, individuals and businesses can build online e-commerce stores with user interface of their choice - be it a website, a mobile app or any other possible interface. Krama AI provides basic e-commerce platform components such as catalog, orders & shopping cart as well as advanced features such as sophisticated search capabilities, recommendation features to increase sales conversion rates and customer engagement. Krama AI also plans to provide sophisticated analytics & insights API that will give a competitive edge to businesses and will open new avenues for better customer engagement, inventory planning and price optimization and much more ... Stay tuned!

Krama AI is using state-of-art open-source software platforms and technologies to provide you with the best performance, better than any other SaaS platform in the market today. 

Krama AI is powered by:

- Golang            
- Redis             
- Elasticsearch     
- MongoDB
- Kafka (upcoming)
- TensorFlow (upcoming)
- Kubernetes
- Other supporting tools           


<aside class="success">
The current API version is: v1 Please replace {API version} with v1 in your API calls
</aside>


# Data Structures

## Product

> Sample valid product object:

```json
{
  "Sku": "B07K3BHGL5",
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
  "RegularPrice": 3300,
  "PromotionPrice": 3000,
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
  "Selectors": {
    "RAM": "64 GB",
    "Display Size": "15 inches",
    "Memory Speed": "2.88 GHz"
  },
  "Attributes": {
    "Wireless Standard": "802.11ac",
    "Number of USB 2 Ports": "1",
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

At the heart of that SaaS platform lies the concept of product. Product is the most basic unit of data on this platform. The product data structure provides you with a skeleton to model your inventory products for this platform. The product data structure provide a flexible way to allow the modelling of any kind of product that you may have in your inventory. The data structure comes with some mandatory fields, some optional fields and some flexible fields. Mandatory fields are those fields in the data structure that expects you to pass in relevant information following certain rules (as mentioned in the table below). Optional fields are those fields that are not necessary for the system but are good to have. 

The two most important fields that are needed to be discussed in detail are the Attributes and Selectors fields. These two fields in the data structure provide immense power as it allows the modelling of any type of real-world product. Attributes field is a key-value field that allows you to mention any other field that has not been included already. The attributes field is optional - use it if you need it. Selectors field is another optional key-value field that allows you to define all those product attributes that you think are crucial in helping identify a variation of the product in the product group. As a general thumb of rule, you should not define a product attribute more than once in the data-structure. 

The product data structure provides with all essential product attributes that are useful in modelling your product for this platform. Please feel have a look in the table below to find more details on each field and its constraints. 

<aside class="warning">
To ensure that the platform has clean data, the API is strict and will decline malformed data objects. It is essential for you to follow these rules for the API to normally accept your requests. 
</aside>

Please find the field definitions, types and constraints below:

|   Field          |   Type         |     Short Description                                         |    Constraints                                                          |
|------------------|----------------|---------------------------------------------------------------|-------------------------------------------------------------------------|
|  Sku             |   String       | Unique product identifier                                     | Mandatory, Unique, Alphanumeric, Less than 50 characters                |
|  Name            |   String       | Name of the product                                           | Mandatory, Less than 100 characters                                     |
|  GroupID         |   String       | Product group identifier. More details in other section       | Mandatory, Less than 50 characters                                      |
|  Description     |   String       | Product description field                                     | Mandatory, less than 10240 characters                                   |
|  RegularPrice    |   Float        | Product's everyday price                                      | Mandatory, Cannot be negative                                           |
|  PromotionPrice  |   Float        | Product's promotion price. Typically less than regular price. | Mandatory, Cannot be negative                                           |
|  Images          |   String       | Product image links                                           | Mandatory, need to be valid URLs, cannot be more than 100 URLs          |
|  SearchKeywords  |   String[]     | Product search keywords                                       | Mandatory, cannot be more than 100 search keywords                      |
|  Quantity        |   Integer      | Product stock quantity field                                  | Mandatory, cannot be negative                                           |
|  Category        |   String[]     | Product category path. Please see product object example      | Mandatory, '>' separated category path                                  |
|  Color           |   String       | Product color field                                           | Optional, cannot be greater than 100 characters                         |         
|  Brand           |   String       | Product brand field                                           | Optional, cannot be greater than 100 characters                         |
|  Size            |   String       | Product size field                                            | Optional, cannot be greater than 100 characters                         |
|  Active          |   Boolean      | Field to mark product available for sale                      | Boolean - either true or false                                          |
|  IsMain          |   Boolean      | Field to mark the product as a main product in the group      | Boolean - either true of false                                          |
|  Currency        |   String       | Product purchase currency                                     | Either - "USD", "CAD", "CDN", "INR", "GBP" or "EUR" (for now!)          |
|  Attributes      |   Map{k,v}     | Product custom attributes that fit your needs                 | Keys should be strings (alphanumeric with single space and "-" or "_") and values either: Int, Float, String or Boolean |
|  Selectors       |   Map{k,v}     | Those attributes that are used to identify a variation of a product in the group | Keys should be strings (alphanumeric with single space and "-" or "_") and values either: Int, Float, String or Boolean |


<aside class="notice">
The isMain field in the product data structure marks a product as the main product in the group. This is particularly useful if you want to ensure a specific version
of the product name, images, etc ... to show up on the product page by default.  
</aside>



## ProductGroup

> Sample valid ProductGroup object:


```json
{
    "Code": 200,
    "Success": true,
    "Message": "Product Group Found ...",
    "Time": 1585609244627387726,
    "Response": {
        "GroupID": "MSLAPS2",
        "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
        "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
        "RegularPriceMin": 2000,
        "RegularPriceMax": 5000,
        "PromotionPriceMin": 1500,
        "PromotionPriceMax": 3000,
        "Skus": [
            "B07K3BHGL3",
            "B07K3BHGL4",
            "B07K3BHGL5"
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
        "Updated": 1585602487817024375,
        "Products": {
            "B07K3BHGL3": {
                "Sku": "B07K3BHGL3",
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                "GroupID": "MSLAPS2",
                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                "RegularPrice": 5000,
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
                "Quantity": 0,
                "Category": [
                    "Computers & Tablets>Laptops"
                ],
                "Color": "Black",
                "Brand": "Microsoft",
                "Size": "13.5 inches",
                "Active": true,
                "Selectors": {
                    "Display Size": "13.5 inches",
                    "Memory Speed": "1 GHz",
                    "RAM": "16 GB"
                },
                "Attributes": {
                    "ASIN": "B07K3BHGL3",
                    "Batteries": "1",
                    "Color": "Black",
                    "Date First Available": "Nov. 4 2018",
                    "Flash Memory Size": "512.00",
                    "Item Weight": "1.28 Kg",
                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                    "Item model number": "DAL-00092",
                    "Number of USB 2 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac"
                },
                "IsMain": true,
                "Currency": "CDN",
                "ComplementaryProducts": null,
                "Updated": 1585602487747404715
            },
            "B07K3BHGL4": {
                "Sku": "B07K3BHGL4",
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                "GroupID": "MSLAPS2",
                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                "RegularPrice": 2000,
                "PromotionPrice": 1500,
                "Images": [
                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                ],
                "SearchKeywords": [
                    "Laptop",
                    "Microsoft",
                    "Surface"
                ],
                "Quantity": 0,
                "Category": [
                    "Computers & Tablets>Laptops"
                ],
                "Color": "Black",
                "Brand": "Microsoft",
                "Size": "13.5 inches",
                "Active": true,
                "Selectors": {
                    "Display Size": "15 inches",
                    "Memory Speed": "4 GHz",
                    "RAM": "32 GB"
                },
                "Attributes": {
                    "ASIN": "B07K3BHGL3",
                    "Batteries": "1",
                    "Color": "Black",
                    "Date First Available": "Nov. 4 2018",
                    "Flash Memory Size": "512.00",
                    "Item Weight": "1.28 Kg",
                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                    "Item model number": "DAL-00092",
                    "Number of USB 2 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac"
                },
                "IsMain": true,
                "Currency": "CDN",
                "ComplementaryProducts": null,
                "Updated": 1585602582118590850
            },
            "B07K3BHGL5": {
                "Sku": "B07K3BHGL5",
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                "GroupID": "MSLAPS2",
                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                "RegularPrice": 3300,
                "PromotionPrice": 3000,
                "Images": [
                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                ],
                "SearchKeywords": [
                    "Laptop",
                    "Microsoft",
                    "Surface"
                ],
                "Quantity": 0,
                "Category": [
                    "Computers & Tablets>Laptops"
                ],
                "Color": "Black",
                "Brand": "Microsoft",
                "Size": "13.5 inches",
                "Active": true,
                "Selectors": {
                    "Display Size": "15 inches",
                    "Memory Speed": "2.88 GHz",
                    "RAM": "64 GB"
                },
                "Attributes": {
                    "ASIN": "B07K3BHGL3",
                    "Batteries": "1",
                    "Color": "Black",
                    "Date First Available": "Nov. 4 2018",
                    "Flash Memory Size": "512.00",
                    "Item Weight": "1.28 Kg",
                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                    "Item model number": "DAL-00092",
                    "Number of USB 2 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac"
                },
                "IsMain": true,
                "Currency": "CDN",
                "ComplementaryProducts": null,
                "Updated": 1585607727430437732
            }
        },
        "Selectors": {
            "Display Size": [
                "13.5 inches",
                "15 inches"
            ],
            "Memory Speed": [
                "1 GHz",
                "4 GHz",
                "2.88 GHz"
            ],
            "RAM": [
                "16 GB",
                "32 GB",
                "64 GB"
            ]
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
            "Number of USB 2 Ports": [
                "1"
            ],
            "Operating System": [
                "Windows 10 Home"
            ],
            "Processor Count": [
                "16"
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
        },
        "MainProductSKU": "B07K3BHGL3",
        "ComplementaryProducts": null,
        "CumulativeReviewStars": 0,
        "CumulativeReviewCount": 0
    }
}
```


ProductGroup objects are created and maintained by the platform. These objects have a very important purpose to solve - To optimize search, recommendation and other platform features by logically and automatically grouping a set of similar products. A group of similar products in your catalog could be same product with certain variations. Example: A shirt with same style but different colors and sizes logically forms a product group. Krama AI platform automatically groups similar products into ProductGroup objects. The platform does it using the **GroupID** field in your product object. **Product objects with same GroupID are grouped together in a ProductGroup object** 

The platform does more than what you expect. It groups these Product objects into ProductGroup objects and also aggregates product specific fields. ProductObjects also have price ranges computed using the RegularPrice and PromotionPrice fields in the Product objects. Attributes field map keys are also aggregated into an array of unique values. This is especially helpful in making faceted search requests. The platform takes care of the ProductGroup objects and builds and maintain these objects using the data from the Product objects - **in real-time**. 

Please find the field definitions, types, and constraints below:


|   Field                   |   Type         |     Short Description                                                                                                |
|---------------------------|----------------|----------------------------------------------------------------------------------------------------------------------|
|  GroupID                  |   String       | Unique product identifier                                                                                            |
|  Name                     |   String       | Name of the products in the group                                                                                    |
|  Description              |   String       | Product description in the product group                                                                             |
|  RegularPriceMin          |   Float        | Min value of the range computed of RegularPrice over the product group                                               |
|  RegularPriceMax          |   Float        | Max value of the range computed of RegularPrice over the product group                                               |
|  PromotionPriceMin        |   Float        | Min value of the range computed of PromotionPrice over the product group                                             |
|  PromotionPriceMax        |   Float        | Max value of the range computed of PromotionPrice over the product group                                             |
|  Skus                     |   String[]     | Aggregated array of product SKUs in the product group                                                                |
|  Images                   |   String[]     | Images from the main product in the product group (product where isMain == true)                                     |
|  SearchKeywords           |   String[]     | Aggregated array of unique search keywords from all the products in the group                                        |
|  Category                 |   String[]     | Aggregated array of unique categories from all the products in the group                                             |
|  Colors                   |   String[]     | Aggregated array of unique colors from all the products in the group                                                 |
|  Brands                   |   String[]     | Aggregated array of unique brands from all the products in the group                                                 |
|  Sizes                    |   String[]     | Aggregated array of unique sizes from all the products in the group                                                  |
|  Active                   |   boolean      | Flag to mark the group as active. If all the products are inactive group gets inactive else it is marked as active   |
|  Currency                 |   String       | Currency as mentioned in the product group.                                                                          |
|  Updated                  |   Integer      | Unix timestamp of the last time when the product group was updated                                                   |
|  Products                 |   Map{k,v}     | Map of products in the group. Key is the product SKU, value is the Product object                                    |
|  Attributes               |   Map{k,v}     | Aggregated Map of custom product attributes. Unique values are grouped in arrays for each key                        |
|  Selectors                |   Map{k,v}     | All those fields that help in identification of a specific variation of a product in the group                       |
|  CumulativeReviewStars    |   Float        | Cumulative Average of star rating of the product group                                                               |
|  CumulativeReviewCount    |   Integer      | Cumulative count of the reviews on the product group                                                                 | 




## Customer

> Sample valid customer object:

```json
{
  "Active": true,
  "FirstName": "Tom",
  "LastName": "Hanks",
  "Email": "tom.hanks@gmail.com",
  "PhoneNumbers": [
    "000-000-0000"
  ],
  "Password": "password",
  "AddressBook": [
    {
      "FirstName": "Tom",
      "LastName": "Hanks",
      "AddressLineOne": "101 Broad St",
      "AddressLineTwo": "",
      "City": "Santa Barbara",
      "State": "California",
      "Country": "United States",
      "Pincode": "00000",
      "Default": true
    }
  ],
  "PaymentOptions": [
    {
      "Name": "TOM HANKS",
      "CardNumber": "0000000000000000",
      "CardExpiryMM": "01",
      "CardExpiryYY": "20",
      "SecurityCode": "000",
      "ZipCode": "00000",
      "Default": true,
      "SaveInformation": true
    }
  ],
  "WishList": [
    "83947DSDS",
    "84378DFDW"
  ],
  "SaveForLater": [
    "FSDF3434",
    "ERF4432D"
  ]
}
```

Please find customer object fields and constraints/rules associated with each, below:


|   Field          |   Type         |     Short Description                                         |    Constraints                                                          |
|------------------|----------------|---------------------------------------------------------------|-------------------------------------------------------------------------|
| Active           | Boolean        | Flag to mark if the customer is an active customer            | Mandatory, boolean value - either true or false                         |
| CustomerID       | String         | Platform generated customer identifier                        | Mandatory, generated automatically                                      |
| FirstName        | String         | First name of the customer                                    | Mandatory, less than 100 characters                                     |
| LastName         | String         | Last name of the customer                                     | Mandatory, less than 100 characters                                     |
| Email            | String         | Email address. Used as primary field to identify a customer   | Mandatory, Valid email address                                          |
| PhoneNumbers     | String[]       | An array of phone numbers                                     | Multiple, A customer can have at most 10 phone numbers                  |
| Password         | String         | Customer password                                             | Mandatory, Cannot have less than 5 or more than 1024 characters         |
| AddressBook      | Address[]      | Object containing customer's valid addresses                  | Multiple, A customer can have at most 10 adresses at a time             |
| PaymentOptions   | PaymentOption[]| Object containing customer's valid payment information        | Multiple, A customer can have at most 50 payment options at a time      |
| WishList         | String[]       | An array of product SKUs                                      | Multiple, At a time a customer can have at most 1000 SKUs               |
| SaveForLater     | String[]       | An array of product SKUs                                      | Multiple, At a time a customer can have at most 1000 SKUs               |




## Address

> Sample valid Address object:

```json
{
      "FirstName": "Tom",
      "LastName": "Hanks",
      "AddressLineOne": "101 Broad St",
      "AddressLineTwo": "",
      "City": "Santa Barbara",
      "State": "California",
      "Country": "United States",
      "Pincode": "00000",
      "Default": true
}
```

Address object fields, definitions, and constraints below:


|   Field          |   Type         |     Short Description                                         |    Constraints                                                          |
|------------------|----------------|---------------------------------------------------------------|-------------------------------------------------------------------------|
| FirstName        | String         | First name of the address                                     | Mandatory, less than 100 characters                                     |
| LastName         | String         | Last name of the address                                      | Mandatory, less than 100 characters                                     |
| AddressLineOne   | String         | First line for address description                            | Mandatory, less than 200 characters                                     |
| AddressLineTwo   | String         | Second line for address description                           | Optional, less than 200 characters                                      |
| City             | String         | Address city                                                  | Mandatory, less than 100 characters                                     |
| State            | String         | Address state                                                 | Mandatory, less than 100 characters                                     |
| Country          | String         | Address country                                               | Mandatory, less than 100 characters                                     |
| Pincode          | String         | Addres pincode                                                | Mandatory, less than 10 characters                                      |
| Default          | Boolean        | Is the address the default address?                           | Boolean value - either true or false                                    |



## PaymentOption

> Sample valid PaymentOption object:

```json
{
      "Name": "TOM HANKS",
      "CardNumber": "0000000000000000",
      "CardExpiryMM": "01",
      "CardExpiryYY": "20",
      "SecurityCode": "000",
      "ZipCode": "00000",
      "Default": true,
      "SaveInformation": true
}
```


PaymentOption object fields, definitions, and constraints below:


|   Field          |   Type         |     Short Description                                         |    Constraints                                                          |
|------------------|----------------|---------------------------------------------------------------|-------------------------------------------------------------------------|
| Name             | String         | Name associated with credit card                              | Mandatory, less than 100 characters                                     |
| CardNumber       | String         | Card number associated with payment option                    | Mandatory, 16 digits                                                    |
| CardExpiryMM     | String         | Card expiry month                                             | Mandatory, 2 digits greater than 00 and less than equal to 12           |
| CardExpiryYY     | String         | Card expiry year                                              | Mandatory, 2 digits                                                     |
| SecurityCode     | String         | Card security code                                            | Mandatory, 3 digits                                                     |
| ZipCode          | String         | Card pincode                                                  | Mandatory, valid pincode (based on country)                             |
| Default          | Boolean        | Flag to make if the payment information is the default option | Mandatory, boolean - either true or false                               |
| SaveInformation  | Boolean        | Flag to mark the payment information to be saved for use later| Mandatory, boolean - either true or false                               |                                          

## Response

> Sample response object:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Search Result ...",
    "Time": 1579463722416637832,
    "Response": {
        "count": 1,
        "results": {
            "0": {
                "Skus": [
                    "B07K3BHGL4"
                ],
                "Images": [
                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                ],
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)"
            }
        }
    }
}
```

Every request to this API gets a standard response object. Details on the fields of the response object is elaborated in the table below.


|  Key              |    Description                                                                        |
|-------------------|---------------------------------------------------------------------------------------|
| Code              | [HTTP Response code](#response-codes) for the request                                 |
| Success           | Flag that tells if the request was successful or not                                  |
| Message           | Message for additional information                                                    |
| Time              | Unix timestamp of the response                                                        |
| Response          | Response object containing response information (variable field for each request)     |


# API access

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
        "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjeHMiOiJDb250ZUFtZXJpY2EiLCJleHAiOjE1NzkxMDY5NTQsImlhdCI6MTU3OTAyNjk1NCwibmJmIjoxNTc5MDI2ODU0LCJ1aWQiOiIwMjQ0Zjg1NS1jMWQ3LTQyNGYtOWI5OS04NGZmYWNiYzYwOGUifQ.6IhX3X321NlZFtSSf3JUPisD7fTxqeVrCpHQ6WDDgIk",
        "ValidForSeconds": 80000
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

This endpoint gets you your API access token. You need to send your customer ID and the API key that we provided you for using our platform. Upon receiving your valid credentials, the API will respond with a token with additional information including the **ValidForSeconds** key which tells you how long this access token is valid for. Please set **Authorization** field in the HTTP request header to the **token** in your API calls. 


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
You do not need to invoke login too often. Please include the token that you receive upon a successful login in your subsequent API calls until the token expires. You can keep a track of token expiry using the ValidForSeconds field in the response. 
</aside>

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |
| Token             | Included in response object that should be included in subsequent API calls   |
| ValidForSeconds   | Included in response object that tells the validity of access token in seconds|


# Catalog API

## Add a new product

> Sample HTTP request body:

```json
{
  "Sku": "B07K3BHGL5",
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
  "RegularPrice": 3300,
  "PromotionPrice": 3000,
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
  "Selectors": {
    "RAM": "64 GB",
    "Display Size": "15 inches",
    "Memory Speed": "2.88 GHz"
  },
  "Attributes": {
    "Wireless Standard": "802.11ac",
    "Number of USB 2 Ports": "1",
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
    "Time": 1585607727604209821,
    "Response": {
        "Sku": "B07K3BHGL5",
        "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
        "GroupID": "MSLAPS2",
        "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
        "RegularPrice": 3300,
        "PromotionPrice": 3000,
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
        "Selectors": {
            "Display Size": "15 inches",
            "Memory Speed": "2.88 GHz",
            "RAM": "64 GB"
        },
        "Attributes": {
            "ASIN": "B07K3BHGL3",
            "Batteries": "1",
            "Color": "Black",
            "Date First Available": "Nov. 4 2018",
            "Flash Memory Size": "512.00",
            "Item Weight": "1.28 Kg",
            "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
            "Item model number": "DAL-00092",
            "Number of USB 2 Ports": "1",
            "Operating System": "Windows 10 Home",
            "Processor Count": "16",
            "Series": "Surface Laptop 2",
            "Shipping Weight": "2.2 kg",
            "Wireless Standard": "802.11ac"
        },
        "IsMain": true,
        "Currency": "CDN",
        "ComplementaryProducts": null,
        "Updated": 1585607727430437732
    }
}
```

Use this API endpoint to add a new product in the products collection. When a product is added in the products collection it also gest added in product group collection. If the product group with the matching groupID is missing, a new product group is formed. Search index and the cache are also automatically updated with a valid call to this API endpoint. 


### HTTP Request URL

`POST https://api.krama.ai/catalog/{API version}/products`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |

### HTTP Request Body Parameters

|   Field          |   Type         |     Short Description                                         |    Constraints                                                          |
|------------------|----------------|---------------------------------------------------------------|-------------------------------------------------------------------------|
|  Sku             |   String       | Unique product identifier                                     | Mandatory, Unique, Alphanumeric, Less than 50 characters                |
|  Name            |   String       | Name of the product                                           | Mandatory, Less than 100 characters                                     |
|  GroupID         |   String       | Product group identifier. More details in other section       | Mandatory, Less than 50 characters                                      |
|  Description     |   String       | Product description field                                     | Mandatory, less than 10240 characters                                   |
|  RegularPrice    |   Float        | Product's everyday price                                      | Mandatory, Cannot be negative                                           |
|  PromotionPrice  |   Float        | Product's promotion price. Typically less than regular price. | Mandatory, Cannot be negative                                           |
|  Images          |   String       | Product image links                                           | Mandatory, need to be valid URLs, cannot be more than 100 URLs          |
|  SearchKeywords  |   String[]     | Product search keywords                                       | Mandatory, cannot be more than 100 search keywords                      |
|  Quantity        |   Integer      | Product stock quantity field                                  | Mandatory, cannot be negative                                           |
|  Category        |   String[]     | Product category path. Please see product object example      | Mandatory, '>' separated category path                                  |
|  Color           |   String       | Product color field                                           | Optional, cannot be greater than 100 characters                         |         
|  Brand           |   String       | Product brand field                                           | Optional, cannot be greater than 100 characters                         |
|  Size            |   String       | Product size field                                            | Optional, cannot be greater than 100 characters                         |
|  Active          |   Boolean      | Field to mark product available for sale                      | Boolean - either true or false                                          |
|  IsMain          |   Boolean      | Field to mark the product as a main product in the group      | Boolean - either true of false                                          |
|  Currency        |   String       | Product purchase currency                                     | Either - "USD", "CAD", "CDN", "INR", "GBP" or "EUR" (for now!)          |
|  Attributes      |   Map{k,v}     | Product custom attributes that fit your needs                 | keys should be strings and values either: int, float, string or boolean |  
|  Selectors       |   Map{k,v}     | Those attributes that are used to identify a variation of a product in the group | Keys should be strings (alphanumeric with single space and "-" or "_") and values either: Int, Float, String or Boolean |

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
    "Time": 1585608587275689785,
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
        "Selectors": {
            "Display Size": "13.5 inches",
            "Memory Speed": "1 GHz",
            "RAM": "16 GB"
        },
        "Attributes": {
            "ASIN": "B07K3BHGL3",
            "Batteries": "1",
            "Color": "Black",
            "Date First Available": "Nov. 4 2018",
            "Flash Memory Size": "512.00",
            "Item Weight": "1.28 Kg",
            "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
            "Item model number": "DAL-00092",
            "Number of USB 2 Ports": "1",
            "Operating System": "Windows 10 Home",
            "Processor Count": "16",
            "Series": "Surface Laptop 2",
            "Shipping Weight": "2.2 kg",
            "Wireless Standard": "802.11ac"
        },
        "IsMain": true,
        "Currency": "CDN",
        "ComplementaryProducts": null,
        "Updated": 1585602487747404715
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

When you want to get a specific product you can use this endpoint. All you need to pass is your access token and the SKU. This endpoint is cached for efficiency but also ensures that updated product data is served when the product update is invoked for this product. 


### HTTP Request URL

`GET https://api.krama.ai/catalog/{API version}/products/{SKU}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |

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
| Selectors         | All those product attributes which help choose a specific product variation   |


## Get products (bulk)

> Sample HTTP request body:

```json
{
  "Skus" : ["B07K3BHGL3"]
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": false,
    "Message": "Products Found ...",
    "Time": 1585608824004904929,
    "Response": [
        {
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
            "Quantity": 0,
            "Category": [
                "Computers & Tablets>Laptops"
            ],
            "Color": "Black",
            "Brand": "Microsoft",
            "Size": "13.5 inches",
            "Active": true,
            "Selectors": {
                "Display Size": "13.5 inches",
                "Memory Speed": "1 GHz",
                "RAM": "16 GB"
            },
            "Attributes": {
                "ASIN": "B07K3BHGL3",
                "Batteries": "1",
                "Color": "Black",
                "Date First Available": "Nov. 4 2018",
                "Flash Memory Size": "512.00",
                "Item Weight": "1.28 Kg",
                "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                "Item model number": "DAL-00092",
                "Number of USB 2 Ports": "1",
                "Operating System": "Windows 10 Home",
                "Processor Count": "16",
                "Series": "Surface Laptop 2",
                "Shipping Weight": "2.2 kg",
                "Wireless Standard": "802.11ac"
            },
            "IsMain": true,
            "Currency": "CDN",
            "ComplementaryProducts": null,
            "Updated": 1585602487747404715
        }
    ]
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Product Not Found ...",
    "Time": 1579846902585974734,
    "Response": null
}
```

Use this API endpoint to get an array of products by passing an array of SKUs. 


### HTTP Request URL

`PUT https://api.krama.ai/catalog/{API version}/bulk/products`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

| Parameter     | Description                                      |
|---------------|--------------------------------------------------|
| Skus          | string[] - An array of product SKUs              |


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
| Selectors         | All those product attributes which help choose a specific product variation   |


## Update a product

> Sample HTTP request body:

```json
{
    "Sku": "B07K3BHGL3",
    "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
    "GroupID": "MSLAPS2",
    "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
    "RegularPrice": 5000,
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
    "Quantity": 0,
    "Category": [
        "Computers & Tablets>Laptops"
    ],
    "Color": "Black",
    "Brand": "Microsoft",
    "Size": "13.5 inches",
    "Active": true,
    "Selectors": {
        "Display Size": "13.5 inches",
        "Memory Speed": "1 GHz",
        "RAM": "16 GB"
    },
    "Attributes": {
        "ASIN": "B07K3BHGL3",
        "Batteries": "1",
        "Color": "Black",
        "Date First Available": "Nov. 4 2018",
        "Flash Memory Size": "512.00",
        "Item Weight": "1.28 Kg",
        "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
        "Item model number": "DAL-00092",
        "Number of USB 2 Ports": "1",
        "Operating System": "Windows 10 Home",
        "Processor Count": "16",
        "Series": "Surface Laptop 2",
        "Shipping Weight": "2.2 kg",
        "Wireless Standard": "802.11ac"
    },
    "IsMain": true,
    "Currency": "CDN",
    "ComplementaryProducts": null,
    "Updated": 1585602487747404715
}
```

> Sample valid API response:

```json
{
    "Code": 202,
    "Success": true,
    "Message": "Product Updated ...",
    "Time": 1585608990922156125,
    "Response": {
        "Sku": "B07K3BHGL3",
        "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
        "GroupID": "MSLAPS2",
        "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
        "RegularPrice": 5000,
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
        "Quantity": 0,
        "Category": [
            "Computers & Tablets>Laptops"
        ],
        "Color": "Black",
        "Brand": "Microsoft",
        "Size": "13.5 inches",
        "Active": true,
        "Selectors": {
            "Display Size": "13.5 inches",
            "Memory Speed": "1 GHz",
            "RAM": "16 GB"
        },
        "Attributes": {
            "ASIN": "B07K3BHGL3",
            "Batteries": "1",
            "Color": "Black",
            "Date First Available": "Nov. 4 2018",
            "Flash Memory Size": "512.00",
            "Item Weight": "1.28 Kg",
            "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
            "Item model number": "DAL-00092",
            "Number of USB 2 Ports": "1",
            "Operating System": "Windows 10 Home",
            "Processor Count": "16",
            "Series": "Surface Laptop 2",
            "Shipping Weight": "2.2 kg",
            "Wireless Standard": "802.11ac"
        },
        "IsMain": true,
        "Currency": "CDN",
        "ComplementaryProducts": null,
        "Updated": 1585602487747404715
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

Use this API endpoint to update your product information in the catalog. For now you need to pass the entire product object with updated parts (this functionality will be improved very soon). When you hit this API endpoint, the data in the products collection gets updated, product group data also gets updated automatically, search index is also updated and the cache entry for this product is removed first and updated on the next GET call. 


### HTTP Request URL

`PUT https://api.krama.ai/catalog/{API version}/products/{SKU}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

|   Field          |   Type         |     Short Description                                         |    Constraints                                                          |
|------------------|----------------|---------------------------------------------------------------|-------------------------------------------------------------------------|
|  Sku             |   String       | Unique product identifier                                     | Mandatory, Unique, Alphanumeric, Less than 50 characters                |
|  Name            |   String       | Name of the product                                           | Mandatory, Less than 100 characters                                     |
|  GroupID         |   String       | Product group identifier. More details in other section       | Mandatory, Less than 50 characters                                      |
|  Description     |   String       | Product description field                                     | Mandatory, less than 10240 characters                                   |
|  RegularPrice    |   Float        | Product's everyday price                                      | Mandatory, Cannot be negative                                           |
|  PromotionPrice  |   Float        | Product's promotion price. Typically less than regular price. | Mandatory, Cannot be negative                                           |
|  Images          |   String       | Product image links                                           | Mandatory, need to be valid URLs, cannot be more than 100 URLs          |
|  SearchKeywords  |   String[]     | Product search keywords                                       | Mandatory, cannot be more than 100 search keywords                      |
|  Quantity        |   Integer      | Product stock quantity field                                  | Mandatory, cannot be negative                                           |
|  Category        |   String[]     | Product category path. Please see product object example      | Mandatory, '>' separated category path                                  |
|  Color           |   String       | Product color field                                           | Optional, cannot be greater than 100 characters                         |         
|  Brand           |   String       | Product brand field                                           | Optional, cannot be greater than 100 characters                         |
|  Size            |   String       | Product size field                                            | Optional, cannot be greater than 100 characters                         |
|  Active          |   Boolean      | Field to mark product available for sale                      | Boolean - either true or false                                          |
|  IsMain          |   Boolean      | Field to mark the product as a main product in the group      | Boolean - either true of false                                          |
|  Currency        |   String       | Product purchase currency                                     | Either - "USD", "CAD", "CDN", "INR", "GBP" or "EUR" (for now!)          |
|  Attributes      |   Map{k,v}     | Product custom attributes that fit your needs                 | keys should be strings and values either: int, float, string or boolean |  
|  Selectors       |   Map{k,v}     | Those attributes that are used to identify a variation of a product in the group | Keys should be strings (alphanumeric with single space and "-" or "_") and values either: Int, Float, String or Boolean |

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
| Selectors         | All those product attributes which help choose a specific product variation   |


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
ensuring that the product groups and the search index are synced as well in real-time. The response object contains three lists - the list of updated skus 
(essentially all the skus for which the prices were updated), a list of skus which were not updated (most likely that you submitted same old prices), and
finally a list of skus that were not found. You can use this information to check if your update request was executed as per your expectations or not. 


### HTTP Request URL

`PUT https://api.krama.ai/catalog/{API version}/products/price/update`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
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

Use this API endpoint to update your product quantity. You can use this endpoint to update inventory for multiple products. Please find the API request details below. The API updates product groups as well as search index. The response object contains three lists - a list that gives you the skus of those
products where the product quantities were updated. The response object also provides you with list of those skus where there was no change and also
those skus which were not found in the catalog. You can use the information from the response to verify that the changes were done as per your expectation. 


### HTTP Request URL

`PUT https://api.krama.ai/catalog/{API version}/products/inventory/update`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Body Parameters

|    Parameter    |          Constraints                    |        Description                                           |
|-----------------|-----------------------------------------|--------------------------------------------------------------|
|Quantity         |                                         |  The map containing skus-quantity mappings                   |
|{sku : quantity} | {unique identifier, int non-negative}   |  sku-quantity mappings in the Quantity object                |


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
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |

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
    "Time": 1585609244627387726,
    "Response": {
        "GroupID": "MSLAPS2",
        "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
        "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
        "RegularPriceMin": 2000,
        "RegularPriceMax": 5000,
        "PromotionPriceMin": 1500,
        "PromotionPriceMax": 3000,
        "Skus": [
            "B07K3BHGL3",
            "B07K3BHGL4",
            "B07K3BHGL5"
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
        "Updated": 1585602487817024375,
        "Products": {
            "B07K3BHGL3": {
                "Sku": "B07K3BHGL3",
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                "GroupID": "MSLAPS2",
                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                "RegularPrice": 5000,
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
                "Quantity": 0,
                "Category": [
                    "Computers & Tablets>Laptops"
                ],
                "Color": "Black",
                "Brand": "Microsoft",
                "Size": "13.5 inches",
                "Active": true,
                "Selectors": {
                    "Display Size": "13.5 inches",
                    "Memory Speed": "1 GHz",
                    "RAM": "16 GB"
                },
                "Attributes": {
                    "ASIN": "B07K3BHGL3",
                    "Batteries": "1",
                    "Color": "Black",
                    "Date First Available": "Nov. 4 2018",
                    "Flash Memory Size": "512.00",
                    "Item Weight": "1.28 Kg",
                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                    "Item model number": "DAL-00092",
                    "Number of USB 2 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac"
                },
                "IsMain": true,
                "Currency": "CDN",
                "ComplementaryProducts": null,
                "Updated": 1585602487747404715
            },
            "B07K3BHGL4": {
                "Sku": "B07K3BHGL4",
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                "GroupID": "MSLAPS2",
                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                "RegularPrice": 2000,
                "PromotionPrice": 1500,
                "Images": [
                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                ],
                "SearchKeywords": [
                    "Laptop",
                    "Microsoft",
                    "Surface"
                ],
                "Quantity": 0,
                "Category": [
                    "Computers & Tablets>Laptops"
                ],
                "Color": "Black",
                "Brand": "Microsoft",
                "Size": "13.5 inches",
                "Active": true,
                "Selectors": {
                    "Display Size": "15 inches",
                    "Memory Speed": "4 GHz",
                    "RAM": "32 GB"
                },
                "Attributes": {
                    "ASIN": "B07K3BHGL3",
                    "Batteries": "1",
                    "Color": "Black",
                    "Date First Available": "Nov. 4 2018",
                    "Flash Memory Size": "512.00",
                    "Item Weight": "1.28 Kg",
                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                    "Item model number": "DAL-00092",
                    "Number of USB 2 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac"
                },
                "IsMain": true,
                "Currency": "CDN",
                "ComplementaryProducts": null,
                "Updated": 1585602582118590850
            },
            "B07K3BHGL5": {
                "Sku": "B07K3BHGL5",
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                "GroupID": "MSLAPS2",
                "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                "RegularPrice": 3300,
                "PromotionPrice": 3000,
                "Images": [
                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                ],
                "SearchKeywords": [
                    "Laptop",
                    "Microsoft",
                    "Surface"
                ],
                "Quantity": 0,
                "Category": [
                    "Computers & Tablets>Laptops"
                ],
                "Color": "Black",
                "Brand": "Microsoft",
                "Size": "13.5 inches",
                "Active": true,
                "Selectors": {
                    "Display Size": "15 inches",
                    "Memory Speed": "2.88 GHz",
                    "RAM": "64 GB"
                },
                "Attributes": {
                    "ASIN": "B07K3BHGL3",
                    "Batteries": "1",
                    "Color": "Black",
                    "Date First Available": "Nov. 4 2018",
                    "Flash Memory Size": "512.00",
                    "Item Weight": "1.28 Kg",
                    "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                    "Item model number": "DAL-00092",
                    "Number of USB 2 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac"
                },
                "IsMain": true,
                "Currency": "CDN",
                "ComplementaryProducts": null,
                "Updated": 1585607727430437732
            }
        },
        "Selectors": {
            "Display Size": [
                "13.5 inches",
                "15 inches"
            ],
            "Memory Speed": [
                "1 GHz",
                "4 GHz",
                "2.88 GHz"
            ],
            "RAM": [
                "16 GB",
                "32 GB",
                "64 GB"
            ]
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
            "Number of USB 2 Ports": [
                "1"
            ],
            "Operating System": [
                "Windows 10 Home"
            ],
            "Processor Count": [
                "16"
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
        },
        "MainProductSKU": "B07K3BHGL3",
        "ComplementaryProducts": null,
        "CumulativeReviewStars": 0,
        "CumulativeReviewCount": 0
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

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |

### HTTP Request URL Parameters

| Parameter     | Description          |
|---------------|----------------------|
|PGID           | The product Group ID |

### HTTP Response

|  Key                  |    Description                                                                |
|-----------------------|-------------------------------------------------------------------------------|
| Code                  |  Response code for the request                                                |
| Success               |  Flag that tells if the request was successful                                |
| Message               |  Message for additional information                                           |
| Time                  |  Unix timestamp of the response                                               |
| Response              |  Response object containing response information                              |
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
| Selectors             |  All those product attributes which help choose a specific product variation  |


## Get product groups (bulk)

> Sample HTTP request body:

```json
{
  "Skus" : ["B07K3BHGL3"]
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": false,
    "Message": "Product Groups Found ...",
    "Time": 1585611290280810346,
    "Response": [
        {
            "GroupID": "MSLAPS2",
            "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
            "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
            "RegularPriceMin": 2000,
            "RegularPriceMax": 5000,
            "PromotionPriceMin": 1500,
            "PromotionPriceMax": 3000,
            "Skus": [
                "B07K3BHGL3",
                "B07K3BHGL4",
                "B07K3BHGL5"
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
            "Updated": 1585602487817024375,
            "Products": {
                "B07K3BHGL3": {
                    "Sku": "B07K3BHGL3",
                    "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                    "GroupID": "MSLAPS2",
                    "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                    "RegularPrice": 5000,
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
                    "Quantity": 0,
                    "Category": [
                        "Computers & Tablets>Laptops"
                    ],
                    "Color": "Black",
                    "Brand": "Microsoft",
                    "Size": "13.5 inches",
                    "Active": true,
                    "Selectors": {
                        "Display Size": "13.5 inches",
                        "Memory Speed": "1 GHz",
                        "RAM": "16 GB"
                    },
                    "Attributes": {
                        "ASIN": "B07K3BHGL3",
                        "Batteries": "1",
                        "Color": "Black",
                        "Date First Available": "Nov. 4 2018",
                        "Flash Memory Size": "512.00",
                        "Item Weight": "1.28 Kg",
                        "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                        "Item model number": "DAL-00092",
                        "Number of USB 2 Ports": "1",
                        "Operating System": "Windows 10 Home",
                        "Processor Count": "16",
                        "Series": "Surface Laptop 2",
                        "Shipping Weight": "2.2 kg",
                        "Wireless Standard": "802.11ac"
                    },
                    "IsMain": true,
                    "Currency": "CDN",
                    "ComplementaryProducts": null,
                    "Updated": 1585602487747404715
                },
                "B07K3BHGL4": {
                    "Sku": "B07K3BHGL4",
                    "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                    "GroupID": "MSLAPS2",
                    "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                    "RegularPrice": 2000,
                    "PromotionPrice": 1500,
                    "Images": [
                        "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                        "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                    ],
                    "SearchKeywords": [
                        "Laptop",
                        "Microsoft",
                        "Surface"
                    ],
                    "Quantity": 0,
                    "Category": [
                        "Computers & Tablets>Laptops"
                    ],
                    "Color": "Black",
                    "Brand": "Microsoft",
                    "Size": "13.5 inches",
                    "Active": true,
                    "Selectors": {
                        "Display Size": "15 inches",
                        "Memory Speed": "4 GHz",
                        "RAM": "32 GB"
                    },
                    "Attributes": {
                        "ASIN": "B07K3BHGL3",
                        "Batteries": "1",
                        "Color": "Black",
                        "Date First Available": "Nov. 4 2018",
                        "Flash Memory Size": "512.00",
                        "Item Weight": "1.28 Kg",
                        "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                        "Item model number": "DAL-00092",
                        "Number of USB 2 Ports": "1",
                        "Operating System": "Windows 10 Home",
                        "Processor Count": "16",
                        "Series": "Surface Laptop 2",
                        "Shipping Weight": "2.2 kg",
                        "Wireless Standard": "802.11ac"
                    },
                    "IsMain": true,
                    "Currency": "CDN",
                    "ComplementaryProducts": null,
                    "Updated": 1585602582118590850
                },
                "B07K3BHGL5": {
                    "Sku": "B07K3BHGL5",
                    "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)",
                    "GroupID": "MSLAPS2",
                    "Description": "Clean, elegant design — thin and light, starting at just 2.76 pounds, Surface Laptop 2 fits easily in your bag Choose from rich tone-on-tone color combinations: Platinum, Burgundy, and Cobalt Blue, plus an all-new finish in classic Matte Black Improved speed and performance to do what you want, with the latest 8th Generation Intel Core processor",
                    "RegularPrice": 3300,
                    "PromotionPrice": 3000,
                    "Images": [
                        "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                        "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                    ],
                    "SearchKeywords": [
                        "Laptop",
                        "Microsoft",
                        "Surface"
                    ],
                    "Quantity": 0,
                    "Category": [
                        "Computers & Tablets>Laptops"
                    ],
                    "Color": "Black",
                    "Brand": "Microsoft",
                    "Size": "13.5 inches",
                    "Active": true,
                    "Selectors": {
                        "Display Size": "15 inches",
                        "Memory Speed": "2.88 GHz",
                        "RAM": "64 GB"
                    },
                    "Attributes": {
                        "ASIN": "B07K3BHGL3",
                        "Batteries": "1",
                        "Color": "Black",
                        "Date First Available": "Nov. 4 2018",
                        "Flash Memory Size": "512.00",
                        "Item Weight": "1.28 Kg",
                        "Item dimensions L x W x H": "17.8 x 12.7 x 15.2 cm",
                        "Item model number": "DAL-00092",
                        "Number of USB 2 Ports": "1",
                        "Operating System": "Windows 10 Home",
                        "Processor Count": "16",
                        "Series": "Surface Laptop 2",
                        "Shipping Weight": "2.2 kg",
                        "Wireless Standard": "802.11ac"
                    },
                    "IsMain": true,
                    "Currency": "CDN",
                    "ComplementaryProducts": null,
                    "Updated": 1585607727430437732
                }
            },
            "Selectors": {
                "Display Size": [
                    "13.5 inches",
                    "15 inches"
                ],
                "Memory Speed": [
                    "1 GHz",
                    "4 GHz",
                    "2.88 GHz"
                ],
                "RAM": [
                    "16 GB",
                    "32 GB",
                    "64 GB"
                ]
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
                "Number of USB 2 Ports": [
                    "1"
                ],
                "Operating System": [
                    "Windows 10 Home"
                ],
                "Processor Count": [
                    "16"
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
            },
            "MainProductSKU": "B07K3BHGL3",
            "ComplementaryProducts": null,
            "CumulativeReviewStars": 0,
            "CumulativeReviewCount": 0
        }
    ]
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Product Group Not Found ...",
    "Time": 1579847199233752534,
    "Response": null
}
```

Use this API endpoint to get an array of product group objects by passing an array of SKUs. 


### HTTP Request URL

`PUT https://api.krama.ai/catalog/{API version}/bulk/productgroups`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

| Parameter     | Description                                      |
|---------------|--------------------------------------------------|
| Skus          | string[] - An array of product SKUs              |


### HTTP Response

|  Key                  |    Description                                                                |
|-----------------------|-------------------------------------------------------------------------------|
| Code                  |  Response code for the request                                                |
| Success               |  Flag that tells if the request was successful                                |
| Message               |  Message for additional information                                           |
| Time                  |  Unix timestamp of the response                                               |
| Response              |  Response object containing response information                              |
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
| Selectors             |  All those product attributes which help choose a specific product variation  |


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
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |

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



# Orders API

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
        "Number of USB 2 Ports": "1",
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
                    "Number of USB 2 Ports": "1",
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

Use this endpoint to create an order entry in the database. Details on the fields and the constraints are listed below. 


### HTTP Request URL

`POST https://api.krama.ai/orders/{API version}/orders`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
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
                    "Number of USB 2 Ports": "1",
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
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |

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
                    "Number of USB 2 Ports": "1",
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
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |

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
            "Number of USB 2 Ports": "1",
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
                    "Number of USB 2 Ports": "1",
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
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request URL Parameters

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




## Delete order

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
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |

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

## Quick search

> Sample HTTP request body:

```json
{
  "Query": " microsoft",
  "QueryFields": [
    "Name",
    "Sku"
  ],
  "ResponseFields": [
    "Name",
    "Skus",
    "Images"
  ],
  "From": 0,
  "To": 100
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Search Result ...",
    "Time": 1579233804527950065,
    "Response": {
        "count": 1,
        "results": {
            "0": {
                "Skus": [
                    "B07K3BHGL3"
                ],
                "Images": [
                    "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                    "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                ],
                "Name": "Microsoft DAL-00092 Surface Laptop 2 (Intel Core i7, 16GB RAM, 512 GB) - Black (Newest Version)"
            }
        }
    }
}
```

Use this API endpoint to search products in the search index. Use of this endpoint is recommended for quick search feature. 


### HTTP Request

`POST https://api.krama.ai/search/{API version}/quick`


### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

| Parameter             |               Description                                        |
|-----------------------|------------------------------------------------------------------|
| Query                 |  String, Query string                                            |
| QueryFields           |  String[], Multiple, query fields                                |
| ResponseFields        |  String[], Multiple, fields that requested in response           |
| From                  |  Integer, used for pagination returning products in range        |
| To                    |  Integer, used for pagination returning products in range        | 

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |






## Full Page search

> Sample HTTP request body:

```json
{
  "Query": "microsoft",
  "QueryFields": [
    "Name",
    "Sku"
  ],
  "ResponseFields": [
    "Attributes.Color",
    "Skus",
    "Images"
  ],
  "From": 0,
  "To": 100,
  "TermFacetFields": [
    "Brands"
  ],
  "RangeFacetFields": [
    {
      "RegularPriceMin": [
        {
          "from": 0,
          "to": 1000
        },
        {
          "from": 1001,
          "to": 2000
        },
        {
          "from": 2001,
          "to": 3000
        }
      ]
    }
  ]
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Search Result ...",
    "Time": 1579235619698744869,
    "Response": {
        "count": 1,
        "facets": {
            "Brands": {
                "doc_count_error_upper_bound": 0,
                "sum_other_doc_count": 0,
                "buckets": [
                    {
                        "key": "Microsoft xxx yyy",
                        "doc_count": 1
                    }
                ]
            },
            "RegularPriceMin": {
                "buckets": [
                    {
                        "key": "0.0-1000.0",
                        "from": 0.0,
                        "to": 1000.0,
                        "doc_count": 0
                    },
                    {
                        "key": "1001.0-2000.0",
                        "from": 1001.0,
                        "to": 2000.0,
                        "doc_count": 0
                    },
                    {
                        "key": "2001.0-3000.0",
                        "from": 2001.0,
                        "to": 3000.0,
                        "doc_count": 1
                    }
                ]
            }
        },
        "hits": [
            {
                "_score": 0.2876821,
                "_index": "ffaabbccdd.productgroups.index",
                "_type": "_doc",
                "_id": "MSLAPS2",
                "_seq_no": null,
                "_primary_term": null,
                "_source": {
                    "Skus": [
                        "B07K3BHGL3"
                    ],
                    "Images": [
                        "https://images-na.ssl-images-amazon.com/images/I/51JODZveCOL._SL1200_.jpg",
                        "https://images-na.ssl-images-amazon.com/images/I/511Kd0b1WxL._SL1200_.jpg"
                    ],
                    "Attributes": {
                        "Color": [
                            "Black"
                        ]
                    }
                }
            }
        ]
    }
}
```

Use this API endpoint to search products in the search index. Use of this endpoint is recommended for full-page search. 


### HTTP Request

`POST https://api.krama.ai/search/{API version}/fullpage`


### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

| Parameter                     |               Description                                        |
|-------------------------------|------------------------------------------------------------------|
| Query                         |  String,   Query string                                          |
| QueryFields                   |  String[], Multiple, query fields                                |
| ResponseFields                |  String[], Multiple, fields that requested in response           |
| From                          |  Integer,  Used for pagination returning products in range       |
| To                            |  Integer,  Used for pagination returning products in range       | 
| TermFacetFields               |  String[], Eligible fields for term faceting                     |
| RangeFacetFields              |  Complex,  Typically used for a numeric field like price         |

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |



# Customers API

## Get a customer

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": false,
    "Message": "Customer Found ...",
    "Time": 1579240042172013722,
    "Response": {
        "CustomerID": "127452a3-00b6-4a04-a03b-e1db919645cc",
        "Active": true,
        "FirstName": "Tom",
        "LastName": "Hanks",
        "Email": "tom.hanks@gmail.com",
        "PhoneNumbers": [
            "000-000-0000"
        ],
        "Password": "password",
        "AddressBook": [
            {
                "FirstName": "Tom",
                "LastName": "Hanks",
                "AddressLineOne": "101 Broad St",
                "AddressLineTwo": "",
                "City": "Santa Barbara",
                "State": "California",
                "Country": "United States",
                "Pincode": "00000",
                "Default": true
            }
        ],
        "PaymentOptions": [
            {
                "Name": "TOM HANKS",
                "CardNumber": "0000000000000000",
                "CardExpiryMM": "01",
                "CardExpiryYY": "20",
                "SecurityCode": "000",
                "ZipCode": "00000",
                "Default": true,
                "SaveInformation": true
            }
        ],
        "WishList": [
            "83947DSDS"
        ],
        "SaveForLater": [
            "FSDF3434",
            "ERF4432D"
        ],
        "Updated": 1579233562701600551
    }
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Customer Not Found ...",
    "Time": 1579148524152947154,
    "Response": null
}
```


Use this endpoint to get the customer object from the database. 


### HTTP Request URL

`GET https://api.krama.ai/customers/{API version}/customers/{CID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |


### HTTP Request URL Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| CID                   |  Customer identifier                                |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |


## Add a customer

> Sample HTTP request body:

```json
{
  "Active": true,
  "FirstName": "Tom",
  "LastName": "Hanks",
  "Email": "tom.hanks@gmail.com",
  "PhoneNumbers": [
    "000-000-0000"
  ],
  "Password": "password",
  "AddressBook": [
    {
      "FirstName": "Tom",
      "LastName": "Hanks",
      "AddressLineOne": "101 Broad St",
      "AddressLineTwo": "",
      "City": "Santa Barbara",
      "State": "California",
      "Country": "United States",
      "Pincode": "00000",
      "Default": true
    }
  ],
  "PaymentOptions": [
    {
      "Name": "TOM HANKS",
      "CardNumber": "0000000000000000",
      "CardExpiryMM": "01",
      "CardExpiryYY": "20",
      "SecurityCode": "000",
      "ZipCode": "00000",
      "Default": true,
      "SaveInformation": true
    }
  ],
  "WishList": [
    "83947DSDS",
    "84378DFDW"
  ],
  "SaveForLater": [
    "FSDF3434",
    "ERF4432D"
  ]
}
```

> Sample valid API response:

```json
{
    "Code": 201,
    "Success": true,
    "Message": "Customer Added ...",
    "Time": 1579233420045379570,
    "Response": {
        "CustomerID": "127452a3-00b6-4a04-a03b-e1db919645cc",
        "Active": true,
        "FirstName": "Tom",
        "LastName": "Hanks",
        "Email": "tom.hanks@gmail.com",
        "PhoneNumbers": [
            "000-000-0000"
        ],
        "Password": "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
        "AddressBook": [
            {
                "FirstName": "Tom",
                "LastName": "Hanks",
                "AddressLineOne": "101 Broad St",
                "AddressLineTwo": "",
                "City": "Santa Barbara",
                "State": "California",
                "Country": "United States",
                "Pincode": "00000",
                "Default": true
            }
        ],
        "PaymentOptions": [
            {
                "Name": "TOM HANKS",
                "CardNumber": "0000000000000000",
                "CardExpiryMM": "01",
                "CardExpiryYY": "20",
                "SecurityCode": "000",
                "ZipCode": "00000",
                "Default": true,
                "SaveInformation": true
            }
        ],
        "WishList": [
            "83947DSDS",
            "84378DFDW"
        ],
        "SaveForLater": [
            "FSDF3434",
            "ERF4432D"
        ],
        "Updated": 1579233419983053555
    }
}
```

Use this endpoint to add a customer into the database. 


### HTTP Request URL

`POST https://api.krama.ai/customers/{API version}/`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| CID                   |  Customer identifier                                |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |


## Update a customer's information

> Sample HTTP request body:

```json
{
  "Active": true,
  "FirstName": "Tom",
  "LastName": "Hanks",
  "Email": "tom.hanks@gmail.com",
  "PhoneNumbers": [
    "000-000-0000"
  ],
  "Password": "password",
  "AddressBook": [
    {
      "FirstName": "Tom",
      "LastName": "Hanks",
      "AddressLineOne": "101 Broad St",
      "AddressLineTwo": "",
      "City": "Santa Barbara",
      "State": "California",
      "Country": "United States",
      "Pincode": "00000",
      "Default": true
    }
  ],
  "PaymentOptions": [
    {
      "Name": "TOM HANKS",
      "CardNumber": "0000000000000000",
      "CardExpiryMM": "01",
      "CardExpiryYY": "20",
      "SecurityCode": "000",
      "ZipCode": "00000",
      "Default": true,
      "SaveInformation": true
    }
  ],
  "WishList": [
    "83947DSDS"
  ],
  "SaveForLater": [
    "FSDF3434",
    "ERF4432D"
  ]
}
```

> Sample valid API response:

```json
{
    "Code": 202,
    "Success": true,
    "Message": "Customer Updated ...",
    "Time": 1579233562710123000,
    "Response": {
        "CustomerID": "127452a3-00b6-4a04-a03b-e1db919645cc",
        "Active": true,
        "FirstName": "Tom",
        "LastName": "Hanks",
        "Email": "tom.hanks@gmail.com",
        "PhoneNumbers": [
            "000-000-0000"
        ],
        "Password": "password",
        "AddressBook": [
            {
                "FirstName": "Tom",
                "LastName": "Hanks",
                "AddressLineOne": "101 Broad St",
                "AddressLineTwo": "",
                "City": "Santa Barbara",
                "State": "California",
                "Country": "United States",
                "Pincode": "00000",
                "Default": true
            }
        ],
        "PaymentOptions": [
            {
                "Name": "TOM HANKS",
                "CardNumber": "0000000000000000",
                "CardExpiryMM": "01",
                "CardExpiryYY": "20",
                "SecurityCode": "000",
                "ZipCode": "00000",
                "Default": true,
                "SaveInformation": true
            }
        ],
        "WishList": [
            "83947DSDS"
        ],
        "SaveForLater": [
            "FSDF3434",
            "ERF4432D"
        ],
        "Updated": 1579233562701600551
    }
}
```

> Sample invalid API response:

```json
{
    "Code": 304,
    "Success": false,
    "Message": "Customer Not Found ...",
    "Time": 1579147170529434000,
    "Response": null
}
```


Use this information to update a customer's information. In the example shown in this section there is a request to remove one product sku from the customer's wishlist.


### HTTP Request URL

`PUT https://api.krama.ai/customers/{API version}/customers/{CID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request URL Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| CID                   |  Customer identifier                                |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |





## Delete a customer

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Customer Deleted ...",
    "Time": 1579146617669614321,
    "Response": null
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Customer Not Found ...",
    "Time": 1579146368716061070,
    "Response": null
}
```


Use this endpoint to delete the customer object from the database. 


### HTTP Request URL

`DELETE https://api.krama.ai/customers/{API version}/customers/{CID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |


### HTTP Request URL Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| CID                   |  Customer identifier                                |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |



# Shopping Cart API

## Get shopping cart

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Shopping Cart: ",
    "Time": 1579336082414378455,
    "Response": {
        "CartID": "6ea6b222-8634-43f3-bb27-067d58d515ef",
        "CustomerID": "1234",
        "ProductsCount": {
            "B07K3BHGL3": 5
        },
        "Products": {
            "B07K3BHGL3": {
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
                "Brand": "Microsoft xxx yyy",
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
                    "Number of USB 2 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "RAM": "16 GB",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac",
                    "testingb": true,
                    "testingf": 55.234,
                    "testingi": 345
                },
                "IsMain": true,
                "Currency": "CDN",
                "Updated": 1579300684745433221
            }
        },
        "Total": 13000,
        "Currency": "CDN",
        "Updated": 1579336048673188848
    }
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Cart id: 6ea6b222-8634-43f3-bb27-067d58d515efx not found ...",
    "Time": 1579336141745608597,
    "Response": null
}
```

Use this endpoint to get the shopping cart object using the cart ID. 


### HTTP Request URL

`GET https://api.krama.ai/shoppingcart/{API version}/cart/{CID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |


### HTTP Request URL Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| CID                   |  Shopping cart identifier                           |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |



## Add product in cart

> Sample HTTP request body:

```json
{
	"CustomerID" : "1234",
	"Product" : {
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
        "Brand": "Microsoft xxx yyy",
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
            "Number of USB 2 Ports": "1",
            "Operating System": "Windows 10 Home",
            "Processor Count": "16",
            "RAM": "16 GB",
            "Series": "Surface Laptop 2",
            "Shipping Weight": "2.2 kg",
            "Wireless Standard": "802.11ac",
            "testingb": true,
            "testingf": 55.234,
            "testingi": 345
        },
        "IsMain": true,
        "Currency": "CDN",
        "Updated": 1579300684745433221
    },    
    "Count" : 5
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Product added in the cart...",
    "Time": 1579335347860765751,
    "Response": {
        "CartID": "6ea6b222-8634-43f3-bb27-067d58d515ef",
        "CustomerID": "1234",
        "ProductsCount": {
            "B07K3BHGL3": 5
        },
        "Products": {
            "B07K3BHGL3": {
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
                "Brand": "Microsoft xxx yyy",
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
                    "Number of USB 2 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "RAM": "16 GB",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac",
                    "testingb": true,
                    "testingf": 55.234,
                    "testingi": 345
                },
                "IsMain": true,
                "Currency": "CDN",
                "Updated": 1579300684745433221
            }
        },
        "Total": 13000,
        "Currency": "CDN",
        "Updated": 1579335347860479220
    }
}
```


Use this endpoint to add a product in the shopping cart. If the mentioned product already exists in the cart, the quantity is incremented and the total is updated accordingly.
If a cart ID is not passed in the request the API creates a new unique cart ID - essentially a new shopping cart! So to create a shopping cart you need to make sure to NOT pass
a cart ID in your request. 


### HTTP Request URL

`POST https://api.krama.ai/shoppingcart/{API version}/cart/addproduct`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| CartID                | String, Shopping cart identifier                    |
| CustomerID            | String, Customer unique identifier                  |
| Product               | Product Object                                      |
| Count                 | Integer, the quantity of the added product          |

### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |



## Remove product from cart

> Sample valid HTTP request body:

```json
{
	"CartID" : "6ea6b222-8634-43f3-bb27-067d58d515ef",
	"CustomerID" : "1234",
	"SKU" : "B07K3BHGL3",
    "Count" : 4
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Product removed the cart...",
    "Time": 1579336317220235734,
    "Response": {
        "CartID": "6ea6b222-8634-43f3-bb27-067d58d515ef",
        "CustomerID": "1234",
        "ProductsCount": {
            "B07K3BHGL3": 6
        },
        "Products": {
            "B07K3BHGL3": {
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
                "Brand": "Microsoft xxx yyy",
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
                    "Number of USB 2 Ports": "1",
                    "Operating System": "Windows 10 Home",
                    "Processor Count": "16",
                    "RAM": "16 GB",
                    "Series": "Surface Laptop 2",
                    "Shipping Weight": "2.2 kg",
                    "Wireless Standard": "802.11ac",
                    "testingb": true,
                    "testingf": 55.234,
                    "testingi": 345
                },
                "IsMain": true,
                "Currency": "CDN",
                "Updated": 1579300684745433221
            }
        },
        "Total": 15600,
        "Currency": "CDN",
        "Updated": 1579336317220058378
    }
}
```

> Sample valid API response (when all the product quantity is removed):

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Product removed the cart...",
    "Time": 1579336403193133086,
    "Response": {
        "CartID": "6ea6b222-8634-43f3-bb27-067d58d515ef",
        "CustomerID": "1234",
        "ProductsCount": {},
        "Products": {},
        "Total": 0,
        "Currency": "CDN",
        "Updated": 1579336403192931756
    }
}
```


Use this endpoint to remove the product or adjust the product quantity in the shopping cart. 


### HTTP Request URL

`POST https://api.krama.ai/shoppingcart/{API version}/cart/removeproduct`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| CartID                |  Shopping cart identifier                           |
| CustomerID            |  Customer unique identifier                         |
| SKU                   |  Product SKU                                        |
| Count                 |  The quantity to be subtracted from cart            |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |



## Clear shopping cart

> Sample valid API response:

```json
{
    "Code": 202,
    "Success": true,
    "Message": "Cart with id: 6ea6b222-8634-43f3-bb27-067d58d515ef deleted ...",
    "Time": 1579335148388432310,
    "Response": null
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Cart id: 6ea6b222-8634-43f3-bb27-067d58d515ef not found ...",
    "Time": 1579335172915531779,
    "Response": null
}
```


Use this endpoint to reset/clear the shopping cart. 


### HTTP Request URL

`DELETE https://api.krama.ai/shoppingcart/{API version}/cart/clear/{CID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |


### HTTP Request URL Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| CID                   |  Customer identifier                                |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |




# Product Reviews API

## Get product reviews

> Sample valid HTTP request body:

```json
{
	"SortField" : "Time",
	"Order" : -1,
	"From"  : 0,
	"To" : 100 
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Customer Found ...",
    "Time": 1579464266245156552,
    "Response": [
        {
            "ReviewID": "2c588edf-c81e-449d-b9b6-ffefc00908a0",
            "Time": 1579464264414715770,
            "GroupID": "MSLAPS2",
            "CustomerID": "340950843976",
            "Stars": 1,
            "Description": "Test Review"
        },
        {
            "ReviewID": "5dfb83f7-5995-4be5-8c0c-d0dd5e99428c",
            "Time": 1579464261921297128,
            "GroupID": "MSLAPS2",
            "CustomerID": "340950843976",
            "Stars": 1,
            "Description": "Test Review"
        }
    ]
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Reviews Not found ...",
    "Time": 1579464296786186393,
    "Response": null
}
```


Use this endpoint to product reviews for a product group. It must be noted that the reviews are stored and retrieved in context with product groups and not individual products as it 
does not make any sense to store reviews by products which could be indeed versions of other products.  


### HTTP Request URL

`GET https://api.krama.ai/productreviews/{API version}/reviews/{PGID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |


### HTTP Request URL Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| PGID                  |  Customer identifier                                |


### HTTP Request Body Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| From                  | Starting index of reviews                           |
| Count                 | Number of reviews                                   |
| SortField             | Field on which sorting needs to be applied          |
| Order                 | Value -1 Decending ; Value 1 Ascending              |             



### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |




## Post a product review

> Sample valid HTTP request body:

```json
{
	"GroupID" : "MSLAPS2",
	"CustomerID" : "340950843976",
	"Stars" : 1,
	"Description" : "Test Review"
}
```

> Sample invalid API response:

```json
{
    "Code": 201,
    "Success": true,
    "Message": "Review Added and Cumulative review data updated in Product Group object ...",
    "Time": 1579464264416813978,
    "Response": {
        "ReviewID": "2c588edf-c81e-449d-b9b6-ffefc00908a0",
        "Time": 1579464264414715770,
        "GroupID": "MSLAPS2",
        "CustomerID": "340950843976",
        "Stars": 1,
        "Description": "Test Review"
    }
}
```


Use this endpoint to post a review for a product group. 


### HTTP Request URL

`POST https://api.krama.ai/productreviews/{API version}/reviews`

### HTTP Request Header


| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |



### HTTP Request Body Parameters

| Parameter             |               Description                                                 |
|-----------------------|---------------------------------------------------------------------------|
| GroupID               |  String, Product GroupID for which the review is being posted             |
| CustomerID            |  String, CustomerID of the customer who is posting this review            |
| Stars                 |  Float,  star rating on the review  (value between: 1-5)                  |
| Description           |  String, description in the review. Less than 10240 characters            |                  


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |





## Delete a product review

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Review deleted ...",
    "Time": 1579464792787915410,
    "Response": null
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Review Not Found ...",
    "Time": 1579464802608052274,
    "Response": null
}
```


Use this endpoint to delete a review for a product group using the unique ReviewID. 


### HTTP Request URL

`DELETE https://api.krama.ai/productreviews/{API version}/reviews/{RID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |


### HTTP Request URL Parameters

| Parameter             |               Description                           |
|-----------------------|-----------------------------------------------------|
| RID                   |  Review ID of the review to be deleted              |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |





## Delete all product reviews

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Review for product group deleted ...",
    "Time": 1579464873282343364,
    "Response": null
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Reviews for Product group mentioned in request, Not Found ...",
    "Time": 1579464888741724246,
    "Response": null
}
```


Use this endpoint to remove all the reviews for a product group using the product GroupID. 


### HTTP Request URL

`DELETE https://api.krama.ai/productreviews/{API version}/reviews/productgroup/{PGID}`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |


### HTTP Request URL Parameters

| Parameter             |               Description                                                                       |
|-----------------------|-------------------------------------------------------------------------------------------------|
| PGID                  |  Product GroupID of the product group for which the reviews are expected to be deleted          |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |



# Category API


## Get category products

> Sample valid HTTP request body:

```json
{
	"Path" : "Electronics>Computers>Laptops"
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Products in category path ...",
    "Time": 1579752116485856444,
    "Response": {
        "Electronics>Computers>Laptops": [
            "B07K3BHGL3"
        ]
    }
}
```

> Sample invalid API response:

```json
{
    "Code": 400,
    "Success": false,
    "Message": "Category path does not exit ...",
    "Time": 1579753925588641431,
    "Response": null
}
```


Use this endpoint to get product SKUs in the category.


### HTTP Request URL

`GET https://api.krama.ai/categories/{API version}/products`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

| Parameter             |               Description                                                                       |
|-----------------------|-------------------------------------------------------------------------------------------------|
| Path                  |  Valid category path - Example: "Electronics>Computers>Laptop"                                  |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |




## Get root categories


> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Root categories ...",
    "Time": 1579754001879780149,
    "Response": [
        "Electronics",
        "Clothes"
    ]
}
```


Use this endpoint to get all the root categories. 


### HTTP Request URL

`GET https://api.krama.ai/categories/{API version}/root`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |





## Get sub categories

> Sample valid HTTP request body:

```json
{
	"Category" : "Electronics"
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Immediate Sub categories ...",
    "Time": 1579754123461121424,
    "Response": [
        "Computers"
    ]
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Category Laptops does not have a sub category ...",
    "Time": 1579754183920349803,
    "Response": null
}
```


Use this endpoint to get immediate sub categories of a category. 


### HTTP Request URL

`GET https://api.krama.ai/categories/{API version}/sub`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

| Parameter             |               Description                                                                       |
|-----------------------|-------------------------------------------------------------------------------------------------|
| Path                  |  Valid category path - Example: "Electronics>Computers>Laptop"                                  |
| Category              |  Valid category name - Example: "Computers"                                                     |


### HTTP Response

|  Key              |    Description                                                                |
|-------------------|-------------------------------------------------------------------------------|
| Code              | Response code for the request                                                 |
| Success           | Flag that tells if the request was successful                                 |
| Message           | Message for additional information                                            |
| Time              | Unix timestamp of the response                                                |
| Response          | Response object containing response information                               |






## Get parent category

> Sample valid HTTP request body:

```json
{
	"Category" : "Computers"
}
```

> Sample valid API response:

```json
{
    "Code": 200,
    "Success": true,
    "Message": "Category parent ...",
    "Time": 1579754248936370320,
    "Response": "Electronics"
}
```

> Sample invalid API response:

```json
{
    "Code": 404,
    "Success": false,
    "Message": "Category Electronics does not have a parent ...",
    "Time": 1579754265682142471,
    "Response": null
}
```


Use this endpoint to get the name of the parent category of a given category. 


### HTTP Request URL

`GET https://api.krama.ai/categories/{API version}/parent`

### HTTP Request Header

| Key               |                Value                         |
|-------------------|----------------------------------------------|
|Authorization      | Bearer {YOUR_ACCESS_TOKEN}                   |
|Content-Type       | application/json                             |


### HTTP Request Body Parameters

| Parameter             |               Description                                                                       |
|-----------------------|-------------------------------------------------------------------------------------------------|
| Path                  |  Valid category path - Example: "Electronics>Computers>Laptop"                                  |
| Category              |  Valid category name - Example: "Computers"                                                     |


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

| Response Code                      | Meaning                                                                             |
|------------------------------------|-------------------------------------------------------------------------------------|
| 200  **StatusOK**                  | The API call was successful                                                         |
| 201  **StatusCreated**             | The resource was successfully created                                               |
| 202  **StatusAccepted**            | The API request was accepted                                                        |
| 304  **StatusNotModified**         | The request to modify the resource failed for some reason                           |
| 400  **StatusBadRequest**          | The API call was malformed                                                          |
| 401  **StatusUnauthorized**        | The API request is not authorized. Please check your access credentials             |
| 404  **StatusNotFound**            | The requested/referenced resource was not found                                     |
| 409  **StatusConflict**            | Code that is returned when, for example, product with same SKU already exists       |
| 429  **StatusTooManyRequests**     | Your servers are sending more API requests per minute than allowed                  |  
| 500  **StatusInternalServerError** | Something went wrong on our server                                                  |
| 503  **StatusServiceUnavailable**  | The API servers are down for some internal reasons                                  |


# Non-functional Features

## API Rate Limit

> Sample valid API response:

```json
{
    "Code": 429,
    "Success": false,
    "Message": "Rate capacity of 10 requests per minute is reached for this customer ...",
    "Time": 1585673070784773975,
    "Response": null
}
```

This SaaS platform enforces API rate limiting for each customer. API rate limiting is helpful in ensuring that no single enterprise customer unfairly consumes more resources than any other enterprise customer. Rate limiting also adds a level of security to the platform by preventing potential DoS attacks. This scheme also allows the rationing of the API endpoints in the time of need. The API rate limiting unit of measurement is number of requests per minute. Once the rate limit is reached for the customer, all API requests from the customer are responded with an HTTP response code of 429. Please see the sample response object provided on the side. This throttling is implemented using Redis. To know more about the implementation approach please check <a href="https://redislabs.com/redis-best-practices/basic-rate-limiting/" target="_blank">this Redislabs blog</a>