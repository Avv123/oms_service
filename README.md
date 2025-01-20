
### Setup For The Project:


``` golang
go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
go get go.mongodb.org/mongo-driver/mongo

go run main.go

```
### Create your own .env file

### Endpoint:user/register : Creates a user


### input: 
curl --location 'http://localhost:8080/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":"1",
    "username":"aryaman",
    "email":"arya2003@gmail.com",
    "password":"arya2003"

}'
### output:

{"User created":{"id":"1","username":"aryaman","email":"arya2003@gmail.com","password":"$2a$10$wFAylMXPVLN227AwHxUAiueXs4NfXqOh0fxWLI/Au7zysbx4ihfoK","orders":null}}
![image](https://github.com/user-attachments/assets/4552e91c-e66e-4f68-9e35-d9b0572633dc)

### Endpoint:user/login : Logins a user


## input: 
curl --location 'http://localhost:8080/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    
    
    "email":"arya2003@gmail.com",
    "password":"arya2003"

}'

### output:

{"message":"Login successful","user":{"id":"1","username":"aryaman","email":"arya2003@gmail.com","password":"$2a$10$wFAylMXPVLN227AwHxUAiueXs4NfXqOh0fxWLI/Au7zysbx4ihfoK","orders":null}}
![image](https://github.com/user-attachments/assets/4124142b-0726-4203-b63a-8a75bb24753e)

### Endpoint: user/user/users : Get all users

### Input: 
curl --location --request GET 'http://localhost:8080/user/user/users' \user/user/users' \
--header 'Content-Type: application/json' \
--data '
'

### output: 

[{"id":"1","username":"aryaman","email":"arya2003@gmail.com","password":"$2a$10$wFAylMXPVLN227AwHxUAiueXs4NfXqOh0fxWLI/Au7zysbx4ihfoK","orders":null}]

![image](https://github.com/user-attachments/assets/eac4d1e4-ba20-439e-aa6a-68d29bd9a0e5)

### Endpoint:user/user/1

![image](https://github.com/user-attachments/assets/a1e30c73-f2b6-45cf-b32b-d3b579276aa8)

### Endpoint: user/products/create-product: Create a product

### Input:
curl --location 'http://localhost:8080/user/products/create-product' \ects/oms$ curl --location 'http://localhost:8080/user/products/create-product' \
--header 'Content-Type: application/json' \
--data '{
    "id":"a",
    "seller_id":"*12",
    "name":"iphone",
    "description":"a good phone",
    "price":356,
    "stock":20

}'

### output: 

{"product":{"id":"a","seller_id":"*12","name":"iphone","description":"a good phone","price":356,"stock":20}}

![image](https://github.com/user-attachments/assets/145fd7e9-bdb6-4a0c-adc7-acfe941316fc)

### Endpoint: user/products/getallproducts : Get all products

### Input:

curl --location --request GET 'http://localhost:8080/user/products/getallproducts' \
--header 'Content-Type: application/json' \
--data '{
    "id":"d",
    "seller_id":"*56356",
    "name":"iphone 14",
    "SKU":"128gb",
    "description":"a good phone",
    "price":326,
    "stock":20

}'

### Output:

[{"id":"a","seller_id":"*12","name":"iphone","description":"a good phone","price":356,"stock":20},{"id":"b","seller_id":"*562","name":"iphone 13","description":"a good phone","price":326,"stock":20},{"id":"c","seller_id":"*56256","name":"iphone 13","description":"a good phone","price":326,"stock":20},{"id":"d","seller_id":"*56356","name":"iphone 14","description":"a good phone","price":326,"stock":20}]a


![image](https://github.com/user-attachments/assets/7b54b94f-c247-4e71-b04f-93a33361fae1)

### Endpoint: user/products/get-product/:id : Get product 

### Input:

 curl --location --request GET 'http://localhost:8080/user/products/get-product/e' \//localhost:8080/user/products/get-product/e' \
--header 'Content-Type: application/json' \
--data '{
    "id":"e",
    "seller_id":"*59856",
    "name":"iphone 14",
    "SKU":"128gb",
    "description":"a good phone",
    "price":326,
    "stock":20

}'

### output:

![image](https://github.com/user-attachments/assets/330b977d-aac2-4e06-9621-51e001d9382c)

### Routes working:

![image](https://github.com/user-attachments/assets/48ef85b7-db66-43be-84e6-e8935d61dec7)




