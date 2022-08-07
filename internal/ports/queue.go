package ports

type Queue interface {
	PublishAnalytics(key, value []byte) error
	PublishEmail(key, value []byte) error
}
