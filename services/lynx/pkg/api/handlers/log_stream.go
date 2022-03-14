package handlers

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"lynx/pkg/clients/picus"
	"lynx/pkg/entities"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{}
)

func (h *Handlers) GetStreamLogsHandler(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	_, token, err := ws.ReadMessage()
	if err != nil {
		return err
	}
	userID, err := h.auth.GetUserID(string(token))
	if err != nil {
		return err
	}

	functionName := c.Param("name")
	ch := make(chan picus.LogResponse)
	ticker := time.NewTicker(25 * time.Second)

	function, err := h.ibis.GetFunctionByName(c.Request().Context(), userID, functionName)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(c.Request().Context())
	go h.picus.GetOnlineLogs(ctx, entities.FunctionQuery{ID: function.ID}, ch)

loop:
	for {
		var newItem picus.LogResponse
		var ok bool

		select {
		case newItem, ok = <-ch:
			if ok == false {
				break loop
			}
			ticker = time.NewTicker(25 * time.Second)
		case _ = <-ticker.C:
			newItem = picus.LogResponse{Timestamp: *ptypes.TimestampNow(), Message: "keepalive"}
		}
		err = ws.WriteJSON(newItem)
		if err != nil {
			break loop
		}
		_, _, err = ws.ReadMessage()
		if err != nil {
			break loop
		}
	}
	cancel()
	return nil
}
