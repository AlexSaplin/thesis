package nalogi

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"hippo/pkg/config"
	"hippo/pkg/entities"
)

type Client interface {
	SubmitReports([]entities.RunReport) error
}

type kafkaClient struct {
	cfg  config.NalogiClientConfig
	conn *kafka.Conn
}

func NewKafkaClient(cfg config.NalogiClientConfig) (Client, error) {
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Target, cfg.Topic, partition)
	if err != nil {
		return nil, err
	}
	return &kafkaClient{
		cfg:  cfg,
		conn: conn,
	}, nil
}


func (k *kafkaClient) SubmitReports(in []entities.RunReport) error {
	reports, err := serializeRunReports(in)
	if err != nil {
		return err
	}
	_, err = k.conn.WriteMessages(reports...)
	return err
}


type runReport struct {
	OwnerID   string  `json:"owner_id"`
	ModelID   string  `json:"model_id"`
	Duration  float64 `json:"duration"`
	Timestamp int64   `json:"ts"`
}

func serializeRunReports(in []entities.RunReport) (out []kafka.Message, err error) {
	out = make([]kafka.Message, len(in))
	for i, v := range in {
		report := runReport{
			OwnerID:   v.OwnerID.String(),
			ModelID:   v.ModelID.String(),
			Duration:  v.Duration.Seconds(),
			Timestamp: v.Timestamp.Unix(),
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