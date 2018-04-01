package models

type BaseRoom interface {
	GetID() string
	FindClient(id string) Client
}

type room struct {
	id      string
	clients map[string]Client
}
