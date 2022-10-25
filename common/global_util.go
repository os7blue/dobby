package common

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

func StructConvert[F any, T any](from F, to *T) error {

	j, err := json.Marshal(from)

	if err != nil {
		return errors.New("数据转换失败")
	}

	err = json.Unmarshal(j, to)

	if err != nil {
		return errors.New("数据转换失败")
	}

	return nil
}

func RandCodeString(len int) string {

	rand.Seed(time.Now().UnixNano())

	code := ""

	for i := 0; i < len; i++ {
		code = fmt.Sprintf("%s%d", code, rand.Intn(10))
	}

	return code

}
