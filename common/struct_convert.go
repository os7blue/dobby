package common

import (
	"encoding/json"
	"errors"
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
