package handlers

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

const MwUserIDKey = "MwUserIDKey"
const MWConstTokenKey = "MWConstTokenKey"

func (h *Handlers) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	userMap := make(map[string]uuid.UUID)
	for _, user := range h.users {
		userID, err := uuid.FromString(user.UserID)
		if err != nil {
			panic(err)
		}
		userMap[user.Token] = userID
	}
	return func(c echo.Context) (err error) {
		token := c.Request().Header.Get("X-Token")
		if token == "" {
			return ErrUnauthenticated
		}
		constToken := true
		userID, ok := userMap[token]
		if !ok {
			userID, err = h.auth.GetUserID(token)
			if err != nil {
				return
			}
			constToken = false
		}
		c.Set(MwUserIDKey, userID)
		c.Set(MWConstTokenKey, constToken)
		return next(c)
	}
}
