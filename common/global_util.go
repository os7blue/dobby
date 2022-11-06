package common

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const (
	EmailRegx = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	IpV4Regx  = "((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}"
	UrlRegx   = "^([\\w-]+\\.)+[\\w-]+(\\/[\\w-.\\/?%&=]*)?$"
	UUID      = "[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}"
)

var GlobalUtil = new(globalUtil)

type globalUtil struct {
}

func (g *globalUtil) checkArrStr() {

}

// CheckArrStr CheckJsonUtil min: min item num;	max:max item num;	regx:item content regx rule
func (g *globalUtil) CheckArrStr(min int, max int, regx string, data string) error {

	if data != "" {

		dataArr := strings.Split(data, ",")

		if min > 0 && len(dataArr) < min {
			return errors.New(fmt.Sprintf("最少有%d条数据", min))
		}

		if max != 0 && len(dataArr) > max {
			return errors.New(fmt.Sprintf("最多%d条数据", max))
		}

		for _, v := range dataArr {

			b, err := regexp.Match(regx, []byte(v))
			if err != nil || !b {
				return errors.New(fmt.Sprintf("内容： %s 格式不正确", v))
			}

		}

	} else if min > 0 {
		return errors.New(fmt.Sprintf("至少有%d条数据", min))
	}

	return nil
}

func (g *globalUtil) RandCodeString(len int) string {

	rand.Seed(time.Now().UnixNano())

	code := ""

	for i := 0; i < len; i++ {
		code = fmt.Sprintf("%s%d", code, rand.Intn(10))
	}

	return code

}
