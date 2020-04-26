package responses

type Params interface {
	GetRequestId() string
	GetMessage() string
	GetCode() string
	IsOk() bool
}

//公共响应参数
type Common struct {
	RequestId, Message, Code string
}

//状态码的描述
func (c *Common) GetRequestId() string {
	return c.RequestId
}

//状态码的描述
func (c *Common) GetMessage() string {
	return c.Message
}

//请求状态码
func (c *Common) GetCode() string {
	return c.Code
}

//是否请求成功
func (c *Common) IsOk() bool {
	return c.GetCode() == "OK"
}
