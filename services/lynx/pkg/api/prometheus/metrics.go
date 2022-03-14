package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/kafka-go"
)

type runReport struct {
	OwnerID      string  `json:"owner_id"`
	FunctionID   string  `json:"object_id"`
	RunDuration  float64 `json:"run_duration"`
	LoadDuration float64 `json:"load_duration"`
	Timestamp    int64   `json:"ts"`
}

var (
	avgRunTime = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "run_duration",
		Help: "Represents duration of running a function",
	},
		[]string{"function_id", "owner_id"},
	)

	avgLoadTime = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "load_duration",
		Help: "Represents duration of loading a function",
	},
		[]string{"function_id", "owner_id"},
	)
)

func runDurationUpdater() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "run_reports",
	})
	ctx := context.Background()
	var report runReport
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("ERROR: read from kafka failed %s\n", err.Error())
			return
		}
		err = json.Unmarshal(m.Value, &report)
		if err != nil {
			fmt.Printf("ERROR: Unmarshal failed on %s with %s\n", m.Value, err.Error())
			return
		}
		avgRunTime.WithLabelValues(report.FunctionID, report.OwnerID).Observe(report.RunDuration)
		avgLoadTime.WithLabelValues(report.FunctionID, report.OwnerID).Observe(report.LoadDuration)
	}
}

func init() {
	prometheus.MustRegister(avgRunTime)
	prometheus.MustRegister(avgLoadTime)
	go runDurationUpdater()
}
