package bootstrap

import (
	"dobby/common"
	"dobby/model"
)

func dbInit() {

	err := common.DB.AutoMigrate(
		&model.ChannelInfo{},
	)
	if err != nil {
		return
	}

}
