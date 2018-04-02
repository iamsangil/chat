package controller

import (
	"github.com/iamsangil/chat/app/models"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var (
	upgrader = websocket.Upgrader{}
)

func ChatController(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	msg := new(models.Message)
	var client models.Client
	for {
		if err := ws.ReadJSON(msg); err != nil {
			c.Logger().Error(err)
		}

		switch msg.Type {
		case models.TypeConnect:
			client = models.NewClient(msg.ClientID, ws)
			log.Info(msg.ClientID)
			log.Info(msg.RoomID)
			roomTable := models.GetRoomTable()
			log.Info(*roomTable)
			room, err := roomTable.Find(msg.RoomID)
			if err != nil {
				room = models.NewRoom(msg.RoomID)
				roomTable.Register(room)
				c.Logger().Error(err)
			}
			room.Register(client)
			client.SetRoom(room)

			err = ws.WriteMessage(websocket.TextMessage, []byte("connect"))
			if err != nil {
				c.Logger().Error(err)
			}

		case models.TypeSend:
			room := client.GetRoom()
			toClient, err := room.FindClient(msg.ClientID)
			if err != nil {
				c.Logger().Error(err)
			}
			toClient.GetConn().WriteMessage(websocket.TextMessage, []byte(msg.Data))

		case models.TypeDisconnect:
			err := ws.WriteMessage(websocket.TextMessage, []byte("disconnect"))
			if err != nil {
				c.Logger().Error(err)
			}

		}
	}
}
