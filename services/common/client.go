package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
)

//Client http 客户端
type Client struct {
	httpClient      *http.Client
	credential      *Credential
	ContentType     string
	Version         string
	MaxIdleConns    int
	MaxConnsPerHost int
	Debug           bool
}

//NewClient  实例化http client
func NewClient(credential *Credential) *Client {
	client := &Client{
		credential:      credential,
		MaxIdleConns:    30,
		MaxConnsPerHost: 30,
		Debug:           false,
		httpClient: &http.Client{
			Timeout: 8 * time.Second, //设置HTTP超时时间
		},
	}
	return client
}

//Request 发起HTTP请求
func (c *Client) Request(action, url string, input []byte, retry int) (*HTTPResponse, error) {

	var err error
	response := &HTTPResponse{}

	if c.Debug {
		log.Debugf("[req]=>%s to %s \n%s\n", action, url, string(input))
	}

	//处理 http action
	if strings.ToUpper(action) == "POST" {
		action = "POST"
	} else if strings.ToUpper(action) == "GET" {
		action = "GET"
	} else {
		errorMsg := fmt.Sprintf(errors.UnsupportedTypeErrorMessage, action, "POST,GET")
		return nil, errors.NewClientError(errors.UnsupportedTypeErrorCode, errorMsg, err)
	}

	//判断json是否合法
	if !isJSONString(input) {
		errorMsg := fmt.Sprintf(errors.InvalidFormatErrorMessage, "input is not json format")
		return nil, errors.NewClientError(errors.InvalidParamErrorCode, errorMsg, err)
	}

	req, err := http.NewRequest(action, url, bytes.NewReader(input))
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	//添加权限认证字段
	for k, v := range c.credential.GetHeaders() {
		req.Header.Set(k, v)
		if c.Debug {
			log.Debugf("[headers]=>%s:%s\n", k, v)
		}
	}

	//设置Content-Type
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	if c.ContentType != "" {
		req.Header.Set("Content-Type", c.ContentType)
	}

	//默认 retry
	if retry <= 0 {
		retry = 1
	}
	var resp *http.Response
	for i := 0; i < retry; i++ {
		resp, err = c.httpClient.Do(req)
		if err == nil {
			break
		}

		time.Sleep(time.Second * 1)
	}

	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}
	if resp != nil {
		defer func() {
			err = resp.Body.Close()
		}()
	}

	response.StatusCode = resp.StatusCode
	response.Status = resp.Status
	response.OriginHTTPResponse = resp //原始的Http Response
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf(errors.NetWorkErrorMessage, err.Error())
		return response, errors.NewClientError(errors.NetWorkErrorCode, errMsg, err)
	}

	response.ResponseBodyBytes = bytes //http 响应体
	//如果StatusCode不等于200,则错误
	if response.StatusCode != 200 {
		response.Message = string(response.ResponseBodyBytes)
		return response, errors.NewServerError(resp.StatusCode, response.Message, err)
	}

	if c.Debug {
		log.Debugf("[resp]=>%s \n", bytes)
	}

	return response, nil
}

//SetTransport 设置Transport
func (c *Client) SetTransport(transport http.RoundTripper) {
	c.httpClient.Transport = transport
}

//SetHTTPTimeout 设置http 超时时间
func (c *Client) SetHTTPTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

//isJSONString 检查json能否被正确解析
func isJSONString(input []byte) bool {
	var x struct{}
	if err := json.Unmarshal(input, &x); err == nil {
		return true
	}
	return false
}
