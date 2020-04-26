package main

import (
	"fmt"
	"github.com/iaping/goalisms"
	"github.com/iaping/goalisms/requests"
	"github.com/iaping/goalisms/responses"
)

func main() {
	client := goalisms.NewClient("LTAxxxX", "s9PmkoKxittu")

	//发送短信
	sendSms := requests.NewSendSms()
	sendSms.SetSignName("测试")
	sendSms.SetTemplateCode("SMS_888888888")
	sendSms.SetPhoneNumbers("18888888888")
	sendSms.SetTemplateParam(map[string]string{"code": "12345"})
	if c, err := client.Exec(sendSms); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c.(*responses.SendSms).GetCode())
		fmt.Println(c.(*responses.SendSms).GetMessage())
		fmt.Println(c.(*responses.SendSms).IsOk())
		fmt.Println(c.(*responses.SendSms).GetBizId())
	}

	//发送记录
	querySendDetails := requests.NewQuerySendDetails()
	querySendDetails.SetPhoneNumber("18888888888")
	querySendDetails.SetSendDate("20200101")
	if c, err := client.Exec(querySendDetails); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c.(*responses.QuerySendDetails).GetCode())
		fmt.Println(c.(*responses.QuerySendDetails).GetMessage())
		fmt.Println(c.(*responses.QuerySendDetails).IsOk())
		fmt.Println(c.(*responses.QuerySendDetails).GetTotalCount())
		fmt.Println(c.(*responses.QuerySendDetails).GetSmsSendDetailDTOs())
	}
}
