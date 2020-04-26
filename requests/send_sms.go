package requests

import (
	"encoding/json"
	"errors"
	"strings"
)

//发送短信
type SendSms struct {
	Common
	PhoneNumbers, SignName, TemplateCode, OutId, SmsUpExtendCode, TemplateParam string
}

func NewSendSms() *SendSms {
	return &SendSms{Common: NewCommon("SendSms")}
}

//支持对多个手机号码发送短信，手机号码之间以英文逗号（,）分隔。上限为1000个手机号码
func (s *SendSms) SetPhoneNumbers(phoneNumbers string) {
	s.PhoneNumbers = phoneNumbers
}

//短信模板变量对应的实际值，JSON格式。
func (s *SendSms) SetTemplateParam(templateParam map[string]string) {
	json, _ := json.Marshal(templateParam)
	s.TemplateParam = string(json)
}

//短信签名名称。请在控制台签名管理页面签名名称一列查看。
func (s *SendSms) SetSignName(signName string) {
	s.SignName = signName
}

//短信模板ID。请在控制台模板管理页面模板CODE一列查看。
func (s *SendSms) SetTemplateCode(templateCode string) {
	s.TemplateCode = templateCode
}

//外部流水扩展字段。
func (s *SendSms) SetOutId(outId string) {
	s.OutId = outId
}

//上行短信扩展码，无特殊需要此字段的用户请忽略此字段。
func (s *SendSms) SetSmsUpExtendCode(smsUpExtendCode string) {
	s.SmsUpExtendCode = smsUpExtendCode
}

func (s *SendSms) SetAccessKeyId(accessKeyId string) {
	s.AccessKeyId = accessKeyId
}

func (s *SendSms) Check() error {
	if s.AccessKeyId == "" {
		return errors.New("AccessKeyId is empty")
	}
	if s.TemplateCode == "" {
		return errors.New("template code is empty")
	}
	if s.SignName == "" {
		return errors.New("sign name is empty")
	}
	if s.PhoneNumbers == "" {
		return errors.New("phone numbers is empty")
	}
	if len(strings.Split(s.PhoneNumbers, ",")) > 1000 {
		return errors.New("phone numbers are greater than 1000")
	}

	return nil
}
