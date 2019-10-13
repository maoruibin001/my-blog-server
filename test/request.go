package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送GET请求
// url:请求地址
// response:请求返回的内容
func Get(url string) string {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
// content:请求放回的内容
func Post(url string, data interface{}, contentType string) string {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest(`POST`, url, bytes.NewBuffer(jsonStr))
	req.Header.Add(`content-type`, contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func main() {
	result := Get("https://88w87.com")
	fmt.Println(result)
}
//https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx95bd67c3b7c5c989&redirect_uri=https%3A%2F%2F88w87.com&response_type=code&scope=snsapi_userinfo#wechat_redirect
//https://api.weixin.qq.com/sns/oauth2/access_token?appid=wx95bd67c3b7c5c989&secret=2b5bf66816586287281b22c323c11743&code=021Zz77M1K9IH7132z4M1UTN6M1Zz77c&grant_type=authorization_code
//
//https://api.weixin.qq.com/sns/userinfo?access_token=26_JW9lagV1yfAZkWWJkvAakp-ot2Pid9I3NUkdDnUVdRz6iLxxH1VKd0LAUE2O894sq1Q8qyrQ-4WxJ8M7fy9Ip85Q0cjGUH44hQvZvqvg-sY&openid=oym6WuPQYckGFN0PMq1xaAha-8iU&lang=zh_CN
