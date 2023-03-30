package wechat

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func PostJson(url string, jsonBytes []byte) (str string, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body体数据错误", err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		// 处理请求失败的情况
		return "", errors.New("请求失败")
	}
	str = readString(respBytes)
	fmt.Println("返回的字符串数据", str)
	return str, nil
}

func readString(respBytes []byte) string {
	if len(respBytes) == 0 {
		return ""
	}
	return string(respBytes)
}

func GetUrl(url string) (res string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	res = string(body)
	fmt.Println("返回的结果", res)
	return
}
