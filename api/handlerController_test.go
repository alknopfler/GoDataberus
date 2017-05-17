package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alknopfler/tdd_api_mongodb/database"
	"errors"
	"github.com/alknopfler/tdd_api_mongodb/data_model"
	"github.com/gorilla/mux"

	"fmt"
)

type fakeDriver struct {}

func TestMyRouterHandler(t *testing.T){
	cases := []struct {
		description		string
		testURL			string
		method  		string
		expectedResponseCode 	int
	}{
		{
			description:  	      "Test HTTP Status OK",
			testURL:	      "/v0/countries/spain/services/svc1",
			method:		      "GET",
			expectedResponseCode: http.StatusOK,
		},
		{
			description:	      "Test HTTP URL not found",
			testURL:	      "/v0/countries/spain/services",
			method: 	      "GET",
			expectedResponseCode: http.StatusNotFound,
		},
		{
			description:	      "Test HTTP Wrong Method",
			testURL:              "/v0/countries/spain/services/svc1",
			method: 	      "PATCH",
			expectedResponseCode: http.StatusNotFound,
		},
	}

	for _, c := range cases {

		r := mux.NewRouter()
		req, _ := http.NewRequest(c.method, c.testURL, nil)
		res := httptest.NewRecorder()


		f := fakeDriver{}
		db := database.ConnectionDB{"localhost","test"}

		r.HandleFunc("/v0/countries/{country}/services/{serviceid}", func(w http.ResponseWriter, r *http.Request) {
			HandlerServices(res,req,&f,db)
		}).Methods("GET")

		r.ServeHTTP(res, req)

		if res.Code != c.expectedResponseCode{
			fmt.Println(c.description)
			t.Errorf("Error, expected: %d Received: %d",c.expectedResponseCode,res.Code)
		}else{
			fmt.Printf("%s: code result: %d and Code Expected: %d \n",c.description,res.Code,c.expectedResponseCode)
		}


	}
}

func (f *fakeDriver) Initialize(c *database.ConnectionDB) error {
	if c.Dbname == "service" && c.Ipaddress =="localhost"{
		return nil
	}else {
		return errors.New("Error Fake Initialize")
	}
}

func (f *fakeDriver) InsertEntity(i *data_model.Information) error {

	if i.Service != ""{
		return nil
	}else{
		return errors.New("Error Fake insertEntity")
	}
}

func (f *fakeDriver) GetEntity(field,searchItem string) (result []data_model.Information,err error) {
	return result, nil
}

func (f *fakeDriver) IsNew(field string, searchItem string) bool {
	return true
}

