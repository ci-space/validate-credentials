package credentials

type Store interface {
	Get(key string) (string, bool)
}
