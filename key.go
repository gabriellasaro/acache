package acache

type Key string

func (k Key) Add(key string) Key {
	return k + ":" + Key(key)
}
