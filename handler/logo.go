package handler

import (
	"fmt"
	"net/http"
)

type Logo struct {
}

func (o *Logo) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a logo")
}

func (o *Logo) List(w http.ResponseWriter, t *http.Request) {
	fmt.Println("List all logos")
}

func (o *Logo) GetByID(w http.ResponseWriter, t *http.Request) {
	fmt.Println("Get a logo by ID")
}

func (o *Logo) UpdateByID(w http.ResponseWriter, t *http.Request) {
	fmt.Println("Update a logo by ID")
}

func (o *Logo) DeleteByID(w http.ResponseWriter, t *http.Request) {
	fmt.Println("Delete a logo by ID")
}
