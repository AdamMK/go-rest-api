package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
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

func (s *Server) add() []Doc{
	id, _ := uuid.Parse("25f2d428-0670-4edd-8a4d-df3e0595d716")
	var d = Doc{ id, "Sell", Content{"Car", "Volvo"}, "Adam" }
	newS := append(s.Documents, d)
	return newS
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

	id, _ := uuid.Parse("25f2d428-0670-4edd-8a4d-df3e0595d716")
	var d = Doc{ id, "Sell", Content{"Car", "Volvo"}, "Adam" }
	x := append(ForTestServer().Documents, d)


	fmt.Println(x)

	request, _ := http.NewRequest("GET", "/documents/25f2d428-0670-4edd-8a4d-df3e0595d716", nil)
	fmt.Println(ForTestServer().Documents)
	response := httptest.NewRecorder()
	ForTestServer().ServeHTTP(response, request)
	json.NewEncoder(response.Body).Encode(&x)

	assert.Equal(t, 404, response.Code)

}
