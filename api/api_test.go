package api

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerCheckConnections(t *testing.T) {
	cases := []struct {
		description          string
		testURL              string
		method               string
		expectedResponseCode int
	}{
		{
			description:          "Test HTTP Status OK",
			testURL:              "/v0/connections/fake",
			method:               "PUT",
			expectedResponseCode: http.StatusCreated,
		},
		{
			description:          "Test HTTP URL not found",
			testURL:              "/v0/connections",
			method:               "PUT",
			expectedResponseCode: http.StatusNotFound,
		},
		{
			description:          "Test HTTP Wrong Method",
			testURL:              "/v0/connections/fake",
			method:               "PATCH",
			expectedResponseCode: http.StatusNotFound,
		},
	}

	for _, c := range cases {

		r := mux.NewRouter()
		var jsonStr = []byte(`{
	"DBconnection":
	{
		"Proto":"http",
		"Ipaddress":"localhost",
		"Port":"27017",
		"Name":"Test",
		"Username":"",
		"Password":"",
		"Collection":"testing"
	}
}`)
		req, _ := http.NewRequest(c.method, c.testURL, bytes.NewBuffer(jsonStr))
		res := httptest.NewRecorder()

		r.HandleFunc("/v0/connections/{dbType}", func(w http.ResponseWriter, r *http.Request) {
			HandlerCheckConnections(res, req)
		}).Methods("PUT")

		r.ServeHTTP(res, req)

		if res.Code != c.expectedResponseCode {
			fmt.Println(c.description)
			t.Errorf("Error, expected: %d Received: %d", c.expectedResponseCode, res.Code)
		}
		fmt.Printf("%s:**** PASS ****code result: %d and Code Expected: %d \n", c.description, res.Code, c.expectedResponseCode)

	}
}
