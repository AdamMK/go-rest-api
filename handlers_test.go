package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func ForTestServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		Documents: []Doc{},
	}
	s.requests()
	return s
}


func TestListDocuments(t *testing.T) {

	request, _ := http.NewRequest("GET", "/documents", nil)
	response := httptest.NewRecorder()
	ForTestServer().ServeHTTP(response, request)

	//no documents added
	assert.Equal(t,404, response.Code)
}

func TestCreateDocument(t *testing.T) {

	exampleJson :=
		`{"title": "Selling contract","content": {"header": "Sell a car", "data": "Volvo V40"},"signee": "Mr Bean"}`

	r := strings.NewReader(exampleJson)
	request, _ := http.NewRequest("POST", "/documents", r)
	response := httptest.NewRecorder()
	var d Doc

	ForTestServer().ServeHTTP(response, request)
	json.NewDecoder(response.Body).Decode(&d)
	assert.Equal(t, 201, response.Code)
	assert.Equal(t, "Selling contract" , d.Title)
}

func TestFindDocument(t *testing.T) {

	var d Doc

	request, _ := http.NewRequest("GET", "/documents/25f2d428-0670-4edd-8a4d-df3e0595d716", nil)
	response := httptest.NewRecorder()
	ForTestServer().ServeHTTP(response, request)
	json.NewEncoder(response.Body).Encode(&d)

	assert.Equal(t, 404, response.Code)

}
