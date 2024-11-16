package verifykey

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAccessToken() (string, error) {
	url := "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"

	payload := "scope=GIGACHAT_API_PERS"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return "", err
	}
	AuthorizationKey := "MDk1NWUzNjAtNTlhMi00ZmMzLTk0YzktY2YyYjQxMTNhYWEyOmM0YjQ5MTg1LTdjMjktNGY3NC05ODBjLWZkNDJiYTYzMjc5Yg=="
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("RqUID", "0955e360-59a2-4fc3-94c9-cf2b4113aaa2")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", AuthorizationKey))

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Отключение проверки сертификата
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", err
	}

	accessToken, ok := responseData["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in response")
	}
	return accessToken, nil
}
