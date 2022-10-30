package bootstrap

import (
	"message-push/common"
	"message-push/model"
)

func dbInit() {

	err := common.DB.AutoMigrate(
		&model.ChannelInfo{},
	)
	if err != nil {
		return
	}

}
