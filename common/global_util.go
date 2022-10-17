package common

import (
	"fmt"
	"math/rand"
	"time"
)

func RandCodeString(len int) string {

	rand.Seed(time.Now().UnixNano())

	code := ""

	for i := 0; i < len; i++ {
		code = fmt.Sprintf("%s%d", code, rand.Intn(10))
	}

	return code

}
