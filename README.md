My Rest API
===

---
Project source can be downloaded from:
https://github.com/PasinduDiss/MyRestserviceWorkspace
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
      │   └──tests                        # Live tests for each lambda function (you will need to edit)
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
      │    ├── livetest.sh
      │    └── test.sh
      │ 
      ├── serverless.yml        # serverless.yml a yaml file used to specify  
      ├── MyRestservice.json   
      ├── Gopkg.lock
      ├── README.md                     
      └── Gopkg.toml            # MyRestservice.json contains a postman collection to test REST API
                                          
```
**Prerequisites:**

- Serverless framework:
```
npm install -g serverless
```
- Golang install

- Install dep using homebrew:
```
brew install dep && brew upgrade dep
```

Please ensure you add GOROOT and GOPATH variables to the your .bash_profile
Instructions can be found in this [video](https://www.youtube.com/watch?v=FTDOW8UbKjQ&t=252s). 

**Setup**

Go does not allow for projects be in multiple GOPATHS, thus it is recommended to use one go workspace for development. This repository contains a go workspace with the serverless REST API project contained in the source (src) folder of this go workspace. 

This was done deliberately to ensure you are able to run the project in any directory to which this repository is cloned. Please perform the following steps to ensure the GOPATH environment variable points to this workspace.

Ensure you have followed the Prerequisites and set up GOROOT and GOPATH variables in your bash_profile.

```
cd MyRestserviceWorkspace
export GOPATH=`pwd`
```

This will ensure the GOPATH variable points to MyRestserviceWorkspace, you can now navigate to MyRestserviceWorkspace/src/MyRestservice project and deploy!

**Directory Structure and Information** 

The repository contains a golang workspace, a golang workspace contains a package(pkg) directory which contains package objects, a binaries(bin) directory containing executables, and a source(src) directory containing the GO source files. The project MyRestservice is placed in the src folder of this Workspace. The following is and explaination of the contents of MyRestservice project. 

Src (source directory) contains the handlers directory and main directory. The handlers directory contains the handlers.go file, which is used to create the lambda functions used in the various REST API calls. The main directory contains four subdirectories: create, get, delete and list. 

These subdirectories are named according to their corresponding lambda functions. Each of these directories contain main.go files of the main package which is necessary in order to create the separate binary files which correspond to each function. The main.go files need to be separated in this manner due to the Go language not allowing for more than one main function in a given directory. 

The subdirectories under the main directory also contain unit tests for each lambda function. The bin directory contains binaries which are automatically generated when running the build or deploy scripts. These binaries are used by the serverless framework to deploy the lambda functions specified. 

The scripts directory contains the scripts build.sh, deploy.sh and test.sh. The build.sh script should be used if you would like to compile the lambda functions but not deploy as it compiles all main.go files in the subdirectories of the main directory and creates the binary files with the corresponding directory name in the bin directory. 

If you would like to build and deploy, the deploy.sh script can be used, it will run the build script and then proceed to deploy the REST API

There are two different types of tests, live tests which use HTTP requests to interact with the deployed REST API and unit tests for each main.go function. To run the live tests you will need to edit the code, the url endpoints for each HTTP request is hardcoded, thus you will need to replace them with your own endpoints generated. 

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



Windows :

Deployment have not been tested on windows operating system.


How to invoke REST calls
------
Using curl :

**POST Request**

This function will insert items into the dynamodb table

```
curl -H "Content-Type: application/json" -X POST -d '{
"id": "ID","deviceModel": "DEVICE_MODEL", "name": "NAME","note": "NOTE","serial": "SERIAL"}' https://Your-end-point-url
```

**GET Request**

This function will list all items from the dynamodb table

```
curl -H "Content-Type: application/json" -X GET https://Your-end-point-url
```

**GET Request (get single item)**

This function will get a single item from the dynamodb table according to the provided path variable

```
curl -H "Content-Type: application/json" -X GET https://Your-end-point-url/{id}
```
**DELETE Request**

This function will delet a single item from the dynamodb table according to the provided path variable

```
curl -H "Content-Type: application/json" -X DELETE https://Your-end-point-url/{id}
```
Testing
-------
The script test.sh can be used to run the unit tests, testing can be invoked by:

```
scripts/test.sh
```

For live tests you can use the following command, ensure the ENDPOINT variable in live test go files are replaced with your own endpoints generated.

```
scripts/livetest.sh
```

To live test using Postman, please edit the provided Postman collection with the endpoints generated for each POST, GET(list), GET(item), DELETE  request. 



To run unit tests follow the commands below and run the test.sh script

```
cd MyRestservice
scripts/test.sh
```

Postman was also used to test the API, the postman collection and tests are included can be found [here](MyRestservice.postman_collection.json)

Assumptions
-------

- deviceModel value stored in dynamod db without the recource path "/deviceModel/" thus in the case of {id : /devices/id1, deviceModel: /deviceModel/id1} both values would be stored as "id1"

- It is assumed the data sent using POST Http method is of the format {"id":"/devices/id1","deviceModel":"/deviceModel/id1","name":"Sensor","note":"Testing a sensor.","serial":"A020000101"}

- It is assumed the data recived using GET Http method with path parameters (/devices/{id})is of the format {"id":"/devices/id1","deviceModel":"/deviceModel/id1","name":"Sensor","note":"Testing a sensor.","serial":"A020000101"}

- Removing the serverless stack (sls remove) will remove the dynamodb table resource, if this is not the desired outcome please change the dynamodb Deletion policy from "Delete" to "Retain".