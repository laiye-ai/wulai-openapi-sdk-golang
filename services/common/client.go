package common

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/errors"
	log "github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/zlog"
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

//Post 发起post请求
func (c *Client) Post(url string, input []byte) (*HTTPResponse, error) {

	var err error
	response := &HTTPResponse{}

	if c.Debug {
		log.Debugf("[req]=>%s \n%s\n", url, string(input))
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(input))
	if err != nil {
		return response, err
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

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return response, errors.NewClientError("Client Error", err.Error(), err)
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
		return response, err
	}

	response.ResponseBodyBytes = bytes //http 响应体
	//如果StatusCode不等于200,则错误
	if response.StatusCode != 200 {
		response.Message = string(response.ResponseBodyBytes)
		return response, errors.NewServerError(resp.StatusCode, "Server Error", response.Message, err)
	}

	return response, nil
}

//Get 发起Get请求
func (c *Client) Get(url string) (*HTTPResponse, error) {
	var err error
	response := &HTTPResponse{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response, err
	}

	//添加权限认证字段
	for k, v := range c.credential.GetHeaders() {
		req.Header.Set(k, v)
		if c.Debug {
			log.Debugf("[headers]=>%s:%s\n", k, v)
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return response, errors.NewClientError("Client Error", err.Error(), err)
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
		response.Message = string(response.ResponseBodyBytes)
		return response, errors.NewServerError(resp.StatusCode, "Server Error", response.Message, err)
	}

	response.ResponseBodyBytes = bytes //http 响应体
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
