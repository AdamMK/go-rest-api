package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//helper method for getting http requests in one place
func (s *Server) requests() {
	s.HandleFunc("/documents", s.listDocuments()).Methods("GET")
	s.HandleFunc("/documents/{id}", s.findDocument()).Methods("GET")
	s.HandleFunc("/documents", s.createDocument()).Methods("POST")
	s.HandleFunc("/documents/{id}", s.deleteDocument()).Methods("DELETE")
	s.HandleFunc("/documents/{id}", s.updateDocument()).Methods("PUT")
}

func (s *Server) listDocuments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.Documents); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) findDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idStr := mux.Vars(r)["id"]
		docId, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, doc := range s.Documents {
			if doc.ID == docId {
				json.NewEncoder(w).Encode(doc)
				return
			}
		}
		http.Error(w, "", http.StatusNotFound)
	}
}

func (s *Server) createDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d Doc
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		d.ID = uuid.New()
		s.Documents = append(s.Documents, d)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		if err := json.NewEncoder(w).Encode(d); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) deleteDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for d, document := range s.Documents {
			if document.ID == id {
				s.Documents = append(s.Documents[:d], s.Documents[d+1:]... )
				return
			}
		}
		http.Error(w, "", http.StatusNotFound)
	}
}

func (s *Server) updateDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { 
		var nd Doc
		if err := json.NewDecoder(r.Body).Decode(&nd); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for d, document := range s.Documents {
			if document.ID == id {
				nd.ID = id
				s.Documents = append(s.Documents[:d], nd)
			}
		 }

		 //return edited document
		 w.Header().Set("Content-Type", "application/json")
		 if err := json.NewEncoder(w).Encode(nd); err != nil {
			 http.Error(w, err.Error(), http.StatusInternalServerError)
			 return
		 }
	}

}