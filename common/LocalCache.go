package common

var cache = map[string]any{}

func SetLocalCacheItem(key string, value any) {

	cache[key] = value

}

func GetLocalCacheItem[T any](key string) T {

	return cache[key].(T)

}
