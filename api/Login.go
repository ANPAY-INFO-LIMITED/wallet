package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wallet/service/untils"
	"io"
	"net/http"
	"time"
)

var serverHost string = "http://132.232.109.84:8090"

const (
	LoginUrl   = "/login"
	RefreshUrl = "/refresh"
	HttpsUrl   = "http://"
)

func setServerHost(server string) string {
	serverHost = server
	return serverHost
}

func GetServerHost() string {
	return HttpsUrl + Cfg.BtlServerHost
}

func Login(username, password string) (string, error) {
	url := GetServerHost() + LoginUrl
	return login(url, username, password)
}
func Refresh(username, password string) (string, error) {
	url := GetServerHost() + RefreshUrl
	return refresh(url, username, password)
}

func login(url string, username string, password string) (string, error) {
	user := struct {
		Username string `gorm:"unique;column:user_name" json:"userName"` // 正确地将unique和column选项放在同一个gorm标签内
		Password string `gorm:"column:password" json:"password"`
	}{
		Username: username,
		Password: password,
	}
	requestBody, _ := json.Marshal(user)
	a, err := SendPostRequest(url, "", requestBody)
	if err != nil {
		return "", err
	}
	result := struct {
		Error string `json:"error"`
		Token string `json:"token"`
	}{}
	err = json.Unmarshal(a, &result)
	if err != nil {
		fmt.Println("An error occurred while unmarshalling the response body:", err)
	}
	if result.Error != "" {
		return "", fmt.Errorf(result.Error)
	}
	return result.Token, err
}

func refresh(url string, username string, password string) (string, error) {
	user := struct {
		Username string `gorm:"unique;column:user_name" json:"userName"` // 正确地将unique和column选项放在同一个gorm标签内
		Password string `gorm:"column:password" json:"password"`
	}{
		Username: username,
		Password: untils.GenerateExtMD5WithSalt(password),
	}
	requestBody, _ := json.Marshal(user)
	a, err := SendPostRequest(url, "", requestBody)
	if err != nil {
		return "", err
	}
	result := struct {
		Error string `json:"error"`
		Token string `json:"token"`
	}{}
	err = json.Unmarshal(a, &result)
	if err != nil {
		fmt.Println("An error occurred while unmarshalling the response body:", err)
	}
	if result.Error != "" {
		return "", fmt.Errorf(result.Error)
	}
	return result.Token, err
}
func SendPostRequest(url string, token string, requestBody []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 创建HTTP请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("An error occurred while creating an HTTP request:", err)
		return nil, err
	}
	req = req.WithContext(ctx)

	// 设置Authorization Header
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("An error occurred while sending a POST request:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("An error occurred while closing the HTTP response body:", err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}
