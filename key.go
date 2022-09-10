package acache

type Key string

func (k Key) Add(key string) Key {
	return k + ":" + Key(key)
}

func NewKey[K ~string](key K) Key {
	return Key(key)
}
