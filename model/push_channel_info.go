package model

import "message-push/config"

//item -> line -> info

type PushChannelInfo struct {
	Id         uint
	Name       string
	Status     int
	Key        string
	CreateTime uint
}

func (m *PushChannelInfo) create(info PushChannelInfo) {

	config.DB.Create(info)

}
