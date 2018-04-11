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


How to invoke REST calls
------
Using curl :

**POST Request**


curl -H "Content-Type: application/json" -X POST -d '{
"id": "ID","deviceModel": "DEVICE_MODEL", "name": "NAME","note": "NOTE","serial": "SERIAL"}' https://9mm0tq0pc5.execute-api.us-east-1.amazonaws.com/dev/devices

**GET Request**
This function will list the
curl -H "Content-Type: application/json" -X GET https://9mm0tq0pc5.execute-api.us-east-1.amazonaws.com/dev/devices


**GET Request With Parameters**

curl -H "Content-Type: application/json" -X GET https://9mm0tq0pc5.execute-api.us-east-1.amazonaws.com/dev/devices/{id}

**DELETE Request**

curl -H "Content-Type: application/json" -X DELETE https://9mm0tq0pc5.execute-api.us-east-1.amazonaws.com/dev/devices/{id}

Testing
_______
Postman was used to test the API, the postman collection and tests are included
[collection](MyRestservice.json)
