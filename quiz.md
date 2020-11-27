# Quiz

create an api for a merchant seller.

## Requirement
- language: golang
- database: any

## Point (100 point)
- separate packages via structure (25 point)
- separate config between SIT and PROD (10 point)
- show product `amount` and `stock` as humanize (5 point)
- cicd (20 point)
- unit test (20 point)
- interface usages (20 point)

## Merchant Fields
- Name
- Bank Account 
- Username
- Password

## Product Fields
- Name
- Amount
- Stocks

## APIs
| Name                 | Method | Endpoint                          |
|----------------------|--------|-----------------------------------|
| Register Merchant    | POST   | /merchant/register                |
| Merchant Information | GET    | /merchant/information/:id         |
| Update Merchant      | POST   | /merchant/update                  |
| List All Products    | GET    | /merchant/:id/products            |
| Add Product          | POST   | /merchant/:id/product             |

## Batch
- sell reports as json file

### Register Merchant
- auto generate username and password
- each api must check username/password except register and buy product
- cannot register using the same bank account

### Merchant Information
- response merchant information

### Update Merchant
- can only update name

### List All Products
- list all merchant products with name and amount

### Add Product
- add product for each merchant 
- amount can be present in 2 precision, ex. 100.01, 250.35
- maximum products is 5

### Sell Reports
- sell report range only by date
- provide list of selling products and amount accumulate with precision point 2 digit

```json
{
	"date": "2018-11-01",
	"products": [
		{"name": "ABC", "selling_volume": 10},
		{"name": "DEF", "selling_volume": 5}
	],
	"accumulate": 100.25 
}
```
