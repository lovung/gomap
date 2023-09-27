package gomap

type Map[K comparable, V any] interface {
	Store(key K, val V)
	Load(key K) (V, bool)
	LoadAndDelete(key K) (V, bool)
	Delete(key K)
	Contain(key K) bool
	Clear()
}

type Option func(o *option)

type option struct {
	cap int // 0 mean no cap
}

func WithCap(cap int) Option {
	return func(o *option) {
		o.cap = cap
	}
}
