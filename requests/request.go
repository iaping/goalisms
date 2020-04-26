package requests

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

type Params interface {
	Check() error
	GetAction() string
	SetAccessKeyId(accessKeyId string)
}

//公共请求参数
type Common struct {
	Format, AccessKeyId, Action, RegionId, SignatureMethod, SignatureNonce, SignatureVersion, Timestamp, Version string
}

func NewCommon(action string) Common {
	return Common{
		Action:           action,
		Format:           "json",
		RegionId:         "cn-hangzhou",
		SignatureMethod:  "HMAC-SHA1",
		SignatureNonce:   fmt.Sprintf("%s", uuid.Must(uuid.NewV4())),
		SignatureVersion: "1.0",
		Timestamp:        time.Now().UTC().Format(time.RFC3339),
		Version:          "2017-05-25",
	}
}

func (c *Common) GetAction() string {
	return c.Action
}
