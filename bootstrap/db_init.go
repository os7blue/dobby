package bootstrap

import (
	"message-push/common"
	"message-push/model"
)

func dbInit() {

	err := common.DB.AutoMigrate(
		&model.PushChannelInfo{},
		&model.PushChannelLine{},
	)
	if err != nil {
		return
	}

}
