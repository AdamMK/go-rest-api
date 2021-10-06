package main

import (
	"bytes"
	"encoding/json"
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

func serverWithDocument() *Server {
	tS := ForTestServer()
	id, _ := uuid.Parse("25f2d428-0670-4edd-8a4d-df3e0595d716")
	d := Doc{ id, "Sell", Content{"Car", "Volvo"}, "Adam" }
	tS.Documents = append(tS.Documents, d)
	return tS
}



func TestListDocuments(t *testing.T) {

	request, _ := http.NewRequest("GET", "/documents", nil)
	response := httptest.NewRecorder()

	tS := serverWithDocument()
	tS.ServeHTTP(response, request)

	var docs []Doc
	json.NewDecoder(response.Body).Decode(&docs)

	assert.Equal(t,200, response.Code)
	assert.Equal(t, tS.Documents, docs)

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

	request, _ := http.NewRequest("GET", "/documents/25f2d428-0670-4edd-8a4d-df3e0595d716", nil)
	response := httptest.NewRecorder()
	var d Doc

	tS := serverWithDocument()
	tS.ServeHTTP(response, request)
	json.NewDecoder(response.Body).Decode(&d)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, tS.Documents[0], d )

}

func TestDeleteDocument(t *testing.T) {

	request, _ := http.NewRequest("DELETE", "/documents/25f2d428-0670-4edd-8a4d-df3e0595d716", nil)
	response := httptest.NewRecorder()
	var d Doc

	tS := serverWithDocument()
	tS.ServeHTTP(response, request)
	json.NewDecoder(response.Body).Decode(&d)


	assert.Equal(t, 200, response.Code)
	assert.Equal(t, 0, len(tS.Documents))

}

func TestUpdateDocument(t *testing.T)  {

	id, _ := uuid.Parse("25f2d428-0670-4edd-8a4d-df3e0595d716")
	updatedDoc := Doc{ id, "Buy", Content{"House", "Detached"}, "John Snow" }
	w := new(bytes.Buffer)
	json.NewEncoder(w).Encode(&updatedDoc)

	request, _ := http.NewRequest("PUT", "/documents/25f2d428-0670-4edd-8a4d-df3e0595d716", w)
	response := httptest.NewRecorder()

	var d Doc
	tS := serverWithDocument()
	tS.ServeHTTP(response, request)
	json.NewDecoder(response.Body).Decode(&d)


	assert.Equal(t, 200, response.Code)
	assert.Equal(t, updatedDoc, d)

}
