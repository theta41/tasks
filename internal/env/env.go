package env

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"gitlab.com/g6834/team41/tasks/internal/kafka"
	"os"

	"gitlab.com/g6834/team41/tasks/internal/grpc"

	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/cfg"
	"gitlab.com/g6834/team41/tasks/internal/pg"
	"gitlab.com/g6834/team41/tasks/internal/ports"
	"gitlab.com/g6834/team41/tasks/internal/repositories"
)

type Environment struct {
	C cfg.Config

	K         ports.Queue
	LR        repositories.Letters
	TR        repositories.Tasks
	Auth      ports.AuthService
	Analytics ports.AnalyticsService
}

var E *Environment

var (
	ConfigPath = "CONFIG_PATH"
)

func init() {
	// Get config path from environment variable
	path := os.Getenv(ConfigPath)
	if path == "" {
		path = "config.yaml"
	}

	var err error
	E, err = NewEnvironment(path)
	if err != nil {
		logrus.Panic(fmt.Errorf("failed to load config: %w", err))
	}

	configureLogger()

	// Create postgres connection
	logrus.Debug("Connecting to postgres...")
	db, err := pg.NewPG(E.C.DB.Host, E.C.DB.User, E.C.DB.Password, E.C.DB.Name, E.C.DB.SSL, E.C.DB.Port)
	if err != nil {
		logrus.Panic(fmt.Errorf("failed to connect to postgres: %w", err))
	}

	initSentry(E.C.SentryDSN)

	E.K, err = kafka.NewClient(E.C.Kafka.Brokers, E.C.Kafka.Topic)
	if err != nil {
		logrus.Panic(fmt.Errorf("failed to connect to kafka: %w", err))
	}

	E.LR = pg.NewLetters(db)
	E.TR = pg.NewTasks(db)
	E.Auth = grpc.NewClient(E.C.AuthAddress)
	E.Analytics = grpc.NewClientEvents(E.C.AnalyticsAddress)
}

func NewEnvironment(yamlFile string) (*Environment, error) {
	conf, err := cfg.NewConfig(yamlFile)
	if err != nil {
		return nil, err
	}

	return &Environment{C: *conf}, nil
}

func initSentry(dsn string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   dsn,
		Debug: true,
	})
	if err != nil {
		panic(fmt.Errorf("sentry.Init: %w", err))
	}
}

func configureLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if E.C.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
