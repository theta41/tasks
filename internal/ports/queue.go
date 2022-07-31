package ports

type Queue interface {
	Publish(key, value []byte) error
}
