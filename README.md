My Rest API
===

---
Project source can be downloaded from
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
"id": "ID","deviceModel": "DEVICE_MODEL", "name": "NAME","note": "NOTE","serial": "SERIAL"}' https://q5b7gvb7r8.execute-api.us-east-1.amazonaws.com/dev/api/devices

**GET Request**
This function will list the
curl -H "Content-Type: application/json" -X GET https://q5b7gvb7r8.execute-api.us-east-1.amazonaws.com/dev/api/devices


**GET Request With Parameters**

curl -H "Content-Type: application/json" -X GET https://q5b7gvb7r8.execute-api.us-east-1.amazonaws.com/dev/api/devices?id=ID

**DELETE Request**

curl -H "Content-Type: application/json" -X DELETE https://q5b7gvb7r8.execute-api.us-east-1.amazonaws.com/dev/api/devices?id=ID

Resources
_________
