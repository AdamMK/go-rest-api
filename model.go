package main

import (
	"github.com/google/uuid"
)

type Doc struct {
	ID 		uuid.UUID 	`json:"id"`
	Title   string   	`json:"title"`
	Content Content 	`json:"content"`
	Signee  string   	`json:"signee"`
}

type Content struct {
	Header 	string 		`json:"header"`
	Data 	string 		`json:"data"`
}