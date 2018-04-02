package models

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

var (
	rt            *roomTable
	roomTableOnce sync.Once
)

type roomTable struct {
	mux   sync.Mutex
	rooms map[string]Room
}

type Room interface {
	GetID() string
	FindClient(id string) (Client, error)
	Register(c Client)
}

type baseRoom struct {
	id      string
	clients map[string]Client
}

func GetRoomTable() *roomTable {
	roomTableOnce.Do(func() {
		rt = new(roomTable)
		rt.rooms = make(map[string]Room)
	})
	return rt
}

func (rt *roomTable) Register(room Room) {
	rt.mux.Lock()
	defer rt.mux.Unlock()
	rt.rooms[room.GetID()] = room
}

func (rt *roomTable) Find(id string) (Room, error) {
	room, ok := rt.rooms[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("not found room: %v", id))
	}
	return room, nil
}

func NewRoom(id string) Room {
	r := new(baseRoom)
	r.clients = make(map[string]Client, 0)
	r.id = id
	return r
}

func (r *baseRoom) Register(c Client) {
	r.clients[c.GetID()] = c
}

func (r baseRoom) GetID() string {
	return r.id
}

func (r baseRoom) FindClient(id string) (Client, error) {
	client, ok := r.clients[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("not found client: %v", id))
	}
	return client, nil
}
