package main_test

import(
  "testing"
  main "github.com/aws-samples/lambda-go-samples"
  "github.com/aws/lambda-go-samples/events"
  "github.com/stretchr/testify/assert"
)

func UnitTest(t *testing.T){

  struct Test{
    request events.APIGateWayProxyRequest
    expectedbody  string
    expectedstatus int
    err     error
  }

  struct UnitTestList{
    UnitTests []Test
  }

  unittestlist_create := &UnitTestList{
    {
      //HTTP event
      request events.APIGatewayProxyRequest{Body: {"id":"/devices/id3","deviceModel":"/deviceModel/id3","name":"Sensor","note":"Testing a sensor.","serial":"A020000103"}}
      expected 		"Created"
      error  nil
    },
    {
      request events.APIGatewayProxyRequest{URL: "https://9mm0tq0pc5.execute-api.us-east-1.amazonaws.com/dev/devices/id1"}
      expectedbody 		"Bad Request"
      expectedstatus    200
      error  nil
    }
  }

  unittestlist_get := &UnitTestList{
    {
      //HTTP event
      request events.APIGatewayProxyRequest{Body: {}
      expected 		"Created"
      error  nil
    },
    {
      request events.APIGatewayProxyRequest{Body: {}}
      expectedbody 		"Bad Request"
      expectedstatus 400
      error  nil
    }
  }





}
