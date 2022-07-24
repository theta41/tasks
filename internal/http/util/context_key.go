package util

type ContextKey string

func (c ContextKey) String() string {
	return "tasks context key " + string(c)
}
