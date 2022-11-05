package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"message-push/common"
	"message-push/model"
	"message-push/model/constant"
	"message-push/service"
	"regexp"
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

func (i *channelInfoApi) Update(c *gin.Context) {

	var updateVm model.ChannelInfoUpdateValidator

	err := c.ShouldBindJSON(&updateVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}
	err = common.GlobalUtil.CheckArrStr(
		0,
		0,
		common.IpV4Regx,
		updateVm.WhiteListStr,
	)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	//check channel type
	existInfo, err := service.Services.ChannelInfoService.GetOne(updateVm.ID)
	if err != nil {
		common.R.FailWithMsg(c, fmt.Sprintf("未找到id为 %d 的通道信息", updateVm.ID))
	}

	err = checkOption(existInfo.ChannelType, updateVm.OptionJsonStr)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	var info = model.ChannelInfo{}
	err = common.StructConvert[model.ChannelInfoUpdateValidator, model.ChannelInfo](updateVm, &info)
	err = service.Services.ChannelInfoService.Update(info)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}
	common.R.Success(c)

}

func checkOption(channelType int, jsonStr string) error {

	validate := validator.New()
	switch channelType {
	case constant.WEBHOOK:
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
	case constant.EMAIL:
		email := model.EmailChannel{}
		err := json.Unmarshal([]byte(jsonStr), &email)
		if err != nil {
			return errors.New("email项格式错误")
		}
		err = validate.Struct(email)
		if err != nil {
			return err
		}

		b, err := regexp.Match(common.UrlRegx, []byte(email.Host))
		if err != nil || !b {
			return errors.New("内容： host格式不正确")
		}

		err = common.GlobalUtil.CheckArrStr(
			1,
			0,
			common.EmailRegx,
			email.ToEmailListStr,
		)
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

func (i *channelInfoApi) RefreshKey(c *gin.Context) {

	var updateVm model.ChannelInfoUpdateValidator
	err := c.ShouldBindJSON(&updateVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	key, err := service.Services.ChannelInfoService.RefreshKey(updateVm.ID)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.SuccessWithData(c, key)

}
