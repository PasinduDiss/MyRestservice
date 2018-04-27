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
The following is a simple rest api created using the tech stack below
- Serverless framework  
- Golang
- AWS Lambda functions
- AWS DynamoDb no sql database
- AWS API gateway

Deploying Serverless Rest API
-----------------------------
**Project Contents**
```
MyRestservice
      ├── src
      │   ├──handlers               
      │   │         └── handlers.go         # contains lambda the functions   
      │   │             
      │   │                             
      │   ├── main                         # main package to create binaries of lambda functions
      │   │      ├── create
      │   │      │     ├── create_test.go
      │   │      │     └── main.go 
      │   │      │          
      │   │      ├── delete
      │   │      │     ├── delete_test.go
      │   │      │     └── main.go      
      │   │      │
      │   │	 ├── get
      │   │      │    ├── get_test.go
      │   │      │    └── main.go 
      │   │      │
      │   │      └── list
      │   │            ├── get_test.go      
      │   │            └── main.go  
      │   │
      │   └──tests                        # Live unit tests for each lambda function
      │        ├── listlive_test.go
      │        │
      │        ├── deletelive_test.go
      │        │
      │        ├── getlive_test.go
      │        │ 
      │        └── createlive_test.go
      │   
      ├── bin                    # bin file contains binary files of lambda functions
      │   └── handlers               
      │       ├── create
      │       ├── get
      │       ├── list	     
      │       └── delete
      │ 
      ├── scripts                # Scripts used to build, deploy and test lambda functions
      │    ├── build.sh
      │    ├── deploy.sh
      │    └── test.sh
      │ 
      ├── serverless.yml        # serverless.yml a yaml file used to specify  
      ├── MyRestservice.json   
      ├── Gopkg.lock
      ├── README.md                     
      └── Gopkg.toml            # MyRestservice.json contains a postman collection to test REST API
                                          
```
**Prerequisites:**

Ensure the repository is cloned in to a known GOPATH/src 

- Serverless framework:
```
npm install -g serverless
```
- Golang install

- dep brew:
```
install dep && brew upgrade dep
```
Please ensure you add GOROOT and GOPATH variables to the your .bash_profile
Instructions can be found in this [video](https://www.youtube.com/watch?v=FTDOW8UbKjQ&t=252s) 

**Setup**
Go does not allow for projects in multiple GOPATHS, thus it is recommended to use one go workspace for development. This repository contains a go workspace with the serverless REST API go project contained in the source (src) folder of the go workspace. This was done deliberately to ensure you are able to run the project in any directory to which this repository is cloned. Please follow the following steps to ensure the GOPATH environment variable points to this workspace.

Ensure you have followed the Prerequisites and set up GOROOT and GOPATH variables in your bash_profile
```
cd MyRestserviceWorkspace
export GOPATH=`pwd`
```
This will ensure the GOPATH variable points to MyRestserviceWorkspace, you can now navigate to src/MyRestservice project and deploy!

**Directory Structure and Information** 

Src (source directory) contains the handlers directory and main directory. The handlers directory contains the handlers.go file, which is used to create the lambda functions used in the various REST API calls. The main directory contains four subdirectories: create, get, delete and list. 

These subdirectories are named according to their corresponding lambda functions. Each of these directories contain main.go files of the main package which is necessary in order to create the separate binary files which correspond to each function. The main.go files need to be separated in this manner due to the Go language not allowing for more than one main function in a given directory. 

The subdirectories under the main directory also contain unit tests for each lambda function. The bin directory contains binaries which are automatically generated when running the build or deploy scripts. These binaries are used by the serverless framework to deploy the lambda functions specified. 

The scripts directory contains the scripts build.sh, deploy.sh and test.sh. The build.sh script should be used if you would like to compile the lambda functions but not deploy as it compiles all main.go files in the subdirectories of the main directory and creates the binary files with the corresponding directory name in the bin directory. 

If you would like to build and deploy, the deploy.sh script can be used, it will run the build script and then proceed to deploy the REST API

There are two different types of tests, live tests which use HTTP requests to interact with the deployed REST API and unit tests for each main.go function.

**Deploy:**

Mac/linux :

Through terminal navigate to root directory of project:

```
scripts/deploy.sh
```
This script will build and place all executable files in the bin folder
It also deploys the functions using "serverless deploy --aws-profile serverless"
command, you may need to edit this command if you do not have your aws-profile
configured for serverless. If aws-profile is configured as default there is no
need for the --aws-profile flag. The command "severless deploy" should be used.


The script build.sh can be used to build the handler functions without deploying
the service, build can be invoked by:

```
scripts/build.sh
```

The script test.sh can be used to run the unit tests as well as the live tests, testing can be invoked by:

```
scripts/test.sh
```

Windows :

Deployment have not been tested on windows operating system.


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
To run unit tests and the live tests follow the commands below and run the test.sh script

```
cd MyRestservice
/scripts/test.sh
```

Postman was also used to test the API, the postman collection and tests are included can be found [here](MyRestservice.postman_collection.json)

Assumptions
-------

- deviceModel value stored in dynamod db without the recource path "/deviceModel/" thus in the case of {id : /devices/id1, deviceModel: /deviceModel/id1} both values would be stored as "id1"

- It is assumed the data sent using POST Http method is of the format {"id":"/devices/id1","deviceModel":"/deviceModel/id1","name":"Sensor","note":"Testing a sensor.","serial":"A020000101"}

- It is assumed the data recived using GET Http method with path parameters (/devices/{id})is of the format {"id":"/devices/id1","deviceModel":"/deviceModel/id1","name":"Sensor","note":"Testing a sensor.","serial":"A020000101"}

- It is assumed that a dynamodb resource named Devices exists in region us-east-1 in the aws account used for deployment. 
