package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"time"
)

type MetricItem struct {
	Day       time.Time `json:"day"`
	RunTime   float64   `json:"runTime"`
	RunCount  int       `json:"runCount"`
	LoadTime  float64   `json:"loadTime"`
	LoadCount int       `json:"loadCount"`
}

type GetMetricsResponse struct {
	TotalRunTime   float64      `json:"totalRunTime"`
	TotalRunCount  int          `json:"totalRunCount"`
	TotalLoadTime  float64      `json:"totalLoadTime"`
	TotalLoadCount int          `json:"totalLoadCount"`
	Details        []MetricItem `json:"details"`
}

func (h *Handlers) GetMetricsHandler(c echo.Context) error {
	offset := c.Request().URL.Query()["offset"][0]
	offsetNumber, _ := strconv.Atoi(offset)
	userID := c.Get(MwUserIDKey).(uuid.UUID)

	name := c.Param("name")
	function, err := h.ibis.GetFunctionByName(c.Request().Context(), userID, name)
	if err != nil {
		return err
	}
	functionID := function.ID.String()

	value, err := h.prom.GetFunctionMetrics(c.Request().Context(), functionID, offsetNumber+1)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var response GetMetricsResponse
	response.TotalRunTime, _ = strconv.ParseFloat(value.Metrics[len(value.Metrics)-1].RunTime, 64)
	response.TotalLoadTime, _ = strconv.ParseFloat(value.Metrics[len(value.Metrics)-1].LoadTime, 64)
	response.TotalRunCount, _ = strconv.Atoi(value.Metrics[len(value.Metrics)-1].RunCount)
	response.TotalLoadCount, _ = strconv.Atoi(value.Metrics[len(value.Metrics)-1].LoadCount)
	response.Details = make([]MetricItem, offsetNumber)

	for i := 0; i < offsetNumber; i++ {
		runPrev, _ := strconv.ParseFloat(value.Metrics[i].RunTime, 64)
		runNext, _ := strconv.ParseFloat(value.Metrics[i+1].RunTime, 64)

		loadPrev, _ := strconv.ParseFloat(value.Metrics[i].LoadTime, 64)
		loadNext, _ := strconv.ParseFloat(value.Metrics[i+1].LoadTime, 64)

		runCountPrev, _ := strconv.Atoi(value.Metrics[i].RunCount)
		runCountNext, _ := strconv.Atoi(value.Metrics[i+1].RunCount)

		loadCountPrev, _ := strconv.Atoi(value.Metrics[i].LoadCount)
		loadCountNext, _ := strconv.Atoi(value.Metrics[i+1].LoadCount)
		response.Details[i] = MetricItem{
			Day:       time.Now().AddDate(0, 0, i-offsetNumber+1),
			RunTime:   runNext - runPrev,
			RunCount:  runCountNext - runCountPrev,
			LoadTime:  loadNext - loadPrev,
			LoadCount: loadCountNext - loadCountPrev,
		}
	}
	return c.JSON(http.StatusOK, response)
}
