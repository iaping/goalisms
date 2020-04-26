package goalisms

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/iaping/goalisms/requests"
	"github.com/iaping/goalisms/responses"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var responsesMap map[string]responses.Params = map[string]responses.Params{
	"SendSms":          &responses.SendSms{},
	"QuerySendDetails": &responses.QuerySendDetails{},
}

type Client struct {
	ServerUrl, AccessKeyId, AccessKeySecret string
}

func NewClient(accessKeyId, accessKeySecret string) *Client {
	return &Client{
		ServerUrl:       "http://dysmsapi.aliyuncs.com",
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
}

func (c *Client) Exec(r requests.Params) (responses.Params, error) {
	r.SetAccessKeyId(c.AccessKeyId)
	if err := r.Check(); err != nil {
		return nil, err
	}

	url, err := c.createRequestUrl(r)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	src, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := needResponse(r)
	if err := json.Unmarshal(src, response); err != nil {
		return nil, err
	}

	return response, nil
}

func needResponse(r requests.Params) responses.Params {
	if val, ok := responsesMap[r.GetAction()]; ok {
		return val
	}

	return &responses.Common{}
}

func (c *Client) createRequestUrl(r requests.Params) (string, error) {
	urlValues, err := query.Values(r)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s?%s&Signature=%s", c.ServerUrl, urlValues.Encode(), c.Signature(urlValues)), nil
}

func (c *Client) Signature(urlValues url.Values) string {
	signature := urlValues.Encode()
	signature = fmt.Sprintf("GET&%%2F&%s", specialReplace(url.QueryEscape(signature)))

	mac := hmac.New(sha1.New, []byte(c.AccessKeySecret+"&"))
	mac.Write([]byte(signature))
	base64String := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return specialReplace(url.QueryEscape(base64String))
}

//特殊URL编码这个是POP特殊的一种规则，即在一般的URLEncode后再增加三种字符替换：加号 （+）替换成 %20、星号 （*）替换成 %2A、 %7E 替换回波浪号 （~）
func specialReplace(url string) string {
	special := map[string]string{"+": "%20", "*": "%2A", "%7E": "~"}
	for old, new := range special {
		url = strings.ReplaceAll(url, old, new)
	}
	return url
}
