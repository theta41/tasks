package kafka

import (
	"context"
	"errors"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Client struct {
	Writer          *kafka.Writer
	topicAnalytics  string
	topicMailsender string
}

func NewClient(brokers []string, topicAnalytics, topicMailsender string) (*Client, error) {
	if len(brokers) == 0 || brokers[0] == "" {
		return nil, errors.New("invalid kafka configuration")
	}

	c := Client{
		topicAnalytics:  topicAnalytics,
		topicMailsender: topicMailsender,
	}

	c.Writer = &kafka.Writer{
		Addr:     kafka.TCP(brokers[0]),
		Balancer: &kafka.LeastBytes{},
	}

	return &c, nil
}

func (c *Client) PublishAnalytics(key, value []byte) error {
	return c.publish(c.topicAnalytics, key, value)
}

func (c *Client) PublishEmail(key, value []byte) error {
	return c.publish(c.topicMailsender, key, value)
}

func (c *Client) publish(topic string, key, value []byte) error {
	msg := kafka.Message{
		Topic: topic,
		Key:   key,
		Value: value,
	}
	err := c.Writer.WriteMessages(context.TODO(), msg)
	if err != nil {
		return fmt.Errorf("could not write message %w", err)
	}

	return nil
}
