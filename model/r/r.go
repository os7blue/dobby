package R

type R struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Count int    `json:"count"`
}

func Success() R {
	return R{
		Code: 1,
		Msg:  "success",
	}
}

func SuccessWithData[T any](data T) R {
	return R{
		Code: 1,
		Msg:  "success",
		Data: data,
	}
}

func SuccessWithDataCount[T any](data T, count int) R {
	return R{
		Code:  1,
		Msg:   "success",
		Data:  data,
		Count: count,
	}
}

func Fail() R {
	return R{
		Code: 0,
		Msg:  "fail",
	}
}

func FailWithData[T any](data T) R {
	return R{
		Code: 0,
		Msg:  "fail",
		Data: data,
	}
}
