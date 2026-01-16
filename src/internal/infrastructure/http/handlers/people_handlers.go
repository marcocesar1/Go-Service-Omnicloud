package handlers

import "net/http"

type PeopleHandlers struct {
}

func CreatePeopleHandlers() *PeopleHandlers {
	return &PeopleHandlers{}
}

func (p *PeopleHandlers) FindOne() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FindOne"))
	}
}

func (p *PeopleHandlers) FindAll() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FindAll"))
	}
}

func (p *PeopleHandlers) Create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Create"))
	}
}

func (p *PeopleHandlers) Patch() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Update"))
	}
}
