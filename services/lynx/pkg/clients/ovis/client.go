package ovis

import (
	"encoding/json"

	"github.com/streadway/amqp"

	"lynx/pkg/config"
	"lynx/pkg/entities"
)

type Client interface {
	VerifyModel(model entities.Model) error
}

type message struct {
	ModelID string `json:"model_id"`
}

type RabbitMQClient struct {
	channel *amqp.Channel
	cfg     config.OvisClientConfig
}

func NewRabbitMQClient(cfg config.OvisClientConfig) (result *RabbitMQClient, err error) {
	client := &RabbitMQClient{
		cfg: cfg,
	}
	if err = client.init(); err != nil {
		return
	}
	return client, nil
}

func (c *RabbitMQClient) init() error {
	conn, err := amqp.Dial(c.cfg.Target)
	if err != nil {
		return err
	}

	c.channel, err = conn.Channel()
	if err != nil {
		return err
	}

	_, err = c.channel.QueueDeclare(
		c.cfg.Queue,
		false,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (c *RabbitMQClient) VerifyModel(model entities.Model) error {
	body, err := json.Marshal(message{
		ModelID: model.ID.String(),
	})
	if err != nil {
		return err
	}

	return c.channel.Publish(
		"",
		c.cfg.Queue,
		false,
		false,
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        body,
		},
	)
}