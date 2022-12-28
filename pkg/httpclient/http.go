package httpclient

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	client = &http.Client{}
)

func Request(method string, u *url.URL, body io.Reader, header map[string]string) ([]byte, error) {
	var req *http.Request
	var resp *http.Response
	var respBody []byte
	var err error

	// 创建请求
	if req, err = http.NewRequest(method, u.String(), body); err != nil {
		return nil, err
	}

	// 设置header
	for k, v := range header {
		req.Header.Set(k, v)
	}

	// 发起请求
	if resp, err = client.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 获取结果
	if respBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return respBody, err
}
