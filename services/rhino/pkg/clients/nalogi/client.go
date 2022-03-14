package nalogi

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"rhino/pkg/config"
	"rhino/pkg/entities"
)

type Client interface {
	SubmitReports([]entities.RunReport) error
}

type kafkaClient struct {
	cfg    config.NalogiClientConfig
	writer *kafka.Writer
}

func NewKafkaClient(cfg config.NalogiClientConfig) (Client, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{cfg.Target},
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &kafkaClient{
		cfg:    cfg,
		writer: writer,
	}, nil
}

func (k *kafkaClient) SubmitReports(in []entities.RunReport) error {
	reports, err := serializeRunReports(in)
	if err != nil {
		return err
	}
	err = k.writer.WriteMessages(context.Background(), reports...)
	return err
}

type runReport struct {
	OwnerID      string  `json:"owner_id"`
	FunctionID   string  `json:"object_id"`
	RunDuration  float64 `json:"run_duration"`
	LoadDuration float64 `json:"load_duration"`
	Timestamp    int64   `json:"ts"`
}

func serializeRunReports(in []entities.RunReport) (out []kafka.Message, err error) {
	out = make([]kafka.Message, len(in))
	for i, v := range in {
		report := runReport{
			OwnerID:      v.OwnerID.String(),
			FunctionID:   v.FunctionID.String(),
			RunDuration:  v.RunDuration.Seconds(),
			LoadDuration: v.LoadDuration.Seconds(),
			Timestamp:    v.Timestamp.Unix(),
		}
		var jsonData []byte
		jsonData, err = json.Marshal(report)
		if err != nil {
			return
		}
		out[i] = kafka.Message{
			Value: jsonData,
		}
	}
	return
}
