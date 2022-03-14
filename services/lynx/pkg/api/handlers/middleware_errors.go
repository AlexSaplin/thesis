package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const StatusClientClosedRequest = 499

func (h *Handlers) ErrorsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			st, ok := status.FromError(err)
			if !ok {
				return err
			}
			switch st.Code() {
			case codes.NotFound:
				return echo.NewHTTPError(http.StatusNotFound, st.Message())
			case codes.AlreadyExists:
				return echo.NewHTTPError(http.StatusConflict, st.Message())
			case codes.InvalidArgument:
				return echo.NewHTTPError(http.StatusBadRequest, st.Message())
			case codes.Unavailable:
				return c.NoContent(http.StatusProcessing)
			case codes.FailedPrecondition:
				return echo.NewHTTPError(http.StatusPreconditionFailed, st.Message())
			case codes.PermissionDenied:
				return echo.NewHTTPError(http.StatusForbidden, st.Message())
			case codes.Unauthenticated:
				return echo.NewHTTPError(http.StatusUnauthorized, st.Message())
			case codes.Canceled:
				return c.NoContent(StatusClientClosedRequest)
			default:
				return err
			}
		}
		return nil
	}
}
