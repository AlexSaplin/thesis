package picus

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	picus "lynx/pkg/clients/picus/pb"
)

type LogResponse struct {
	Timestamp timestamp.Timestamp `json:"time"`
	Message   string              `json:"message"`
}

type FullLogResponse struct {
	Log []LogResponse `json:"log"`
}

func newFullLogResponse(log []*picus.LogEntry) (result FullLogResponse) {
	for _, le := range log {
		result.Log = append(result.Log, LogResponse{
			Timestamp: *le.Time,
			Message:   le.Message,
		})
	}
	return
}
