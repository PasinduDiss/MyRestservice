My Rest API
===

---
Project source can be downloaded from:
https://github.com/PasinduDiss/MyRestservice.git
----

Author
------

Pasindu Dissanayake

Abstract
--------
The following is a simple rest api created using the following tech stack
- Serverless framework  
- Golang
- AWS Lambda functions
- AWS DynamoDb no sql database
- AWS API gateway

Deploying Serverless Rest API
-----------------------------
**Prerequisites:**

- Serverless framework:
```
npm install -g serverless
```
- Golang
- dep brew:
```
install dep && brew upgrade dep
```

**Deploy:**

Mac :

Through terminal navigate to MyRestservice directory enter command:

```
scripts/deploy.sh
```


How to invoke REST calls
------
Using curl :

**POST Request**

This function will insert items into the dynamodb table
```
curl -H "Content-Type: application/json" -X POST -d '{
"id": "ID","deviceModel": "DEVICE_MODEL", "name": "NAME","note": "NOTE","serial": "SERIAL"}' https://9mm0tq0pc5.execute-api.us-east-1.amazonaws.com/dev/devices
```
**GET Request**

This function will list all items from the dynamodb table
```
curl -H "Content-Type: application/json" -X GET https://9mm0tq0pc5.execute-api.us-east-1.amazonaws.com/dev/devices
```
**GET Request With Parameters**
This function will get a single item from the dynamodb table according to the provided path variable
```
curl -H "Content-Type: application/json" -X GET https://9mm0tq0pc5.execute-api.us-east-1.amazonaws.com/dev/devices/{id}
```
**DELETE Request**
This function will delet a single item from the dynamodb table according to the provided path variable 
```
curl -H "Content-Type: application/json" -X DELETE https://9mm0tq0pc5.execute-api.us-east-1.amazonaws.com/dev/devices/{id}
```
Testing
_______
Postman was used to test the API, the postman collection and tests are included can be found [here](MyRestservice.json)
