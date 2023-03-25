package gohelper_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RequestInfo struct {
	Method  string
	URL     string
	Header  map[string]string
	body    io.Reader
	Timeout time.Duration
}

func HttpGetWithValues(request *RequestInfo, param url.Values, proxy string) (string, error) {
	uri, err := url.ParseRequestURI(request.URL)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	uri.RawQuery = param.Encode()
	request.URL = uri.String()
	request.Header["Content-Type"] = hycons.Form

	result, err := send(request, proxy)
	return result, err
}

func HttpPostWithValues(request *RequestInfo, param url.Values, proxy string) (string, error) {
	request.body = strings.NewReader(param.Encode())
	request.Header["Content-Type"] = hycons.Form
	result, err := send(request, proxy)
	return result, err
}

func HttpPostWithString(request *RequestInfo, param string, proxy string) (string, error) {
	request.body = strings.NewReader(param)
	request.Header["Content-Type"] = hycons.Form
	result, err := send(request, proxy)
	return result, err
}

func HttpPostWithJson(request *RequestInfo, param interface{}, proxy string) (string, error) {
	paramBytes, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	paramBody := bytes.NewReader(paramBytes)
	request.body = paramBody
	request.Header["Content-Type"] = hycons.Json
	result, err := send(request, proxy)
	return result, err
}

// send 发送请求
func send(client *RequestInfo, proxy string) (string, error) {

	if client == nil {
		return "", &ClientError{Info: "client is nil"}
	}

	if !(strings.EqualFold(client.Method, http.MethodPost) ||
		strings.EqualFold(client.Method, http.MethodPut) ||
		strings.EqualFold(client.Method, http.MethodDelete) ||
		strings.EqualFold(client.Method, http.MethodPatch) ||
		strings.EqualFold(client.Method, http.MethodGet)) {
		return "", &ClientError{Info: "the method is not in the scope, please edit"}
	}

	req, err := http.NewRequest(client.Method, client.URL, client.body)
	if err != nil {
		return "", err
	}

	if client.Header != nil {
		// set header
		for key, value := range client.Header {
			req.Header.Set(key, value)
		}
	}

	var resp *http.Response

	realClient := http.DefaultClient
	if client.Timeout >= 0 {
		realClient.Timeout = client.Timeout
	}

	if proxy != "" {
		//HTTP代理
		proxyAddress, _ := url.Parse(proxy)
		realClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyAddress),
		}
	}

	// start
	resp, err = realClient.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(resp.Body)

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return string(all), nil

}

type ClientError struct {
	Info string
}

func (ce *ClientError) Error() string {
	return ce.Info
}
