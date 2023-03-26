package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Openid      string `json:"openid"`
}

// 使用微信公众号的 appID 和 appSecret 获取 openid  静默授权
func GetWechatH5Openid(appID, appSecret, code string) (string, error) {
	// 获取 access_token
	tokenURL := "https://api.weixin.qq.com/sns/oauth2/access_token"
	tokenURL += fmt.Sprintf("?appid=%s&secret=%s&code=%s&grant_type=authorization_code", appID, appSecret, code)

	tokenResp, err := http.Get(tokenURL)
	if err != nil {
		panic(err)
	}
	defer tokenResp.Body.Close()

	var token AccessTokenResponse
	err = json.NewDecoder(tokenResp.Body).Decode(&token)
	if err != nil {
		panic(err)
	}
	return token.Openid, nil
}
