package constant

var PlanKeyKey = cacheKeyOption{
	"plan-%s",
	7200,
}

var PlanIdKey = cacheKeyOption{
	"plan-%d",
	7200,
}

var CookieKey = cacheKeyOption{
	"u-%s",
	7200,
}

type cacheKeyOption struct {
	KeyTemp string
	Time    int64
}
