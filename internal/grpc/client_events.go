package grpc

import (
	"context"
	"fmt"

	pb "gitlab.com/g6834/team41/tasks/api/events"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientEvents struct {
	addr string
}

func NewClientEvents(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (c *Client) CreateTask(ctx context.Context, objectId uint32) error {
	conn, err := grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("can't connect to events: %w", err)
	}
	defer conn.Close()

	analytics := pb.NewAnalyticsServiceClient(conn)

	in := pb.TaskRequest{
		ObjectId: objectId,
	}

	_, err = analytics.CreateTask(ctx, &in)
	if err != nil {
		return fmt.Errorf("events CreateTask error: %w", err)
	}
	return nil
}

func (c *Client) FinishTask(ctx context.Context, objectId uint32) error {
	conn, err := grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("can't connect to events: %w", err)
	}
	defer conn.Close()

	analytics := pb.NewAnalyticsServiceClient(conn)

	in := pb.TaskRequest{
		ObjectId: objectId,
	}

	_, err = analytics.FinishTask(ctx, &in)
	if err != nil {
		return fmt.Errorf("events CreateTask error: %w", err)
	}
	return nil
}

func (c *Client) CreateLetter(ctx context.Context, objectId uint32, email string) error {
	conn, err := grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("can't connect to events: %w", err)
	}
	defer conn.Close()

	analytics := pb.NewAnalyticsServiceClient(conn)

	in := pb.LetterRequest{
		ObjectId: objectId,
		Email:    email,
	}

	_, err = analytics.CreateLetter(ctx, &in)
	if err != nil {
		return fmt.Errorf("events CreateLetter error: %w", err)
	}
	return nil
}

func (c *Client) AcceptedLetter(ctx context.Context, objectId uint32, email string) error {
	conn, err := grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("can't connect to events: %w", err)
	}
	defer conn.Close()

	analytics := pb.NewAnalyticsServiceClient(conn)

	in := pb.LetterRequest{
		ObjectId: objectId,
		Email:    email,
	}

	_, err = analytics.AcceptedLetter(ctx, &in)
	if err != nil {
		return fmt.Errorf("events AcceptedLetter error: %w", err)
	}
	return nil
}

func (c *Client) DeclinedLetter(ctx context.Context, objectId uint32, email string) error {
	conn, err := grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("can't connect to events: %w", err)
	}
	defer conn.Close()

	analytics := pb.NewAnalyticsServiceClient(conn)

	in := pb.LetterRequest{
		ObjectId: objectId,
		Email:    email,
	}

	_, err = analytics.DeclinedLetter(ctx, &in)
	if err != nil {
		return fmt.Errorf("events DeclinedLetter error: %w", err)
	}
	return nil
}
