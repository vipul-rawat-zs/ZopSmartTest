package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_Handler(t *testing.T) {
	// initializing db
	conf := mysqlConfig{
		host:     "localhost",
		user:     "vips",
		password: "1234",
		port:     "3306",
		db:       "test",
	}
	var err error
	db, err = connectToMySQL(conf)
	if err != nil {
		t.Errorf("could not connect to sql, err:%v", err)
	}

	testcases := []struct {
		// input
		method string
		body   []byte
		// output
		expectedStatusCode int
		expectedResponse   []byte
	}{
		{"GET", nil, http.StatusOK, []byte(`[{"Name":"Hippo","Age":10}]`)},
		{"POST", []byte(`{"Name":"Dog","Age":12}`), http.StatusOK, []byte(`success`)},
		{"DELETE", nil, http.StatusMethodNotAllowed, nil},
	}

	for _, v := range testcases {
		req := httptest.NewRequest(v.method, "/animal", bytes.NewReader(v.body))
		w := httptest.NewRecorder()

		h := http.HandlerFunc(handler)
		h.ServeHTTP(w, req)

		if w.Code != v.expectedStatusCode {
			t.Errorf("Expected %v\tGot %v", v.expectedStatusCode, w.Code)
		}

		expected := bytes.NewBuffer(v.expectedResponse)
		if !reflect.DeepEqual(w.Body, expected) {
			t.Errorf("Expected %v\tGot %v", expected.String(), w.Body.String())
		}
	}
}
