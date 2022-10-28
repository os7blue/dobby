package common

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// CheckJsonUtil min: min item num;	max:max item num;	regx:item content regx rule
func CheckJsonUtil(max int, regx string, data string) error {
	dataArr := strings.Split(data, ",")

	if len(dataArr) == 0 {
		return errors.New("不可为空")
	}

	if len(dataArr) > max {
		return errors.New(fmt.Sprintf("最多%d条数据", max))
	}

	for _, v := range dataArr {
		b, err := regexp.Match(regx, []byte(v))
		if err != nil || !b {
			return errors.New(fmt.Sprintf("内容： %s 格式不正确", v))
		}
	}

	return nil
}

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
