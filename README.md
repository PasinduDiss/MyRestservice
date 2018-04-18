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
**Contents**
```
MyRestservice 
      ├── src
      │   ├──handlers               
      │   │         └── handlers.go        #contains lambda the functions   
      │   │             
      │   │                             
      │   └── main                        # main package to create binaries of lambda
      │          ├── create                 functions  
      │          │     ├── create_test.go
      │          │     └── main.go 
      │          │          
      │          ├── delete
      │          │     ├── delete_test.go
      │          │     └── main.go      
      │          │
      │  	 ├── get          
      │          │    ├── get_test.go
      │          │    └── main.go 
      │          │
      │          └── list
      │               ├── get_test.go      
      │               └── main.go  
      │ 
      ├── bin                            # bin file contains binary files of lambda 
      │   └── handlers                     functions
      │       ├── create
      │       ├── get
      │       ├── list	     
      │       └── delete
      │ 
      ├── scripts                        # Scripts used to build, deploy and test
      │    ├── build.sh                    lambda functions
      │    ├── deploy.sh
      │    └── test.sh
      │ 
      ├── serverless.yml                # serverless.yml a yaml file used to specify  
      ├── MyRestservice.json              resources to be used by api as well as 
      ├── Gopkg.lock                      stating stating the functions to be created
      ├── README.md                      
      └── Gopkg.toml                    # MyRestservice.json contains a postman
                                          collection to test REST API 
                                          
```

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

Mac/linux :

Through terminal navigate to MyRestservice directory enter command:

```
scripts/deploy.sh
```
This script will build and place all executable files in the bin folder
It also deploys the functions using "serverless deploy --aws-profile serverless"
command, you may need to edit this command if you do not have your aws-profile
configured for serverless. If aws-profile is configured as default there is no
need for the --aws-profile flag. The command "severless deploy" should be used.


The script remove.sh can be used to remove the deployed service, invoked by the
command:

```
scripts/remove.sh
```

The script build.sh can be used to build the handler functions without deploying
the service, build can be invoked by:

```
scripts/build.sh
```

How to invoke REST calls
------
Using curl :

**POST Request**

This function will insert items into the dynamodb table
```
curl -H "Content-Type: application/json" -X POST -d '{
"id": "ID","deviceModel": "DEVICE_MODEL", "name": "NAME","note": "NOTE","serial": "SERIAL"}' https://xox3imgc04.execute-api.us-east-1.amazonaws.com/dev/devices
```
**GET Request**

This function will list all items from the dynamodb table
```
curl -H "Content-Type: application/json" -X GET https://xox3imgc04.execute-api.us-east-1.amazonaws.com/dev/devices
```
**GET Request (get single item)**

This function will get a single item from the dynamodb table according to the provided path variable
```
curl -H "Content-Type: application/json" -X GET https://xox3imgc04.execute-api.us-east-1.amazonaws.com/dev/devices/{id}
```
**DELETE Request**

This function will delet a single item from the dynamodb table according to the provided path variable
```
curl -H "Content-Type: application/json" -X DELETE https://xox3imgc04.execute-api.us-east-1.amazonaws.com/dev/devices/{id}
```
Testing
-------

Postman was used to test the API, the postman collection and tests are included can be found [here](MyRestservice.postman_collection.json)

To run unit tests follow the commands below and run the test.sh script
```
cd MyRestservice
/scripts/test.sh
```