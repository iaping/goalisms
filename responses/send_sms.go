package responses

//发送短信响应
type SendSms struct {
	Common
	BizId string
}

//发送回执ID，可根据该ID在接口QuerySendDetails中查询具体的发送状态。
func (s *SendSms) GetBizId() string {
	return s.BizId
}
