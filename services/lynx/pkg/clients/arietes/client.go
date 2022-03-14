package arietes

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"lynx/pkg/config"
	"lynx/pkg/entities"
)

type Client interface {
	SubmitFunctionQuery(in entities.FunctionQuery) error
}

type kafkaClient struct {
	cfg  config.ArietesClientConfig
	writer *kafka.Writer
}

func NewKafkaClient(cfg config.ArietesClientConfig) (Client, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
    	Brokers: []string{cfg.Target},
    	Topic: cfg.Topic,
    	Balancer: &kafka.LeastBytes{},
	})
	return &kafkaClient{
		cfg:  cfg,
		writer: writer,
	}, nil
}

func (k *kafkaClient) SubmitFunctionQuery(in entities.FunctionQuery) error {
	query, err := serializeFunctionQuery(in)
	if err != nil {
		return err
	}
	err = k.writer.WriteMessages(context.Background(), query)
	return err
}

type functionQuery struct {
	ID string `json:"function_id"`
}

func serializeFunctionQuery(in entities.FunctionQuery) (out kafka.Message, err error) {
	query := functionQuery{
		ID: in.ID.String(),
	}
	var jsonData []byte
	jsonData, err = json.Marshal(query)
	if err != nil {
		return
	}
	out = kafka.Message{
		Value: jsonData,
	}
	return
}
