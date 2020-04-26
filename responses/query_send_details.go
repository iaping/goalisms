package responses

//发送短信响应
type QuerySendDetails struct {
	Common
	TotalCount        int
	SmsSendDetailDTOs map[string][]SmsSendDetailDTO
}

//短信发送总条数。
func (q *QuerySendDetails) GetTotalCount() int {
	return q.TotalCount
}

//短信发送明细。
func (q *QuerySendDetails) GetSmsSendDetailDTOs() map[string][]SmsSendDetailDTO {
	return q.SmsSendDetailDTOs
}

type SmsSendDetailDTO struct {
	Content, ErrCode, OutId, PhoneNum, ReceiveDate, SendDate, TemplateCode string
	SendStatus                                                             int
}

//短信内容。
func (s *SmsSendDetailDTO) GetContent() string {
	return s.Content
}

//运营商短信状态码。
func (s *SmsSendDetailDTO) GetErrCode() string {
	return s.ErrCode
}

//外部流水扩展字段。
func (s *SmsSendDetailDTO) GetOutId() string {
	return s.OutId
}

//接收短信的手机号码。
func (s *SmsSendDetailDTO) GetPhoneNum() string {
	return s.PhoneNum
}

//
//短信接收日期和时间。
func (s *SmsSendDetailDTO) GetReceiveDate() string {
	return s.ReceiveDate
}

//短信发送日期和时间。
func (s *SmsSendDetailDTO) GetSendDate() string {
	return s.SendDate
}

//短信发送状态，包括：
//1：等待回执。
//2：发送失败。
//3：发送成功。
func (s *SmsSendDetailDTO) GetSendStatus() int {
	return s.SendStatus
}

//短信模板ID
func (s *SmsSendDetailDTO) GetTemplateCode() string {
	return s.TemplateCode
}
