package common

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	log "github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/zlog"
)

//Client http 客户端
type Client struct {
	httpClient      *http.Client
	credential      *Credential
	Timeout         int
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
		Timeout:         15,
		MaxIdleConns:    30,
		MaxConnsPerHost: 30,
		Debug:           false,
	}

	client.httpClient = client.defaultHTTPClient()
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
	req.Header.Set("Content-Type", "application/json")
	if c.ContentType != "" {
		req.Header.Set("Content-Type", c.ContentType)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return response, err
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
		//TODO:添加错误处理
		response.Message = string(response.ResponseBodyBytes)
		return response, errors.New(response.Message)
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
		return response, err
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
	return response, nil
}

//SetTransport 设置Transport
func (c *Client) SetTransport(transport http.RoundTripper) {
	c.httpClient.Transport = transport
}

//defaultHTTPClient 默认Transport
func (c *Client) defaultHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: time.Minute * time.Duration(c.Timeout), //设置超时时间,默认0不设置超时时间
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   15 * time.Second, //限制建立TCP连接的时间
				KeepAlive: 15 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second, //限制 TLS握手的时间
			ResponseHeaderTimeout: 10 * time.Second, //限制读取response header的时间
			ExpectContinueTimeout: 1 * time.Second,  //限制client在发送包含 Expect: 100-continue的header到收到继续发送body的response之间的时间等待。
			MaxIdleConns:          c.MaxIdleConns,   //连接池对所有host的最大连接数量，默认无穷大
			MaxConnsPerHost:       c.MaxIdleConns,   //连接池对每个host的最大连接数量。
			IdleConnTimeout:       30 * time.Minute, //how long an idle connection is kept in the connection pool.
			DisableKeepAlives:     false,
		},
	}

	return client
}
