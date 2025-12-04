package mongo

import "time"

type Option func(*Mongo)

func MaxPoolSize(size int) Option {
	return func(m *Mongo) {
		m.maxPoolSize = size
	}
}

func ConnAttempts(attempts int) Option {
	return func(m *Mongo) {
		m.connAttempts = attempts
	}
}

func ConnTimeout(timeout time.Duration) Option {
	return func(m *Mongo) {
		m.connTimeout = timeout
	}
}
