package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TextMessage struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

// 推送文本消息
func PushTextMessage(accessToken string, toUser string, content string) (res interface{}, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)

	message := TextMessage{
		ToUser:  toUser,
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}

	body, err := json.Marshal(message)
	if err != nil {
		return res, err
	}
	res, err = PostJson(url, body)
	return res, err
}

type TemplateMessage struct {
	ToUser     string                 `json:"touser"`
	TemplateID string                 `json:"template_id"`
	Data       map[string]interface{} `json:"data"`
}

// 推送模板消息
func SendTemplateMessage(accessToken string, templateMessage *TemplateMessage) error {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken)

	client := &http.Client{Timeout: 5 * time.Second}
	requestBody, err := json.Marshal(templateMessage)
	if err != nil {
		return err
	}

	resp, err := client.Post(url, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return fmt.Errorf("%s", result.ErrMsg)
	}

	return nil
}
