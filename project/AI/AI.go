package AI

import (
	verifykey "project/AI/verifyKey"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func AiResponse(nameRecruit, specialityRecruit string, birthDateRecruit, expirienceRecruit string, nameEmployee, specialityEmployee string, birthDateEmployee, expirienceEmployee string) (string, error) {
	accessToken, err := verifykey.GetAccessToken()
	if err != nil {
		log.Fatalf("Error getting access token: %v", err)
	}

	request := fmt.Sprintf("Возвращает ТОЛЬКО число с плавающей точкой (в формате 0.XX) (округленное до двух знаков) - совместимость между двумя людьми на основе их данных."+
		"Имя первого человека (того, кто устраивается на работу) - %s, Год рождения первого человека %s, Опыт работы первого человека (в годах) %s, Специальность первого человека в компании %s \n"+
		"Имя второго человека (уже работающего) - %s, Год рождения второго человека %s, Опыт работы второго человека (в годах) %s, Специальность второго человека в компании %s", nameRecruit, birthDateRecruit, expirienceRecruit, specialityRecruit,
		nameEmployee, birthDateEmployee, expirienceEmployee, specialityEmployee)

	content, err := getChatResponse(accessToken, request)
	if err != nil {
		return "0", err
	}

	return content, nil
}

func getChatResponse(accessToken string, answer string) (string, error) {
	url := "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"

	payload := map[string]interface{}{
		"model": "GigaChat",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": answer,
			},
		},
		"function_call": "auto",
		"functions": []map[string]interface{}{
			{
				"name":        "Совместимость",
				"description": "Возвращает ТОЛЬКО число с плавающей точкой (в формате 0.XX) (округленное до двух знаков) - совместимость между двумя людьми на основе их данных.",
				"parameters": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"first_person_name": map[string]interface{}{
							"type":        "string",
							"description": "Имя первого человека (того, кто устраивается на работу)",
						},
						"second_person_name": map[string]interface{}{
							"type":        "string",
							"description": "Имя второго человека (уже работающего)",
						},
						"first_person_specialty": map[string]interface{}{
							"type":        "string",
							"description": "Специальность первого человека в компании",
						},
						"second_person_specialty": map[string]interface{}{
							"type":        "string",
							"description": "Специальность второго человека в компании",
						},
						"first_person_birth_year": map[string]interface{}{
							"type":        "number",
							"description": "Год рождения первого человека",
						},
						"second_person_birth_year": map[string]interface{}{
							"type":        "number",
							"description": "Год рождения второго человека",
						},
						"first_person_experience": map[string]interface{}{
							"type":        "number",
							"description": "Опыт работы первого человека (в годах)",
						},
						"second_person_experience": map[string]interface{}{
							"type":        "number",
							"description": "Опыт работы второго человека (в годах)",
						},
					},
					"required": []interface{}{
						"first_person_name",
						"second_person_name",
						"first_person_specialty",
						"second_person_specialty",
						"first_person_birth_year",
						"second_person_birth_year",
						"first_person_experience",
						"second_person_experience",
					},
				},
			},
		},
		"stream":             false,
		"repetition_penalty": 1,
	}

	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytePayload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

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

	choices, ok := responseData["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("no choices found in response")
	}

	choice := choices[0].(map[string]interface{})
	message := choice["message"].(map[string]interface{})
	content := message["content"].(string)

	return content, nil
}
