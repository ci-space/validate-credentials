package credentials

import "os"

type EnvStore struct{}

func NewEnvStore() *EnvStore {
	return &EnvStore{}
}

func (s *EnvStore) Get(key string) (string, bool) {
	return os.LookupEnv(key)
}
