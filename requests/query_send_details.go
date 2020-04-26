package requests

//查看短信发送记录和发送状态
type QuerySendDetails struct {
	Common
	CurrentPage, PageSize        uint8
	PhoneNumber, SendDate, BizId string
}

func NewQuerySendDetails() *QuerySendDetails {
	return &QuerySendDetails{Common: NewCommon("QuerySendDetails")}
}

//分页查看发送记录，指定发送记录的的当前页码。
func (q *QuerySendDetails) SetCurrentPage(currentPage uint8) {
	q.CurrentPage = currentPage
}

//分页查看发送记录，指定每页显示的短信记录数量。取值范围为1~50。
func (q *QuerySendDetails) SetPageSize(pageSize uint8) {
	q.PageSize = pageSize
}

//接收短信的手机号码。
//格式：
//国内短信：11位手机号码，例如15900000000。
//国际/港澳台消息：国际区号+号码，例如85200000000。
func (q *QuerySendDetails) SetPhoneNumber(phoneNumber string) {
	q.PhoneNumber = phoneNumber
}

//短信发送日期，支持查询最近30天的记录。格式为yyyyMMdd，例如20181225。
func (q *QuerySendDetails) SetSendDate(sendDate string) {
	q.SendDate = sendDate
}

//发送回执ID，即发送流水号。调用发送接口SendSms或SendBatchSms发送短信时，返回值中的BizId字段。
func (q *QuerySendDetails) SetBizId(bizId string) {
	q.BizId = bizId
}

func (q *QuerySendDetails) SetAccessKeyId(accessKeyId string) {
	q.AccessKeyId = accessKeyId
}

func (q *QuerySendDetails) Check() error {
	return nil
}
