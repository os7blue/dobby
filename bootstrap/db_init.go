package bootstrap

import (
	"message-push/common"
	"message-push/model"
)

func dbInit() {

	err := common.DB.AutoMigrate(
		&model.ChannelInfo{},
		&model.ChannelLine{},
	)
	if err != nil {
		return
	}

}
