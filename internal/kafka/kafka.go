package kafka

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type Client struct {
	Writer *kafka.Writer
}

func NewClient(brokers []string, topic string) (*Client, error) {
	if len(brokers) == 0 || brokers[0] == "" {
		return nil, errors.New("invalid kafka configuration")
	}

	c := Client{}

	c.Writer = &kafka.Writer{
		Addr:     kafka.TCP(brokers[0]),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &c, nil
}

func (c *Client) Publish(key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}
	err := c.Writer.WriteMessages(context.TODO(), msg)
	if err != nil {
		return fmt.Errorf("could not write message %w", err)
	}

	return nil
}
