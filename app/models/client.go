package models

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

var (
	clientTableOne  *clientTable
	clientTableOnce *sync.Once
)

type clientTable struct {
	mux     sync.Mutex
	clients map[string]Client
}

type Client interface {
	GetID() string
	GetConn() *websocket.Conn
	GetRoom() Room
	SetRoom(room Room)
}

type baseClient struct {
	id   string
	conn *websocket.Conn
	room Room
}

func GetClientTable() *clientTable {
	clientTableOnce.Do(func() {
		clientTableOne = new(clientTable)
	})
	return clientTableOne
}

func (ct *clientTable) Register(client Client) {
	ct.mux.Lock()
	defer ct.mux.Unlock()
	ct.clients[client.GetID()] = client
}

func (ct clientTable) Find(id string) (Client, error) {
	client, ok := ct.clients[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("not found client: %v", id))
	}
	return client, nil
}

func NewClient(id string, conn *websocket.Conn) Client {
	return &baseClient{
		id:   id,
		conn: conn,
	}
}

func (c baseClient) GetID() string {
	return c.id
}

func (c baseClient) GetConn() *websocket.Conn {
	return c.conn
}

func (c baseClient) GetRoom() Room {
	return c.room
}

func (c *baseClient) SetRoom(room Room) {
	c.room = room
}
