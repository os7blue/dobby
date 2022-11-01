package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"message-push/common"
	model "message-push/model"
	"message-push/service"
)

type channelInfoApi struct {
}

func (i *channelInfoApi) Create(c *gin.Context) {

	var createVm model.ChannelInfoCreateValidator

	err := c.ShouldBindJSON(&createVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}
	err = common.GlobalUtil.CheckArrStr(
		0,
		0,
		common.IpV4Regx,
		createVm.WhiteListStr,
	)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = checkOption(createVm.ChannelType, createVm.OptionJsonStr)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	var info = model.ChannelInfo{}
	err = common.StructConvert[model.ChannelInfoCreateValidator, model.ChannelInfo](createVm, &info)
	err = service.Services.ChannelInfoService.CreateOne(info)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.Success(c)

}

func checkOption(channelType int, jsonStr string) error {

	validate := validator.New()
	switch channelType {
	case 1:
		hook := model.WebhookChannel{}
		err := json.Unmarshal([]byte(jsonStr), &hook)
		if err != nil {
			return errors.New("通道设置项格式错误")
		}
		err = validate.Struct(hook)
		if err != nil {
			return err
		}

		break

	}

	return nil
}

func (i *channelInfoApi) Load(c *gin.Context) {

	var loadVm model.PageBody[model.ChannelInfoSearchValidator]
	err := c.ShouldBindJSON(&loadVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	result, count, err := service.Services.ChannelInfoService.LoadPage(loadVm.Page, loadVm.Limit, loadVm.Param)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.SuccessWithDataCount(c, result, count)

}

func (i *channelInfoApi) Update(c *gin.Context) {

	var updateVm model.ChannelInfoUpdateValidator
	err := c.ShouldBindBodyWith(&updateVm, binding.JSON)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	paramMap := make(map[string]any) //注意该结构接受的内容
	err = c.ShouldBindBodyWith(&paramMap, binding.JSON)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	//var info model.ChannelInfo
	//err = common.StructConvert[model.ChannelInfoUpdateValidator, model.ChannelInfo](updateVm, &info)
	//if err != nil {
	//	common.R.FailWithMsg(c, err.Error())
	//	return
	//}
	err = service.Services.ChannelInfoService.Update(paramMap)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.Success(c)

}

func (i *channelInfoApi) Delete(c *gin.Context) {

	var updateVm model.ChannelInfoUpdateValidator
	err := c.ShouldBindJSON(&updateVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = service.Services.ChannelInfoService.Delete(updateVm.ID)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.Success(c)

}
