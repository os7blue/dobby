package common

var cache = map[string]any{}
var cacheTime = map[string]int64{}

var LocalCache = new(localCache)

type localCache struct {
}

func (l *localCache) SetLocalCacheItem(key string, value any) {

	cache[key] = value

}

/*
SetLocalCacheItemWithTimeOut
outTime: time out -> ms
*/
func (l *localCache) SetLocalCacheItemWithTimeOut(key string, outTime uint, value any) {

	cache[key] = value

}

func GetLocalCacheItem[T any](key string) T {

	return cache[key].(T)

}
