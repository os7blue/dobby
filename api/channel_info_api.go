package api

import (
	"dobby/common"
	"dobby/model"
	"dobby/model/constant"
	"dobby/service"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = checkOption(createVm.ChannelType, createVm.OptionJsonStr)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	if createVm.ChannelType == constant.WS {
		createVm.OptionJsonStr = fmt.Sprintf(
			`{"key":"%s"}`,
			uuid.New().String(),
		)
	}

	var info = model.ChannelInfo{}
	err = common.StructConvert[model.ChannelInfoCreateValidator, model.ChannelInfo](createVm, &info)
	err = service.ChannelInfoService.CreateOne(info)
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

	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	//check channel type
	existInfo, err := service.ChannelInfoService.GetOne(updateVm.ID)
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
	err = service.ChannelInfoService.Update(info)
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

	case constant.WS:

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

	result, count, err := service.ChannelInfoService.LoadPage(loadVm.Page, loadVm.Limit, loadVm.Param)
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

	err = service.ChannelInfoService.Delete(updateVm.ID)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.Success(c)

}

func (i *channelInfoApi) PushWS(c *gin.Context) {
	var vm model.PushWs
	err := c.ShouldBindJSON(&vm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = service.WsService.Send(vm.ID, vm.Content)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}
	common.R.Success(c)

}
