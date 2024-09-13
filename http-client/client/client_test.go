package client_test

import (
	"http-client/client"
	"http-client/types"
	"testing"
)

func Test_Can_Connect_To_HTTP_Server_And_Get_JSON_Response(t *testing.T) {
	expected := &types.Response{
		Status:  "success",
		Message: "OK",
	}

	res, err := client.Get("http://127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}

	if expected != res {
		t.Fatal(err)
	}
}
