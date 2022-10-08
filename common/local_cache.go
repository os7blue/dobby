package common

var cache = map[string]any{}
var cacheTime = map[string]int64{}

func SetLocalCacheItem(key string, value any) {

	cache[key] = value

}

/*
SetLocalCacheItemWithTimeOut
outTime: time out -> ms
*/
func SetLocalCacheItemWithTimeOut(key string, outTime uint, value any) {

	cache[key] = value

}

func GetLocalCacheItem[T any](key string) T {

	return cache[key].(T)

}
