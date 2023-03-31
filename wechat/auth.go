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

// 获取 access_token
func GetWechatAccessToken(appID, appSecret, code string) (string, error) {
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
	return token.AccessToken, nil
}

// 获取h5的openid和access_token
func GetWechatH5OpenidAndAccessToken(appID, appSecret, code string) (openid, accessToken string, err error) {
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
	openid = token.Openid
	accessToken = token.AccessToken
	return
}

// 判断用户是否关注公众号
func IsSubscribe(openid, accessToken string) (bool, error) {
	url := "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openid + "&lang=zh_CN"
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	var result struct {
		Subscribe int `json:"subscribe"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return false, err
	}
	return result.Subscribe == 1, nil
}
