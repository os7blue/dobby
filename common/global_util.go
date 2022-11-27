package common

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"math/rand"
	"net/http"
	"net/url"
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

func (g *globalUtil) ArrStrItemFormat(arrStr string, formatTemp string, separator string, newSeparator string) string {

	arr := strings.Split(arrStr, separator)
	newArrStr := ""
	for i, item := range arr {

		if i != 0 && len(arr) > 1 {
			newArrStr += newSeparator + fmt.Sprintf(formatTemp, item)
		} else {
			newArrStr += fmt.Sprintf(formatTemp, item)
		}

	}
	return newArrStr

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

		if regx != "" {
			for _, v := range dataArr {

				b, err := regexp.Match(regx, []byte(v))
				if err != nil || !b {
					return errors.New(fmt.Sprintf("内容： %s 格式不正确", v))
				}

			}
		}

	} else if min > 0 {
		return errors.New(fmt.Sprintf("至少有%d条数据", min))
	}

	return nil
}

func (g *globalUtil) SendSimpleGet(getUrl string) (string, error) {

	parseUrl, err := url.Parse(getUrl)
	if err != nil {
		return "", err
	}

	params, err := url.ParseQuery(parseUrl.RawQuery)
	if err != nil {
		return "", err
	}
	parseUrl.RawQuery = params.Encode()

	rq, err := http.NewRequest("GET", parseUrl.String(), nil)
	if err != nil {
		return "", err
	}
	rq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	rp, err := http.DefaultClient.Do(rq)
	if err != nil {
		return "", err
	}
	defer rp.Body.Close()
	rpBody, err := io.ReadAll(rp.Body)
	if err != nil {
		return "", err
	}

	return string(rpBody), nil
}

// SendSimplePost run a simple post request
func (g *globalUtil) SendSimplePost(url string, param any) (string, error) {

	j, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	rq, err := http.NewRequest("POST", url, strings.NewReader(string(j)))
	if err != nil {
		return "", err
	}

	rq.Header.Set("Content-Type", "application/json;charset=utf-8")

	rp, err := http.DefaultClient.Do(rq)
	if err != nil {
		return "", err

	}
	defer rp.Body.Close()
	rpBody, err := io.ReadAll(rp.Body)
	if err != nil {
		return "", err
	}

	return string(rpBody), nil
}

func (g *globalUtil) RandCodeString(len int) string {

	rand.Seed(time.Now().UnixNano())

	code := ""

	for i := 0; i < len; i++ {
		code = fmt.Sprintf("%s%d", code, rand.Intn(10))
	}

	return code

}
